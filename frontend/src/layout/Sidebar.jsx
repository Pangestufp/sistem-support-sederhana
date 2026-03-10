import { useNavigate, useLocation } from "react-router-dom"
import { useAuth } from "../feature/auth/useAuth"

const navItems = [
  {
    label: "Dashboard",
    path: "/dashboard",
    icon: <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />,
  },
  {
    label: "Create Ticket",
    path: "/create",
    icon: <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />,
  },
  {
    label: "Browse Ticket",
    path: "/browse",
    icon: <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 10h16M4 14h16M4 18h16" />,
  },
]

export { navItems }

export default function Sidebar({ open, setOpen }) {
  const navigate = useNavigate()
  const location = useLocation()
  const { logout } = useAuth()

  const handleLogout = () => {
    try { 
        logout() 
        navigate("/login")
    } catch {}
  }

  return (
    <aside className={`fixed left-0 top-0 z-30 h-screen flex flex-col bg-slate-900 border-r border-slate-700/60 transition-all duration-300 ${open ? "w-56" : "w-16"}`}>

      <div className="flex items-center justify-between px-4 h-16 border-b border-slate-700/60 shrink-0">
        {open && (
          <span className="text-sm font-black tracking-tight text-white truncate" style={{ fontFamily: "'Syne', sans-serif" }}>
            Support
          </span>
        )}
        <button
          onClick={() => setOpen(!open)}
          className="ml-auto p-1.5 rounded-lg text-slate-400 hover:text-white hover:bg-slate-800 transition"
        >
          <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            {open
              ? <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 19l-7-7 7-7m8 14l-7-7 7-7" />
              : <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 5l7 7-7 7M5 5l7 7-7 7" />
            }
          </svg>
        </button>
      </div>

      <nav className="flex-1 px-2 py-4 space-y-1 overflow-y-auto">
        {navItems.map(({ label, path, icon }) => {
          const active = location.pathname === path
          return (
            <button
              key={path}
              onClick={() => navigate(path)}
              title={!open ? label : undefined}
              className={`w-full flex items-center gap-3 px-3 py-2.5 rounded-xl text-sm font-medium transition select-none ${
                active
                  ? "bg-indigo-600 text-white shadow-lg shadow-indigo-500/20"
                  : "text-slate-400 hover:text-white hover:bg-slate-800"
              }`}
            >
              <svg className="w-4 h-4 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                {icon}
              </svg>
              {open && <span className="truncate">{label}</span>}
            </button>
          )
        })}
      </nav>

      <div className="px-2 py-4 border-t border-slate-700/60 shrink-0">
        <button
          onClick={handleLogout}
          title={!open ? "Logout" : undefined}
          className="w-full flex items-center gap-3 px-3 py-2.5 rounded-xl text-sm font-medium text-slate-400 hover:text-red-400 hover:bg-red-500/10 transition select-none"
        >
          <svg className="w-4 h-4 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
          </svg>
          {open && <span>Logout</span>}
        </button>
      </div>

    </aside>
  )
}