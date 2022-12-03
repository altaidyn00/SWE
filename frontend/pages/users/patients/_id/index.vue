<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Patient {{ patient.user.id }}</h1>
      <div
        class="d-flex patient-info flex-column mt-4 w-50 mx-auto justify-content-center align-items-left"
      >
        <div class="d-flex flex-row">
          <div class="mr-2">Birth Date:</div>
          <div class="font-weight-bold">
            {{ patient.patient.date_of_birth }}
          </div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">IIN:</div>
          <div class="font-weight-bold">{{ patient.patient.iin }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">First name:</div>
          <div class="font-weight-bold">{{ patient.user.first_name }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Last name:</div>
          <div class="font-weight-bold">{{ patient.user.last_name }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Email:</div>
          <div class="font-weight-bold">{{ patient.user.email }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Address:</div>
          <div class="font-weight-bold">{{ patient.patient.address }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Contact number:</div>
          <div class="font-weight-bold">
            {{ patient.patient.contact_number }}
          </div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Emergency number:</div>
          <div class="font-weight-bold">
            {{ patient.patient.emergency_contact_number }}
          </div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Address:</div>
          <div class="font-weight-bold">{{ patient.patient.address }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Marital status:</div>
          <div class="font-weight-bold">
            {{ patient.patient.marital_status }}
          </div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Blood group:</div>
          <div class="font-weight-bold">{{ patient.patient.blood_group }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Registration date:</div>
          <div class="font-weight-bold">
            {{ patient.patient.registration_date }}
          </div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">ID:</div>
          <div class="font-weight-bold">{{ patient.patient.id }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Role:</div>
          <div class="font-weight-bold">{{ patient.user.role }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  // middleware: ["loggedin", "admin"],
  name: "patients",
  async asyncData({ store, params }) {
    const query = {};
    console.log(params);
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
