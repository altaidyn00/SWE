const resource = "/api";

export default ($axios) => ({
  create_doctor(payload) {
    return $axios.post(`${resource}/registerDoctor`, payload);
  },

  edit_doctor(params, payload) {
    return $axios.post(`${resource}/modifyDoctorInfo`, params, payload);
  },

  get_doctor(params) {
    return $axios.get(`${resource}/viewDoctorDetails`, params);
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
    return $axios.post(`${resource}/modifyPatientInfo?id=${id}`, payload);
  },

  get_patient(params) {
    return $axios.get(`${resource}/viewPatientDetails`, { params });
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
});
