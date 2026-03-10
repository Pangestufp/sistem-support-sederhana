import api from "../../app/axios";
import Endpoints from "../../shared/endpoints"

const BrowseService = {
  getAll: async (lastId = 0) => {
    const res = await api.get(Endpoints.TICKET.GET_ALL(lastId))
    return res.data
  }
}

export default BrowseService