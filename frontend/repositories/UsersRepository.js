const resource = "/api";

export default ($axios) => ({
  create_doctor(payload) {
    return $axios.post(`${resource}/registerDoctor`, payload);
  },

  edit_doctor(params, payload) {
    return $axios.post(`${resource}/updateDoctor`, params, payload);
  },

  get_doctor(params) {
    return $axios.get(`${resource}/getDoctor`, { params });
  },

  get_doctors(params) {
    return $axios.get(`${resource}/getDoctors`, params);
  },

  create_patient(payload) {
    console.log("repository");
    return $axios.post(`${resource}/registerPatient`, payload);
  },

  edit_patient(id, payload) {
    console.log(id, payload, "repo");
    return $axios.post(`${resource}/updatePatient?id=${id}`, payload);
  },

  get_patient(params) {
    return $axios.get(`${resource}/getPatient`, { params });
  },

  get_patients(params) {
    return $axios.get(`${resource}/getPatients`, params);
  },

  login_admin(payload) {
    return $axios.post(`${resource}/login`, payload);
  },

  get_current_user(payload) {
    return $axios.post(`${resource}/`, payload);
  },

  get_departments() {
    return $axios.get(`${resource}/getDepartments`);
  },

  get_appointments() {
    return $axios.get(`${resource}/getAppointments`);
  },

  get_specializations() {
    return $axios.get(`${resource}/getSpecializations`);
  },

  make_appointment(paylaod) {
    return $axios.post(`${resource}/makeAppointment`, paylaod);
  },
});
