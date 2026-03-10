import { Link } from "react-router-dom";
import useDashboard from "./useDashboard";

function Dashboard() {
  const { jobs, loading, totalJobs, weather } = useDashboard();

  const wLabel  = (code) => {
    if (code === 0) return "☀️ Clear"
    if (code <= 3) return "⛅ Cloudy"
    if (code <= 67) return "🌧️ Rain"
    if (code <= 77) return "❄️ Snow"
    if (code <= 99) return "⛈️ Storm"
    return "🌤️"
  }

  return (
    <div className="text-white p-6 space-y-6">
      <div>
        <p className="text-xs font-semibold tracking-widest text-indigo-400 uppercase mb-4">Overview</p>
        <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">

          <div className="bg-slate-900 border border-slate-700/60 rounded-2xl p-6">
            <p className="text-xs text-slate-500 uppercase tracking-wider mb-3">My Jobs</p>
            <p className="text-5xl font-black text-indigo-400">
              {loading ? "..." : totalJobs}
            </p>
            <p className="text-xs text-slate-500 mt-2">tasks assigned to you</p>
          </div>

          {weather && wLabel ? (
            <div className="sm:col-span-2 bg-slate-900 border border-slate-700/60 rounded-2xl p-6 relative overflow-hidden">
              <div className="absolute -top-6 -right-6 w-32 h-32 bg-indigo-600/10 rounded-full blur-2xl pointer-events-none" />

              <p className="text-xs text-slate-500 uppercase tracking-wider mb-4">Weather · Your Location</p>

              <div className="flex items-end justify-between">
                <div>
                  <div className="flex items-end gap-3 mb-2">
                    <span className="text-6xl font-black text-white">{weather.temperature_2m}°</span>
                    <span className="text-2xl text-slate-400 mb-2">C</span>
                  </div>
                  <p className="text-base text-slate-300 font-medium">{wLabel.icon} {wLabel.label}</p>
                </div>

                <div className="text-right space-y-2">
                  <div className="bg-slate-800/60 border border-slate-700/40 rounded-xl px-4 py-2">
                    <p className="text-xs text-slate-500 mb-0.5">Wind</p>
                    <p className="text-sm font-semibold text-white">💨 {weather.windspeed_10m} km/h</p>
                  </div>
                </div>
              </div>
            </div>
          ) : (
            <div className="sm:col-span-2 bg-slate-900 border border-slate-700/60 rounded-2xl p-6 flex items-center justify-center text-slate-600 text-sm">
              Loading weather...
            </div>
          )}

        </div>
      </div>
      <div>
        <p className="text-xs font-semibold tracking-widest text-indigo-400 uppercase mb-4">My Tasks</p>
        <div className="bg-slate-900 border border-slate-700/60 rounded-2xl overflow-hidden">
          <table className="w-full text-sm">
            <thead>
              <tr className="border-b border-slate-700/60">
                <th className="px-5 py-3.5 text-left text-xs font-semibold text-slate-400 uppercase tracking-wider">Doc No</th>
                <th className="px-5 py-3.5 text-left text-xs font-semibold text-slate-400 uppercase tracking-wider">Activity</th>
                <th className="px-5 py-3.5 text-left text-xs font-semibold text-slate-400 uppercase tracking-wider">Action</th>
              </tr>
            </thead>
            <tbody>
              {jobs.map((job) => (
                <tr key={job.TicketID} className="border-t border-slate-800 hover:bg-slate-800/40 transition">
                  <td className="px-5 py-3.5 font-medium text-white">{job.DocNo}</td>
                  <td className="px-5 py-3.5">
                    <span className="text-xs font-semibold px-2.5 py-1 rounded-full bg-indigo-500/10 border border-indigo-500/20 text-indigo-300 capitalize">
                      {job.Activity}
                    </span>
                  </td>
                  <td className="px-5 py-3.5">
                    <Link
                      to={`/ticket/${job.TicketID}`}
                      className="flex items-center gap-1.5 text-xs font-semibold text-indigo-400 hover:text-indigo-300 transition w-fit"
                    >
                      Open Ticket
                      <svg className="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                      </svg>
                    </Link>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>

          {jobs.length === 0 && !loading && (
            <div className="py-16 text-slate-600 text-center">
              <svg className="w-8 h-8 mx-auto mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
              </svg>
              <p className="text-sm">No jobs assigned</p>
            </div>
          )}
        </div>
      </div>

    </div>
  );
}

export default Dashboard;
