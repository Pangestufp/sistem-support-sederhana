import { useState } from "react"
import { Outlet } from "react-router-dom"
import Sidebar from "./Sidebar"
import Topbar from "./Topbar"

export default function MainLayout() {
  const [open, setOpen] = useState(true)

  return (
    <div>
      <Sidebar open={open} setOpen={setOpen} />
      <Topbar sidebarOpen={open} />

      <main className={`relative z-10 pt-16 min-h-screen transition-all duration-300 ${open ? "ml-56" : "ml-16"}`}>
        <div className="p-6">
          <Outlet />
        </div>
      </main>
    </div>
  )
}