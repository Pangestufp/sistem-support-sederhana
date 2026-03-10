import { Navigate } from "react-router-dom"
import { getToken, getRoles } from "../shared/token"

function ProtectedAdmin({ children }) {
  const token = getToken()
  const roles = getRoles()

  if (!token) {
    return <Navigate to="/login" replace />
  }

  if (!roles.includes("101")) {
    return <Navigate to="/dashboard" replace />
  }

  return children
}

export default ProtectedAdmin