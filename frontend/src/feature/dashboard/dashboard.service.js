import axios from "axios"
import api from "../../app/axios"
import Endpoints from "../../shared/endpoints"

const DashboardService = {
  getJobs: async () => {
    const res = await api.get(Endpoints.TICKET.JOB)
    return res.data
  },

  getWeather: async (lat, lon) => {
    const res = await axios.get("https://api.open-meteo.com/v1/forecast", {
      params: {
        latitude: lat,
        longitude: lon,
        current: "temperature_2m,weathercode,windspeed_10m",
        timezone: "auto"
      }
    })
    return res.data
  }
}

export default DashboardService