package service

import (
	"TicketManagement/dto"
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"TicketManagement/helper"
	"TicketManagement/repository"
	"fmt"
	"log"
	"strconv"
	"time"
)

type TicketService interface {
	CreateTicket(input dto.TicketRequest, filePaths []string, userID int) (*dto.TicketResponse, error)
	Action(ticketID, userID int, action helper.DocumentAction) error
	Return(ticketID, userID int, action helper.DocumentAction) error
	Review(req dto.TicketReviewRequest, filePaths []string, ticketID, userID int, action helper.DocumentAction) error
	GetAllJob(userID int) ([]dto.UserTicketResponse, error)
	GetTicket(ticketID int) (*dto.TicketResponse, error)
	GetTicketDetail(ticketID int) ([]dto.TicketDetailResponse, error)
	GetTicketAttachment(ticketID int) ([]dto.TicketAttachmentResponse, error)
	GetTicketWorkflowV(ticketID int) ([]dto.VTicketWorkflow, error)
	GetPaginateTicket(ticketID int) ([]dto.TicketResponse, error)
}

type ticketService struct {
	repositoryT  repository.TicketRepository
	repositoryTD repository.TicketDetailRepository
	repositoryTA repository.TicketAttachmentRepository
	repositoryTW repository.TicketWorkflowRepository
	repositoryUR repository.UserRoleRepository
	repositoryWP repository.WorkflowPathRepository
}

func NewTicketService(repositoryT repository.TicketRepository, repositoryTD repository.TicketDetailRepository, repositoryTA repository.TicketAttachmentRepository, repositoryTW repository.TicketWorkflowRepository, repositoryUR repository.UserRoleRepository, repositoryWP repository.WorkflowPathRepository) *ticketService {
	return &ticketService{
		repositoryT:  repositoryT,
		repositoryTD: repositoryTD,
		repositoryTA: repositoryTA,
		repositoryTW: repositoryTW,
		repositoryUR: repositoryUR,
		repositoryWP: repositoryWP,
	}
}

func (s *ticketService) CreateTicket(input dto.TicketRequest, filePaths []string, userID int) (*dto.TicketResponse, error) {

	docNo, err := s.repositoryT.GenerateDocNo()

	if err != nil {
		return nil, err
	}

	action := helper.ActionCreate
	ticket := entity.Ticket{
		DocNo:        docNo,
		CreatedBy:    userID,
		Description:  &input.Description,
		TicketTypeID: input.TicketTypeId,
		Status:       helper.GetNextStatusByAction(action),
		CreatedAt:    time.Now(),
	}

	if err := s.repositoryT.CreateTicket(&ticket); err != nil {
		return nil, err
	}

	var note string = "Add by creator"

	for _, path := range filePaths {
		attachment := entity.TicketAttachment{
			TicketID:  ticket.ID,
			FilePath:  path,
			Note:      &note,
			CreatedAt: time.Now(),
		}
		if err := s.repositoryTA.CreateTicketAttachment(&attachment); err != nil {
			continue
		}
	}

	s.assignWorkFlow(ticket.ID)

	data := dto.TicketResponse{
		ID:             ticket.ID,
		DocNo:          ticket.DocNo,
		CreatedBy:      ticket.CreatedBy,
		TicketTypeID:   ticket.TicketTypeID,
		TicketTypeName: "not used",
		Description:    *ticket.Description,
		Status:         *&ticket.Status,
		CreatedAt:      helper.FormatTimeRFC3339(ticket.CreatedAt),
		UpdatedAt:      helper.FormatTimeRFC3339(ticket.UpdatedAt),
	}

	return &data, nil
}

func (s *ticketService) assignWorkFlow(id int) error {

	vTicket, err := s.repositoryT.GetVTicketID(id)

	if err != nil {
		return err
	}

	workflowPaths, err := s.repositoryWP.GetAllWorkflowPath()

	if err != nil {
		return err
	}

	var ticketWorkflows []entity.TicketWorkflow

	for _, path := range workflowPaths {

		ok, err := s.repositoryT.IsConditionMet(
			id,
			path.ReadColumn,
			path.ExeCondition,
		)
		if err != nil || !ok {
			continue
		}

		log.Printf("[WF] condition met path_id=%d wf_id=%d pararel=%d", path.ID, path.WorkflowID, path.ParallelKey)
		userByRole, err := s.repositoryUR.GetUserByRoleName(path.AssignedTo)

		if len(userByRole) > 0 {
			for _, user := range userByRole {
				ticketWorkflows = append(ticketWorkflows, entity.TicketWorkflow{
					TicketID:       id,
					WorkflowPathID: path.ID,
					ParallelKey:    path.ParallelKey,
					AssignedUserID: user.UserID,
					Activity:       path.Activity,
				})
			}

		} else {
			assignedUser := resolveAssignedUser(vTicket, path.AssignedTo)

			if assignedUser == 0 {
				continue
			}

			ticketWorkflows = append(ticketWorkflows, entity.TicketWorkflow{
				TicketID:       id,
				WorkflowPathID: path.ID,
				ParallelKey:    path.ParallelKey,
				AssignedUserID: assignedUser,
				Activity:       path.Activity,
			})
		}

	}

	return s.repositoryTW.CreateTicketWorkflows(ticketWorkflows)

}

