import { useState } from "react"
import { clearToken, setToken } from "../../shared/token"
import AuthService from "./auth.service"
import { useNavigate } from "react-router-dom"

export function useAuth() {
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  const login = async (payload) => {
    setLoading(true)
    setError(null)

    try {
      const data = await AuthService.login(payload)
      setToken(data.data.token)
      return data
    } catch (err) {
      setError("Login gagal")
      throw err
    } finally {
      setLoading(false)
    }
  }

  const logout = () => {
    clearToken()
  }

  return { login, logout, loading, error }
}

