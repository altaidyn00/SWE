<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Patient {{ patient.id }}</h1>
      <div
        class="d-flex patient-info flex-column mt-4 w-50 mx-auto justify-content-center align-items-left"
      >
        <div class="d-flex flex-row">
          <div class="mr-2">Birth Date:</div>
          <div class="font-weight-bold">{{ patient.dateofbirth }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">IIN:</div>
          <div class="font-weight-bold">{{ patient.iin }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Fullname:</div>
          <div class="font-weight-bold">{{ patient.fullname }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Email:</div>
          <div class="font-weight-bold">{{ patient.email }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Address:</div>
          <div class="font-weight-bold">{{ patient.address }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Contact number:</div>
          <div class="font-weight-bold">{{ patient.contactnumber }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Emergency number:</div>
          <div class="font-weight-bold">{{ patient.emergencynumber }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Emergency number:</div>
          <div class="font-weight-bold">{{ patient.emergencynumber }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Marital status:</div>
          <div class="font-weight-bold">{{ patient.maritalstatus }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Blood group:</div>
          <div class="font-weight-bold">{{ patient.bloodgroup }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Registration date:</div>
          <div class="font-weight-bold">{{ patient.registrationdate }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">ID:</div>
          <div class="font-weight-bold">{{ patient.id }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  middleware: ["loggedin", "admin"],
  name: "patients",
  async asyncData({ store, params }) {
    const query = {};
    query.id = params.id;
    await store.dispatch("users/get_patient", query);
  },
  computed: {
    ...mapGetters({
      patient: "users/patient",
    }),
  },
};
</script>