func resolveAssignedUser(v *entity.VTicket, assignedTo string) int {

	if id, err := strconv.Atoi(assignedTo); err == nil {
		return id
	}

	switch assignedTo {
	case "superior":
		if v.SuperiorID == nil {
			return 0
		}
		return *v.SuperiorID
	case "creator":
		return v.CreatedBy
	default:
		return 0
	}

}

func (s *ticketService) Action(ticketID, userID int, action helper.DocumentAction) error {

	ticket, err := s.repositoryT.GetTicketByID(ticketID)

	if err != nil {
		return err
	}

	ticketWorkflow, err := s.repositoryTW.GetCurrentTicketWorkflowByTicketID(ticketID, userID)

	if err != nil {
		return err
	}

	log.Println(ticketWorkflow.AssignedUserID)
	log.Println(userID)

	if ticketWorkflow.AssignedUserID != userID {
		log.Println("NOT EQUAL TRIGGERED")
		return &errorhandler.ForbiddenError{Message: "You cannot do this action"}
	}

	if !helper.CanDoAction(helper.ToDocumentStatus(ticket.Status), action, ticketWorkflow.Activity) {
		log.Println("NOT Pass Validation")
		log.Println(ticket.Status)
		log.Println(action)
		log.Println(ticketWorkflow.Activity)
		return &errorhandler.ForbiddenError{Message: "You cannot do this action"}
	}

	ticket.Status = helper.GetNextStatusByAction(action)
	err = s.repositoryT.UpdateTicket(ticket)

	if err != nil {
		return &errorhandler.InternalServerError{Message: "Error update ticket"}
	}

	now := time.Now()
	var actionStr = string(action)
	ticketWorkflow.ClosedAt = &now
	ticketWorkflow.Action = &actionStr

	err = s.repositoryTW.CloseTicketWorkflow(*ticketWorkflow)

	log.Println("sini 1")
	if err != nil {
		return err
	}

	log.Println("sini 2")

	s.repositoryTW.EnsureClosedIfWorkflowFinished(ticket.ID)
	return nil

}

func (s *ticketService) GetAllJob(userID int) ([]dto.UserTicketResponse, error) {

	ticketWorkflow, err := s.repositoryTW.GetCurrentWorkflow(userID)

	if err != nil {
		return nil, err
	}

	responses := []dto.UserTicketResponse{}

	for _, data := range ticketWorkflow {
		responses = append(responses,
			dto.UserTicketResponse{
				TicketID:       data.TicketID,
				DocNo:          data.DocNo,
				WorkflowPathID: data.WorkflowPathID,
				ParallelKey:    data.ParallelKey,
				AssignedUserID: data.AssignedUserID,
				ClosedAt:       data.ClosedAt,
				Action:         data.Action,
				Activity:       data.Activity,
			},
		)
	}

	return responses, nil
}

func (s *ticketService) Return(ticketID, userID int, action helper.DocumentAction) error {

	ticket, err := s.repositoryT.GetTicketByID(ticketID)

	if err != nil {
		return err
	}

	ticketWorkflow, err := s.repositoryTW.GetCurrentTicketWorkflowByTicketID(ticketID, userID)

	if err != nil {
		return err
	}

	if ticketWorkflow.AssignedUserID != userID {
		return &errorhandler.ForbiddenError{Message: "You cannot do this action"}
	}

	if !helper.CanDoAction(helper.ToDocumentStatus(ticket.Status), action, ticketWorkflow.Activity) {
		return &errorhandler.ForbiddenError{Message: "You cannot do this action"}
	}

	ticket.Status = helper.GetNextStatusByAction(action)
	err = s.repositoryT.UpdateTicket(ticket)

	if err != nil {
		return &errorhandler.InternalServerError{Message: "Error update ticket"}
	}

	return s.repositoryTW.ReopenLastWorkflow(ticket.ID)
}

