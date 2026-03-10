import useBrowse from "./useBrowse"

function Browse() {
  const { tickets, loading, nextPage } = useBrowse()

  return (
    <div>

      <h1 className="text-2xl font-bold mb-6">Ticket Browse</h1>

      <div className="bg-slate-900 border border-slate-700 rounded-xl overflow-hidden">

        <table className="w-full text-sm">
          <thead className="bg-slate-800 text-slate-400">
            <tr>
              <th className="px-4 py-3 text-left">Doc No</th>
              <th className="px-4 py-3 text-left">Description</th>
              <th className="px-4 py-3 text-left">Status</th>
              <th className="px-4 py-3 text-left">Created At</th>
            </tr>
          </thead>

          <tbody>
            {tickets.map((t) => (
              <tr
                key={t.id}
                className="border-t border-slate-800 hover:bg-slate-800/40"
              >
                <td className="px-4 py-3"><a href={`/ticket/${t.id}`} target="_blank">{t.doc_no}</a></td>
                <td className="px-4 py-3">{t.description}</td>
                <td className="px-4 py-3">{t.status}</td>
                <td className="px-4 py-3">
                  {new Date(t.created_at).toLocaleDateString("id-ID")}
                </td>
              </tr>
            ))}
          </tbody>
        </table>

      </div>

      <div className="mt-6 flex justify-center">

        <button
          onClick={nextPage}
          disabled={loading}
          className="px-6 py-2 bg-indigo-600 hover:bg-indigo-500 rounded-lg"
        >
          {loading ? "Loading..." : "Next"}
        </button>

      </div>

    </div>
  )
}

export default Browse