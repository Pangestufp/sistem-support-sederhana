import { useState } from "react";
import TicketService from "./ticket.service";

export function useTicketAction() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const handleAction = async (callback) => {
    setLoading(true);
    setError(null);

    try {
      const response = await callback();

      if (response.code !== 200) {
        throw new Error(response.message);
      }

      return true;
    } catch (err) {
      const message = err.response?.data?.message || err.message || "Action gagal";
      setError(message);
      return false;
    } finally {
      setLoading(false);
    }
  };

  const approve = (id) => handleAction(() => TicketService.approve(id));

  const reject = (id) => handleAction(() => TicketService.reject(id));

  const returnTicket = (id) =>
    handleAction(() => TicketService.returnTicket(id));

  const review = (id, payload) =>
    handleAction(() => TicketService.review(id, payload));

  return {
    approve,
    reject,
    returnTicket,
    review,
    loading,
    error,
  };
}