func (s *ticketService) Review(req dto.TicketReviewRequest, filePaths []string, ticketID, userID int, action helper.DocumentAction) error {

	ticket, err := s.repositoryT.GetTicketByID(ticketID)

	if err != nil {
		return err
	}

	ticketWorkflow, err := s.repositoryTW.GetCurrentTicketWorkflowByTicketID(ticketID, userID)

	if err != nil {
		return err
	}

	if ticketWorkflow.AssignedUserID != userID {
		return &errorhandler.ForbiddenError{Message: "You cannot do this action"}
	}

	if !helper.CanDoAction(helper.ToDocumentStatus(ticket.Status), action, ticketWorkflow.Activity) {
		return &errorhandler.ForbiddenError{Message: "You cannot do this action"}
	}

	ticket.Status = helper.GetNextStatusByAction(action)
	err = s.repositoryT.UpdateTicket(ticket)

	if err != nil {
		return &errorhandler.InternalServerError{Message: "Error update ticket"}
	}

	var note string = fmt.Sprintf("Add By Reviewer with ID %d", userID)

	for _, path := range filePaths {
		attachment := entity.TicketAttachment{
			TicketID:  ticket.ID,
			FilePath:  path,
			Note:      &note,
			CreatedAt: time.Now(),
		}
		if err := s.repositoryTA.CreateTicketAttachment(&attachment); err != nil {
			continue
		}
	}

	ticketDetail := entity.TicketDetail{
		TicketID:  ticket.ID,
		UserID:    userID,
		Review:    &req.Review,
		CreatedAt: time.Now(),
	}

	err = s.repositoryTD.CreateTicketDetail(&ticketDetail)

	if err != nil {
		return err
	}

	now := time.Now()
	var actionStr = string(action)
	ticketWorkflow.ClosedAt = &now
	ticketWorkflow.Action = &actionStr
	err = s.repositoryTW.CloseTicketWorkflow(*ticketWorkflow)

	if err != nil {
		return err
	}

	s.repositoryTW.EnsureClosedIfWorkflowFinished(ticket.ID)
	return nil
}

func (s *ticketService) GetTicket(ticketID int) (*dto.TicketResponse, error) {

	ticket, err := s.repositoryT.GetVTicketID(ticketID)

	if err != nil {
		return nil, err
	}

	data := dto.TicketResponse{
		ID:             ticket.ID,
		DocNo:          ticket.DocNo,
		CreatedBy:      ticket.CreatedBy,
		TicketTypeID:   ticket.TicketTypeID,
		TicketTypeName: ticket.TicketTypeName,
		Description:    *ticket.Description,
		Status:         ticket.Status,
		CreatedAt:      helper.FormatTimeRFC3339(ticket.CreatedAt),
		UpdatedAt:      helper.FormatTimeRFC3339(ticket.UpdatedAt),
	}

	return &data, nil
}

func (s *ticketService) GetTicketDetail(ticketID int) ([]dto.TicketDetailResponse, error) {

	ticketDetails, err := s.repositoryTD.GetTicketDetailByTicketID(ticketID)
	if err != nil {
		return nil, err
	}

	datas := make([]dto.TicketDetailResponse, 0, len(ticketDetails))

	for _, data := range ticketDetails {
		resp := dto.TicketDetailResponse{
			ID:        data.ID,
			TicketID:  data.TicketID,
			UserID:    data.UserID,
			Review:    data.Review,
			CreatedAt: helper.FormatTimeRFC3339(data.CreatedAt),
		}

		datas = append(datas, resp)
	}

	return datas, nil
}

func (s *ticketService) GetTicketAttachment(ticketID int) ([]dto.TicketAttachmentResponse, error) {

	ticketAttachments, err := s.repositoryTA.GetTicketAttachmentByTicketID(ticketID)
	if err != nil {
		return nil, err
	}

	datas := make([]dto.TicketAttachmentResponse, 0, len(ticketAttachments))

	for _, data := range ticketAttachments {
		resp := dto.TicketAttachmentResponse{
			ID:        data.ID,
			TicketID:  data.TicketID,
			FilePath:  data.FilePath,
			Note:      data.Note,
			CreatedAt: helper.FormatTimeRFC3339(data.CreatedAt),
		}

		datas = append(datas, resp)
	}

	return datas, nil
}

func (s *ticketService) GetTicketWorkflowV(ticketID int) ([]dto.VTicketWorkflow, error) {

	datas, err := s.repositoryTW.GetTicketWorkflowV(ticketID)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (s *ticketService) GetPaginateTicket(ticketID int) ([]dto.TicketResponse, error) {

	tickets, err := s.repositoryT.GetTicketWithPagination(ticketID)

	if err != nil {
		return nil, err
	}

	var response []dto.TicketResponse

	for _, data := range tickets {
		response = append(response, dto.TicketResponse{
			ID:             data.ID,
			DocNo:          data.DocNo,
			CreatedBy:      data.CreatedBy,
			TicketTypeID:   data.TicketTypeID,
			TicketTypeName: "-",
			Description:    *data.Description,
			Status:         data.Status,
			CreatedAt:      helper.FormatTimeRFC3339(data.CreatedAt),
			UpdatedAt:      helper.FormatTimeRFC3339(data.UpdatedAt),
		},
		)
	}

	return response, nil
}
