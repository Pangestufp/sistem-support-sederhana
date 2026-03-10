import { useEffect, useState } from "react"
import DashboardService from "./dashboard.service"

export default function useDashboard() {
  const [jobs, setJobs] = useState([])
  const [loading, setLoading] = useState(false)
  const [weather, setWeather] = useState(null)

  const fetchJobs = async () => {
    setLoading(true)
    try {
      const res = await DashboardService.getJobs()
      setJobs(res.data)
    } catch (err) {
      console.error(err)
    } finally {
      setLoading(false)
    }
  }

  const fetchWeather = () => {
    if (!navigator.geolocation) return
    navigator.geolocation.getCurrentPosition(async (pos) => {
      try {
        const data = await DashboardService.getWeather(
          pos.coords.latitude,
          pos.coords.longitude
        )
        setWeather(data.current)
      } catch (err) {
        console.error(err)
      }
    })
  }

  useEffect(() => {
    fetchJobs()
    fetchWeather()
  }, [])

  return { jobs, loading, totalJobs: jobs.length, weather }
}