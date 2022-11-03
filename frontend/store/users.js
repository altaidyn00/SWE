export const state = () => ({
  doctor: null,
  doctors: [],
  patient: null,
  patients: [],
});

export const getters = {
  doctor(state) {
    return state.doctor;
  },
  doctors(state) {
    return state.doctors;
  },
  patient(state) {
    return state.patient;
  },
  patients(state) {
    return state.patients;
  },
};

export const mutations = {
  SET_DOCTOR(state, doctor) {
    state.doctor = doctor;
  },
  SET_DOCTORS(state, doctors) {
    state.doctors = doctors;
  },
  SET_PATIENT(state, patient) {
    state.patient = patient;
  },
  SET_PATIENTS(state, patients) {
    state.patients = patients;
  },
};

export const actions = {
  async create_doctor(_, payload) {
    try {
      const response = await this.$repositories.users.create_doctor(payload);
      this.$bvToast.toast("Doctor successfully created!", {
        title: "Doctor",
        toaster: "b-toaster-bottom-left",
        variant: "success",
        solid: true,
      });
      return response.data;
    } catch (error) {
      this.$bvToast.toast("Error occured during doctor creating.", {
        title: "Doctor",
        toaster: "b-toaster-bottom-left",
        variant: "danger",
        solid: true,
      });
    }
  },

  async create_patient(_, payload) {
    try {
      const response = await this.$repositories.users.create_patient(payload);
      this.$bvToast.toast("Patient successfully created!", {
        title: "Patient",
        toaster: "b-toaster-bottom-left",
        variant: "success",
        solid: true,
      });
      return response.data;
    } catch (error) {
      this.$bvToast.toast("Error occured during patient creating.", {
        title: "Patient",
        toaster: "b-toaster-bottom-left",
        variant: "danger",
        solid: true,
      });
    }
  },

  async edit_doctor(_, { id, payload }) {
    try {
      const response = await this.$repositories.users.edit_doctor(id, payload);
      this.$bvToast.toast("Doctor successfully updated!", {
        title: "Doctor",
        toaster: "b-toaster-bottom-left",
        variant: "success",
        solid: true,
      });
      return response.data;
    } catch (error) {
      this.$bvToast.toast("Error occured during doctor updating.", {
        title: "Doctor",
        toaster: "b-toaster-bottom-left",
        variant: "danger",
        solid: true,
      });
    }
  },

  async edit_patient(_, { id, payload }) {
    try {
      const response = await this.$repositories.users.edit_patient(id, payload);
      this.$bvToast.toast("Patient successfully updated!", {
        title: "Patient",
        toaster: "b-toaster-bottom-left",
        variant: "success",
        solid: true,
      });
      return response.data;
    } catch (error) {
      this.$bvToast.toast("Error occured during patient updating.", {
        title: "Patient",
        toaster: "b-toaster-bottom-left",
        variant: "danger",
        solid: true,
      });
    }
  },

  async get_doctor({ commit }, id) {
    try {
      const response = await this.$repositories.suers.get_doctor(id);
      const doctor = response.data;
      commit("SET_DOCTOR", doctor);
    } catch (error) {}
  },

  async get_patient({ commit }, id) {
    try {
      const response = await this.$repositories.suers.get_patient(id);
      const patient = response.data;
      commit("SET_PATIENT", patient);
    } catch (error) {}
  },

  async get_doctor({ commit }, params) {
    try {
      const response = await this.$repositories.suers.get_doctors(params);
      const doctors = response.data;
      commit("SET_DOCTORS", doctors);
    } catch (error) {}
  },

  async get_patients({ commit }, params) {
    try {
      const response = await this.$repositories.suers.get_patients(params);
      const patients = response.data;
      commit("SET_PATIENTS", patients);
    } catch (error) {}
  },
};
