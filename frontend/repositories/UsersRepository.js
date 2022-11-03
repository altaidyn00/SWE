const resource = "/api/v0/user";

export default ($axios) => ({
  create_doctor(payload) {
    return $axios.post(`${resource}/create/doctor`, payload);
  },

  edit_doctor(id, payload) {
    return $axios.post(`${resource}/${id}/edit/doctor`, payload);
  },

  get_doctor(id) {
    return $axios.get(`${resource}/${id}/doctor`);
  },

  get_doctors(params) {
    return $axios.get(`${resource}/doctors`, params);
  },

  create_patient(payload) {
    return $axios.post(`${resource}/create/patient`, payload);
  },

  edit_patient(id, payload) {
    return $axios.post(`${resource}/${id}/edit/patient`, payload);
  },

  get_patient(id) {
    return $axios.get(`${resource}/${id}/patient`);
  },

  get_patients(params) {
    return $axios.get(`${resource}/patients`, params);
  },
});
