import { useParams } from "react-router-dom"
import { useTicketDetail } from "./useTicketDetail"
import TicketWorkflowAction from "./ticketWorkflowAction"

function TicketDetail() {
  const { id } = useParams()
  const {
    loading,
    error,
    ticket,
    attachments,
    details,
    workflow
  } = useTicketDetail(id)

  if (loading) return <p>Loading ticket detail...</p>
  if (error) return <p style={{ color: "red" }}>{error}</p>
  if (!ticket) return <p>Ticket tidak ditemukan</p>

  return (
    <div>
        <div className="relative max-w-4xl mx-auto space-y-6">

        <div className="fixed -top-20 -left-20 w-72 h-72 bg-indigo-600 opacity-10 rounded-full blur-3xl pointer-events-none" />
        <div className="fixed -bottom-20 -right-20 w-72 h-72 bg-violet-600 opacity-10 rounded-full blur-3xl pointer-events-none" />

        <div className="mb-2">
            <p className="text-xs font-semibold tracking-widest text-indigo-400 uppercase mb-1">Ticket</p>
            <h2 className="text-2xl font-bold text-white">Ticket Detail</h2>
        </div>

        <div className="bg-slate-900 border border-slate-700/60 rounded-2xl overflow-hidden">
            <div className="h-1 w-full bg-gradient-to-r from-indigo-500 via-violet-500 to-purple-500" />
            <div className="p-6">
            <p className="text-xs font-semibold tracking-widest text-indigo-400 uppercase mb-4">Main Info</p>
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                {[
                { label: "ID", value: ticket.id },
                { label: "Doc No", value: ticket.doc_no },
                { label: "Type", value: ticket.ticket_type_name },
                { label: "Created At", value: new Date(ticket.created_at).toLocaleString("id-ID") },
                ].map(({ label, value }) => (
                <div key={label} className="bg-slate-800/50 border border-slate-700/40 rounded-xl px-4 py-3">
                    <p className="text-xs text-slate-500 uppercase tracking-wider mb-1">{label}</p>
                    <p className="text-sm text-white font-medium">{value}</p>
                </div>
                ))}

                <div className="bg-slate-800/50 border border-slate-700/40 rounded-xl px-4 py-3">
                <p className="text-xs text-slate-500 uppercase tracking-wider mb-1">Status</p>
                <span className="inline-flex items-center gap-1.5 text-xs font-semibold px-2.5 py-1 rounded-full bg-indigo-500/20 text-indigo-300 border border-indigo-500/30">
                    <span className="w-1.5 h-1.5 rounded-full bg-indigo-400 animate-pulse" />
                    {ticket.status}
                </span>
                </div>

                <div className="sm:col-span-2 bg-slate-800/50 border border-slate-700/40 rounded-xl px-4 py-3">
                <p className="text-xs text-slate-500 uppercase tracking-wider mb-1">Description</p>
                <p className="text-sm text-slate-300 leading-relaxed">{ticket.description}</p>
                </div>
            </div>
            </div>
        </div>

        <div className="bg-slate-900 border border-slate-700/60 rounded-2xl overflow-hidden">
            <div className="h-1 w-full bg-gradient-to-r from-indigo-500 via-violet-500 to-purple-500" />
            <div className="p-6">
            <p className="text-xs font-semibold tracking-widest text-indigo-400 uppercase mb-4">Attachments</p>
            {attachments.length === 0 ? (
                <div className="flex flex-col items-center justify-center py-8 text-slate-600">
                <svg className="w-8 h-8 mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13" />
                </svg>
                <p className="text-sm">Tidak ada attachment</p>
                </div>
            ) : (
                <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-3">
                {attachments.map((file) => (
                    <div key={file.id} className="group relative bg-slate-800 border border-slate-700/60 rounded-xl overflow-hidden hover:border-indigo-500/50 transition">
                    <img
                        src={`${import.meta.env.VITE_API_URL}/${file.file_path}`}
                        alt="attachment"
                        className="w-full h-32 object-cover"
                    />
                    {file.note && (
                        <div className="px-3 py-2">
                        <p className="text-xs text-slate-400 truncate">{file.note}</p>
                        </div>
                    )}
                    </div>
                ))}
                </div>
            )}
            </div>
        </div>

        <div className="bg-slate-900 border border-slate-700/60 rounded-2xl overflow-hidden">
            <div className="h-1 w-full bg-gradient-to-r from-indigo-500 via-violet-500 to-purple-500" />
            <div className="p-6">
            <p className="text-xs font-semibold tracking-widest text-indigo-400 uppercase mb-4">Reviews</p>
            {details.length === 0 ? (
                <p className="text-sm text-slate-600 text-center py-6">Belum ada review</p>
            ) : (
                <div className="space-y-3">
                {details.map((item) => (
                    <div key={item.id} className="bg-slate-800/50 border border-slate-700/40 rounded-xl px-4 py-4">
                    <div className="flex items-center justify-between mb-2">
                        <span className="text-xs font-semibold text-indigo-300 bg-indigo-500/10 border border-indigo-500/20 px-2 py-0.5 rounded-full">
                        User #{item.user_id}
                        </span>
                        <span className="text-xs text-slate-500">{new Date(item.created_at).toLocaleString("id-ID")}</span>
                    </div>
                    <p className="text-sm text-slate-300 leading-relaxed">{item.review}</p>
                    </div>
                ))}
                </div>
            )}
            </div>
        </div>

        <div className="bg-slate-900 border border-slate-700/60 rounded-2xl overflow-hidden">
            <div className="h-1 w-full bg-gradient-to-r from-indigo-500 via-violet-500 to-purple-500" />
            <div className="p-6">
            <TicketWorkflowAction
                ticketId={id}
                onSuccess={() => window.location.reload()}
            />
            </div>
        </div>

        <div className="bg-slate-900 border border-slate-700/60 rounded-2xl overflow-hidden">
            <div className="h-1 w-full bg-gradient-to-r from-indigo-500 via-violet-500 to-purple-500" />
            <div className="p-6">
            <p className="text-xs font-semibold tracking-widest text-indigo-400 uppercase mb-4">Workflow</p>
            {workflow.length === 0 ? (
                <p className="text-sm text-slate-600 text-center py-6">Tidak ada workflow</p>
            ) : (
                <div className="relative">
                <div className="absolute left-4 top-0 bottom-0 w-px bg-slate-700/60" />
                <div className="space-y-4">
                    {workflow.map((step, index) => (
                    <div key={index} className="relative pl-10">
                        <div className={`absolute left-2.5 top-3 w-3 h-3 rounded-full border-2 -translate-x-1/2 ${step.closed_at ? "bg-indigo-500 border-indigo-400" : "bg-slate-700 border-slate-500"}`} />
                        <div className="bg-slate-800/50 border border-slate-700/40 rounded-xl px-4 py-4">
                        <div className="flex items-center justify-between mb-3">
                            <span className="text-xs font-bold text-white">Step {step.step} — {step.name}</span>
                            {step.closed_at && (
                            <span className="text-xs text-slate-500">{step.closed_at}</span>
                            )}
                        </div>
                        <div className="grid grid-cols-2 gap-x-6 gap-y-1.5 text-xs">
                            <div>
                            <span className="text-slate-500">Assigned: </span>
                            <span className="text-slate-300">{step.assigned_user}</span>
                            </div>
                            <div>
                            <span className="text-slate-500">Activity: </span>
                            <span className="text-slate-300">{step.activity}</span>
                            </div>
                            <div>
                            <span className="text-slate-500">Action: </span>
                            <span className="text-slate-300">{step.action || "—"}</span>
                            </div>
                        </div>
                        </div>
                    </div>
                    ))}
                </div>
                </div>
            )}
            </div>
        </div>

        </div>
    </div>
    )
}

export default TicketDetail