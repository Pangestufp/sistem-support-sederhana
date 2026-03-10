import api from "../../app/axios"
import Endpoints from "../../shared/endpoints"

const AuthService = {
  login: async (payload) => {
    const res = await api.post(Endpoints.AUTH.LOGIN, payload)
    return res.data
  }
}

export default AuthService