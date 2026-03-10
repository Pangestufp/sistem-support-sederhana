import { useLocation } from "react-router-dom"
import { navItems } from "./Sidebar"
import { getName } from "../shared/token"

export default function Topbar({ sidebarOpen }) {
  const location = useLocation()
  const name = getName()
  const currentLabel = navItems.find(n => n.path === location.pathname)?.label ?? "Page"

  return (
    <div className={`fixed top-0 right-0 z-20 h-16 border-b border-slate-700/60 bg-slate-900/80 backdrop-blur-sm flex items-center px-6 transition-all duration-300 justify-between ${sidebarOpen ? "left-56" : "left-16"}`}>
      <p className="text-sm font-semibold text-white">{currentLabel}</p>
      <p>{name}</p>
    </div>
  )
}