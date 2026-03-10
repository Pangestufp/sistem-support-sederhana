import { useEffect, useState } from "react"
import TicketService from "./ticket.service"

export function useTicketDetail(id) {
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  const [ticket, setTicket] = useState(null)
  const [details, setDetails] = useState(null)
  const [attachments, setAttachments] = useState([])
  const [workflow, setWorkflow] = useState([])

  useEffect(() => {
    const fetchAll = async () => {
      setLoading(true)
      setError(null)

      try {
        const [
          ticketRes,
          detailRes,
          attachmentRes,
          workflowRes
        ] = await Promise.all([
          TicketService.getById(id),
          TicketService.getDetails(id),
          TicketService.getAttachments(id),
          TicketService.getWorkflow(id)
        ])

        setTicket(ticketRes.data)
        setDetails(detailRes.data||[])
        setAttachments(attachmentRes.data || [])
        setWorkflow(workflowRes.data || [])

      } catch (err) {
        setError("Gagal load ticket detail")
      } finally {
        setLoading(false)
      }
    }

    if (id) fetchAll()
  }, [id])

  return {
    loading,
    error,
    ticket,
    details,
    attachments,
    workflow
  }
}