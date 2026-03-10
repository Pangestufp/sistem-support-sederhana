import { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"
import { useTicket } from "../ticket/useTicket"
import TicketTypeService from "../ticketType/ticketType.service"

function CreateTicket() {
  const navigate = useNavigate()
  const { createTicket, loading, error } = useTicket()

  const [ticketTypes, setTicketTypes] = useState([])
  const [form, setForm] = useState({
    ticket_type_id: "",
    description: "",
    pictures: []
  })

  const [loadingType, setLoadingType] = useState(false)

  useEffect(() => {
    const fetchTicketTypes = async () => {
      setLoadingType(true)
      try {
        const res = await TicketTypeService.getAll()
        setTicketTypes(res.data)
      } catch (err) {
        console.error("Gagal ambil ticket type")
      } finally {
        setLoadingType(false)
      }
    }

    fetchTicketTypes()
  }, [])

  const handleChange = (e) => {
    const { name, value } = e.target
    setForm((prev) => ({
      ...prev,
      [name]: value
    }))
  }

  const handleFileChange = (e) => {
    const newFiles = Array.from(e.target.files)
    setForm((prev) => ({
        ...prev,
        pictures: [...prev.pictures, ...newFiles]
    }))
  }

  const handleSubmit = async (e) => {
    e.preventDefault()

    try {
      const id = await createTicket(form)
      navigate(`/ticket/${id}`)
    } catch (err) {
      console.log("Create gagal")
    }
  }

  return (
    <div>


      <div className="relative w-full max-w-5xl min-w-2xl">
        <div className="absolute -top-20 -left-20 w-72 h-72 bg-indigo-600 opacity-10 rounded-full blur-3xl pointer-events-none" />
        <div className="absolute -bottom-20 -right-20 w-72 h-72 bg-violet-600 opacity-10 rounded-full blur-3xl pointer-events-none" />

        <div className="relative bg-slate-900 border border-slate-700/60 rounded-2xl shadow-2xl overflow-hidden">

          <div className="h-1 w-full bg-gradient-to-r from-indigo-500 via-violet-500 to-purple-500" />

          <div className="p-8">
            <div className="mb-8">
              <p className="text-xs font-semibold tracking-widest text-indigo-400 uppercase mb-1">New Request</p>
              <h2 className="text-2xl font-bold text-white">Create Ticket</h2>
            </div>

            <form onSubmit={handleSubmit} className="space-y-6">

              <div>
                <label className="block text-xs font-semibold tracking-wider text-slate-400 uppercase mb-2 text-left">
                  Ticket Type
                </label>
                {loadingType ? (
                  <div className="flex items-center gap-2 text-slate-500 text-sm py-3">
                    <svg className="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
                      <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"/>
                      <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8z"/>
                    </svg>
                    Loading ticket types...
                  </div>
                ) : (
                  <select
                    name="ticket_type_id"
                    value={form.ticket_type_id}
                    onChange={handleChange}
                    required
                    className="w-full bg-slate-800 border border-slate-700 text-white text-sm rounded-lg px-4 py-3 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition appearance-none cursor-pointer"
                  >
                    <option value="">-- Select Ticket Type --</option>
                    {ticketTypes.map((type) => (
                      <option key={type.id} value={type.id}>{type.name}</option>
                    ))}
                  </select>
                )}
              </div>

              <div>
                <label className="block text-xs font-semibold tracking-wider text-slate-400 uppercase mb-2 text-left">
                  Description
                </label>
                <textarea
                  name="description"
                  value={form.description}
                  onChange={handleChange}
                  rows="3"
                  required
                  placeholder="Describe your issue..."
                  className="w-full bg-slate-800 border border-slate-700 text-white text-sm rounded-lg px-4 py-3 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition resize-none placeholder-slate-600"
                />
              </div>

              <div>
                <label className="block text-xs font-semibold tracking-wider text-slate-400 uppercase mb-2 text-left">
                  Attachments
                </label>

                <label className="flex flex-col items-center justify-center w-full h-24 border-2 border-dashed border-slate-700 rounded-lg cursor-pointer bg-slate-800/50 hover:border-indigo-500 hover:bg-slate-800 transition">
                  <div className="flex flex-col items-center gap-1">
                    <svg className="w-6 h-6 text-slate-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M12 4v16m8-8H4" />
                    </svg>
                    <span className="text-xs text-slate-500">Click to upload files</span>
                  </div>
                  <input type="file" multiple onChange={handleFileChange} className="hidden" />
                </label>

                {form.pictures.length > 0 && (
                  <div className="mt-3 space-y-2">
                    {form.pictures.map((file, index) => (
                      <div key={index} className="flex items-center justify-between bg-slate-800 border border-slate-700 rounded-lg px-3 py-2">
                        <div className="flex items-center gap-2 min-w-0">
                          <svg className="w-4 h-4 text-indigo-400 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13" />
                          </svg>
                          <span className="text-sm text-slate-300 truncate">{file.name}</span>
                        </div>
                        <button
                          type="button"
                          onClick={() => setForm((prev) => ({ ...prev, pictures: prev.pictures.filter((_, i) => i !== index) }))}
                          className="text-slate-500 hover:text-red-400 transition ml-2 shrink-0"
                        >
                          <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                          </svg>
                        </button>
                      </div>
                    ))}
                  </div>
                )}
              </div>

              {error && (
                <div className="flex items-center gap-2 bg-red-500/10 border border-red-500/30 rounded-lg px-4 py-3">
                  <svg className="w-4 h-4 text-red-400 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  <p className="text-sm text-red-400">{error}</p>
                </div>
              )}

              <button
                type="submit"
                disabled={loading}
                className="w-full bg-gradient-to-r from-indigo-600 to-violet-600 hover:from-indigo-500 hover:to-violet-500 disabled:opacity-50 disabled:cursor-not-allowed text-white font-semibold text-sm py-3 rounded-lg transition flex items-center justify-center gap-2"
              >
                {loading ? (
                  <>
                    <svg className="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
                      <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"/>
                      <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8z"/>
                    </svg>
                    Creating...
                  </>
                ) : "Create Ticket"}
              </button>

            </form>
          </div>
        </div>
      </div>
    </div>
  )
}

export default CreateTicket