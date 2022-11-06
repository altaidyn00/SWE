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

  edit_patient(params, payload) {
    return $axios.post(`${resource}/edit/patient`, params, payload);
  },

  get_patient(params) {
    return $axios.get(`${resource}/getPatient`, params);
  },

  get_patients(params) {
    return $axios.get(`${resource}/getPatients`, params);
  },
  login_admin(payload) {
    return $axios.post(`${resource}/login`, payload);
  },
});
