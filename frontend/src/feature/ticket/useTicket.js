import { useState } from "react"
import TicketService from "./ticket.service"

export function useTicket() {
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  const createTicket = async (payload) => {
    setLoading(true)
    setError(null)

    try {
      const response = await TicketService.create(payload)

      if (response.code !== 201) {
        throw new Error(response.message || "Gagal create ticket")
      }

      return response.data.id

    } catch (err) {
      setError(err.message || "Terjadi kesalahan")
      throw err
    } finally {
      setLoading(false)
    }
  }

  return {
    createTicket,
    loading,
    error
  }
}