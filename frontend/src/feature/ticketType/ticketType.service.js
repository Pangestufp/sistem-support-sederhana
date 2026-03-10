import api from "../../app/axios"
import Endpoints from "../../shared/endpoints"

const TicketTypeService = {
  getAll: async () => {
    const res = await api.get(Endpoints.TICKET_TYPE.GET_ALL)
    return res.data
  },

  getById: async (id) => {
    const res = await api.get(Endpoints.TICKET_TYPE.GET_BY_ID(id))
    return res.data
  },
}

export default TicketTypeService