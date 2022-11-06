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
      this._vm.$bvToast.toast("Doctor successfully created!", {
        title: "Doctor",
        toaster: "b-toaster-bottom-left",
        variant: "success",
        solid: true,
      });
      return response.data;
    } catch (error) {
      this._vm.$bvToast.toast("Error occured during doctor creating.", {
        title: "Doctor",
        toaster: "b-toaster-bottom-left",
        variant: "danger",
        solid: true,
      });
    }
  },

  async create_patient(_, payload) {
    try {
      console.log("sadasd");
      const response = await this.$repositories.users.create_patient(payload);
      this._vm.$bvToast.toast("Patient successfully created!", {
        title: "Patient",
        toaster: "b-toaster-bottom-left",
        variant: "success",
        solid: true,
      });
      return response.data;
    } catch (error) {
      this._vm.$bvToast.toast("Error occured during patient creating.", {
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
      this._vm.$bvToast.toast("Doctor successfully edited!", {
        title: "Doctor",
        toaster: "b-toaster-bottom-left",
        variant: "success",
        solid: true,
      });
      return response.data;
    } catch (error) {
      console.log(error);
      this._vm.$bvToast.toast("Error occured during doctor editing.", {
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
      return response.data;
    } catch (error) {
    }
  },

  async get_doctor({ commit }, id) {
    try {
      const response = await this.$repositories.users.get_doctor(id);
      const doctor = response.data;
      commit("SET_DOCTOR", doctor);
    } catch (error) {}
  },

  async get_patient({ commit }, id) {
    try {
      const response = await this.$repositories.users.get_patient(id);
      const patient = response.data;
      commit("SET_PATIENT", patient);
    } catch (error) {}
  },

  async get_doctor({ commit }, params) {
    try {
      const response = await this.$repositories.users.get_doctors(params);
      const doctors = response.data;
      commit("SET_DOCTORS", doctors);
    } catch (error) {}
  },

  async get_patients({ commit }, params) {
    try {
      const response = await this.$repositories.users.get_patients(params);
      console.log(response.data)
      const patients = response.data;
      commit("SET_PATIENTS", patients);
    } catch (error) {}
  },

  async login({ commit }, payload) {
    try {
      const response = await this.$repositories.users.login_admin(payload);
      console.log(response);
    } catch (error) {}
  },
};
