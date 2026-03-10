import { useEffect, useState } from "react"
import TicketTypeService from "./ticketType.service"

export function useTicketType() {
  const [ticketTypes, setTicketTypes] = useState([])
  const [loading, setLoading] = useState(false)

  const fetchTicketTypes = async () => {
    setLoading(true)
    try {
      const data = await TicketTypeService.getAll()
      setTicketTypes(data)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchTicketTypes()
  }, [])

  return {
    ticketTypes,
    loading,
  }
}