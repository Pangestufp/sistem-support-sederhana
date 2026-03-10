class Endpoints {
  static AUTH = {
    LOGIN: "/api/login",
    REGISTER: "/api/register"
  }

 static TICKET_TYPE = {
    BASE: "/api/ticket-type",
    GET_ALL: "/api/ticket-type",
    GET_BY_ID: (id) => `/api/ticket-type/${id}`,
  }

  static TICKET = {
    BASE: "/api/ticket",
    CREATE: "/api/ticket",
    GET_BY_ID: (id) => `/api/ticket/${id}`,
    GET_DETAILS: (id) => `/api/ticket/details/${id}`,
    GET_ATTACHMENTS: (id) => `/api/ticket/attachments/${id}`,
    GET_WORKFLOW: (id) => `/api/ticket/workflow/${id}`,
    APPROVE: (id) => `/api/ticket/approve/${id}`,
    REJECT: (id) => `/api/ticket/reject/${id}`,
    RETURN: (id) => `/api/ticket/return/${id}`,
    REVIEW: (id) => `/api/ticket/review/${id}`,
    GET_ALL: (id) => `/api/ticket/getAll?last_id=${id}`,
    JOB: "api/ticket/allJob",
  }
  


}

export default Endpoints