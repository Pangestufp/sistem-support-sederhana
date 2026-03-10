import { useEffect, useState } from "react"
import BrowseService from "./browse.service"

export default function useBrowse() {
  const [tickets, setTickets] = useState([])
  const [lastId, setLastId] = useState(0)
  const [nextCursor, setNextCursor] = useState(null)
  const [loading, setLoading] = useState(false)

  const fetchTickets = async (cursor = 0) => {
    setLoading(true)

    try {
      const res = await BrowseService.getAll(cursor)

      setTickets(res.data)
      setNextCursor(res.paginate.last_id)
    } catch (err) {
      console.error(err)
    } finally {
      setLoading(false)
    }
  }

  const nextPage = () => {
    if (!nextCursor) return
    setLastId(nextCursor)
  }

  useEffect(() => {
    fetchTickets(lastId)
  }, [lastId])

  return {
    tickets,
    loading,
    nextPage
  }
}