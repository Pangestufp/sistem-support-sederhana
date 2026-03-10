import api from "../../app/axios";
import Endpoints from "../../shared/endpoints";

const TicketService = {
  create: async (payload) => {
    const formData = new FormData();

    formData.append("ticket_type_id", payload.ticket_type_id);
    formData.append("description", payload.description);

    payload.pictures.forEach((file) => {
      formData.append("pictures", file);
    });

    const res = await api.post(Endpoints.TICKET.CREATE, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });

    return res.data;
  },
  getById: async (id) => {
    const res = await api.get(Endpoints.TICKET.GET_BY_ID(id));
    return res.data;
  },

  getDetails: async (id) => {
    const res = await api.get(Endpoints.TICKET.GET_DETAILS(id));
    return res.data;
  },

  getAttachments: async (id) => {
    const res = await api.get(Endpoints.TICKET.GET_ATTACHMENTS(id));
    return res.data;
  },

  getWorkflow: async (id) => {
    const res = await api.get(Endpoints.TICKET.GET_WORKFLOW(id));
    return res.data;
  },
  approve: async (id) => {
    const res = await api.post(Endpoints.TICKET.APPROVE(id));
    return res.data;
  },

  reject: async (id) => {
    const res = await api.post(Endpoints.TICKET.REJECT(id));
    return res.data;
  },

  returnTicket: async (id) => {
    const res = await api.post(Endpoints.TICKET.RETURN(id));
    return res.data;
  },

  review: async (id, payload) => {
    const formData = new FormData();
    formData.append("review", payload.review);

    payload.pictures?.forEach((file) => {
      formData.append("pictures", file);
    });

    const res = await api.post(Endpoints.TICKET.REVIEW(id), formData, {
      headers: { "Content-Type": "multipart/form-data" },
    });

    return res.data;
  },
};

export default TicketService;
