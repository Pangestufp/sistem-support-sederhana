import { useState } from "react"
import { useTicketAction } from "./useTicketAction"

function ConfirmModal({ action, onConfirm, onCancel, loading }) {
  const config = {
    approve: {
        title: "Approve Ticket?",
        desc: "Ticket akan disetujui dan dilanjutkan ke tahap berikutnya.",
        confirmStyle: { backgroundColor: "#059669" }, // emerald-600
        confirmHoverColor: "#10b981",
        iconBg: "bg-emerald-500/10 border-emerald-500/20",
        icon: <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />,
        iconColor: "text-emerald-400",
        label: "Approve",
    },
    reject: {
        title: "Reject Ticket?",
        desc: "Ticket akan ditolak dan prosesnya dihentikan.",
        confirmStyle: { backgroundColor: "#dc2626" }, // red-600
        iconBg: "bg-red-500/10 border-red-500/20",
        icon: <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />,
        iconColor: "text-red-400",
        label: "Reject",
    },
    return: {
        title: "Return Ticket?",
        desc: "Ticket akan dikembalikan ke step sebelumnya.",
        confirmStyle: { backgroundColor: "#d97706" }, // amber-600
        iconBg: "bg-amber-500/10 border-amber-500/20",
        icon: <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6" />,
        iconColor: "text-amber-400",
        label: "Return",
    },
    }

  const c = config[action]

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center px-4">
      <div className="absolute inset-0 bg-slate-950/80 backdrop-blur-sm" onClick={onCancel} />
      <div className="relative bg-slate-900 border border-slate-700/60 rounded-2xl shadow-2xl w-full max-w-sm overflow-hidden">
        <div className="h-1 w-full bg-gradient-to-r from-indigo-500 via-violet-500 to-purple-500" />
        <div className="p-6">
          <div className={`inline-flex items-center justify-center w-12 h-12 rounded-xl border ${c.iconBg} mb-4`}>
            <svg className={`w-6 h-6 ${c.iconColor}`} fill="none" stroke="currentColor" viewBox="0 0 24 24">
              {c.icon}
            </svg>
          </div>
          <h3 className="text-lg font-bold text-white mb-1">{c.title}</h3>
          <p className="text-sm text-slate-400 mb-6">{c.desc}</p>
          <div className="flex gap-3">
            <button
              onClick={onCancel}
              disabled={loading}
              className="flex-1 py-2.5 rounded-lg text-sm font-semibold bg-slate-800 border border-slate-700 text-slate-300 hover:bg-slate-700 disabled:opacity-40 transition"
            >
              Batal
            </button>
            <button
              onClick={onConfirm}
              disabled={loading}
              style={c.confirmStyle}
              className={`flex-1 py-2.5 rounded-lg text-sm font-semibold text-white disabled:opacity-50 disabled:cursor-not-allowed transition flex items-center justify-center gap-2`}
            >
              {loading ? (
                <>
                  <svg className="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
                    <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"/>
                    <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8z"/>
                  </svg>
                  Processing...
                </>
              ) : c.label}
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}

