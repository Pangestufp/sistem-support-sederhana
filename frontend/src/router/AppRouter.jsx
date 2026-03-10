import { Routes, Route, Navigate } from "react-router-dom"
import { getToken } from "../shared/token"

import ProtectedRoute from "./ProtectedRoute"
import Login from "../feature/auth/Login"
import Dashboard from "../feature/dashboard/Dashboard"
import CreateTicket from "../feature/ticket/createTicket"
import TicketDetail from "../feature/ticket/ticketDetail"
import MainLayout from "../layout/MainLayout"
import Browse from "../feature/browse/Browse"
import ProtectedAdmin from "./ProtectedAdmin"

function AppRouter() {
  const token = getToken()

  return (
    <Routes>
      <Route
        path="/login"
        element={
          token ? <Navigate to="/dashboard" replace /> : <Login />
        }
      />

      <Route
        element={
          <ProtectedRoute>
            <MainLayout/>
          </ProtectedRoute>
        }
      >

        <Route
          path="/dashboard"
          element={
            <ProtectedRoute>
              <Dashboard />
            </ProtectedRoute>
          }
        />

        <Route
          path="/create"
          element={
            <ProtectedRoute>
              <CreateTicket/>
            </ProtectedRoute>
          }
        />

        <Route
          path="/ticket/:id"
          element={
            <ProtectedRoute>
              <TicketDetail/>
            </ProtectedRoute>
          }
        />

        <Route
          path="/browse"
          element={
            <ProtectedAdmin>
              <Browse/>
            </ProtectedAdmin>
          }
        />
      
        <Route path="*" element={<Navigate to="/dashboard" />} /> 
      </Route>

    </Routes>
  )
}

export default AppRouter