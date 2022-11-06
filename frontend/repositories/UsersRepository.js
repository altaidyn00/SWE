const resource = "/api";

export default ($axios) => ({
  create_doctor(payload) {
    return $axios.post(`registerDoctor`, payload);
  },

  edit_doctor(id, payload) {
    return $axios.post(`modifyDoctorInfo`, payload);
  },

  get_doctor(id) {
    return $axios.get(`viewDoctorDetails`, id);
  },

  get_doctors(params) {
    return $axios.get(`getDoctors`, params);
  },

  create_patient(payload) {
    console.log("repository");
    return $axios.post(`${resource}/registerPatient`, payload);
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
  login_admin(payload) {
    return $axios.post(`${resource}/login`, payload);
  },
});