function ReviewModal({ onSubmit, onCancel, loading, reviewText, setReviewText, pictures, setPictures }) {
  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center px-4">
      <div className="absolute inset-0 bg-slate-950/80 backdrop-blur-sm" onClick={onCancel} />

      <div className="relative bg-slate-900 border border-slate-700/60 rounded-2xl shadow-2xl w-full max-w-lg overflow-hidden">
        <div className="h-1 w-full bg-gradient-to-r from-indigo-500 via-violet-500 to-purple-500" />
        <div className="p-6 space-y-5">
          <div>
            <p className="text-xs font-semibold tracking-widest text-indigo-400 uppercase mb-1">Workflow Action</p>
            <h3 className="text-lg font-bold text-white">Submit Review</h3>
          </div>

          <div>
            <label className="block text-xs font-semibold tracking-wider text-slate-400 uppercase mb-2">
              Review
            </label>
            <textarea
              placeholder="Tulis review..."
              value={reviewText}
              onChange={(e) => setReviewText(e.target.value)}
              rows="4"
              className="w-full bg-slate-800 border border-slate-700 text-white text-sm rounded-lg px-4 py-3 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition resize-none placeholder-slate-600"
            />
          </div>

          <div>
            <label className="block text-xs font-semibold tracking-wider text-slate-400 uppercase mb-2">
              Attachments
            </label>
            <label className="flex flex-col items-center justify-center w-full h-20 border-2 border-dashed border-slate-700 rounded-lg cursor-pointer bg-slate-800/50 hover:border-indigo-500 hover:bg-slate-800 transition">
              <div className="flex items-center gap-2">
                <svg className="w-5 h-5 text-slate-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M12 4v16m8-8H4" />
                </svg>
                <span className="text-xs text-slate-500">Click to upload files</span>
              </div>
              <input
                type="file"
                multiple
                onChange={(e) => setPictures(Array.from(e.target.files))}
                className="hidden"
              />
            </label>

            {pictures.length > 0 && (
              <div className="mt-3 space-y-2">
                {pictures.map((file, index) => (
                  <div key={index} className="flex items-center justify-between bg-slate-800 border border-slate-700 rounded-lg px-3 py-2">
                    <div className="flex items-center gap-2 min-w-0">
                      <svg className="w-4 h-4 text-indigo-400 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13" />
                      </svg>
                      <span className="text-sm text-slate-300 truncate">{file.name}</span>
                    </div>
                    <button
                      type="button"
                      onClick={() => setPictures((prev) => prev.filter((_, i) => i !== index))}
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

          <div className="flex gap-3 pt-1">
            <button
              onClick={onCancel}
              disabled={loading}
              className="flex-1 py-2.5 rounded-lg text-sm font-semibold bg-slate-800 border border-slate-700 text-slate-300 hover:bg-slate-700 disabled:opacity-40 transition"
            >
              Batal
            </button>
            <button
              onClick={onSubmit}
              disabled={loading}
              className="flex-1 bg-gradient-to-r from-indigo-600 to-violet-600 hover:from-indigo-500 hover:to-violet-500 disabled:opacity-50 disabled:cursor-not-allowed text-white font-semibold text-sm py-2.5 rounded-lg transition flex items-center justify-center gap-2"
            >
              {loading ? (
                <>
                  <svg className="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
                    <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"/>
                    <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8z"/>
                  </svg>
                  Submitting...
                </>
              ) : "Submit Review"}
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}

function TicketWorkflowAction({ ticketId, onSuccess }) {
  const { approve, reject, returnTicket, review, loading, error } = useTicketAction()

  const [confirmAction, setConfirmAction] = useState(null) // "approve" | "reject" | "return"
  const [showReviewModal, setShowReviewModal] = useState(false)
  const [reviewText, setReviewText] = useState("")
  const [pictures, setPictures] = useState([])

  const handleConfirm = async () => {
    if (loading) return

    let success = false
    if (confirmAction === "approve") success = await approve(ticketId)
    if (confirmAction === "reject") success = await reject(ticketId)
    if (confirmAction === "return") success = await returnTicket(ticketId)
    setConfirmAction(null)
    if (success) {
      if (onSuccess) onSuccess()
    }
  }

  const handleReview = async () => {
    const success = await review(ticketId, { review: reviewText, pictures })
    if (success) {
      setReviewText("")
      setPictures([])
      setShowReviewModal(false)
      if (onSuccess) onSuccess()
    }
  }

  return (
    <div className="space-y-4">
      <p className="text-xs font-semibold tracking-widest text-indigo-400 uppercase">Workflow Action</p>

      {error && (
        <div className="flex items-center gap-2 bg-red-500/10 border border-red-500/30 rounded-lg px-4 py-3">
          <svg className="w-4 h-4 text-red-400 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <p className="text-sm text-red-400">{error}</p>
        </div>
      )}

      <div className="flex flex-wrap gap-3">
        <button
          onClick={() => setConfirmAction("approve")}
          disabled={loading}
          className="flex items-center gap-2 px-4 py-2.5 rounded-lg text-sm font-semibold bg-emerald-500/10 border border-emerald-500/30 text-emerald-400 hover:bg-emerald-500/20 hover:border-emerald-400/50 disabled:opacity-40 disabled:cursor-not-allowed transition active:scale-95 select-none"
        >
          <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
          </svg>
          Approve
        </button>

        <button
          onClick={() => setConfirmAction("reject")}
          disabled={loading}
          className="flex items-center gap-2 px-4 py-2.5 rounded-lg text-sm font-semibold bg-red-500/10 border border-red-500/30 text-red-400 hover:bg-red-500/20 hover:border-red-400/50 disabled:opacity-40 disabled:cursor-not-allowed transition active:scale-95 select-none"
        >
          <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
          </svg>
          Reject
        </button>

        <button
          onClick={() => setConfirmAction("return")}
          disabled={loading}
          className="flex items-center gap-2 px-4 py-2.5 rounded-lg text-sm font-semibold bg-amber-500/10 border border-amber-500/30 text-amber-400 hover:bg-amber-500/20 hover:border-amber-400/50 disabled:opacity-40 disabled:cursor-not-allowed transition active:scale-95 select-none"
        >
          <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6" />
          </svg>
          Return
        </button>

        <button
          onClick={() => setShowReviewModal(true)}
          disabled={loading}
          className="flex items-center gap-2 px-4 py-2.5 rounded-lg text-sm font-semibold bg-indigo-500/10 border border-indigo-500/30 text-indigo-400 hover:bg-indigo-500/20 hover:border-indigo-400/50 disabled:opacity-40 disabled:cursor-not-allowed transition active:scale-95 select-none"
        >
          <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
          </svg>
          Review
        </button>
      </div>

      {confirmAction && (
        <ConfirmModal
          action={confirmAction}
          onConfirm={handleConfirm}
          onCancel={() => setConfirmAction(null)}
          loading={loading}
        />
      )}

      {showReviewModal && (
        <ReviewModal
          onSubmit={handleReview}
          onCancel={() => setShowReviewModal(false)}
          loading={loading}
          reviewText={reviewText}
          setReviewText={setReviewText}
          pictures={pictures}
          setPictures={setPictures}
        />
      )}
    </div>
  )
}

export default TicketWorkflowAction