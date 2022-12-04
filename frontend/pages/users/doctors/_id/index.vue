<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Doctor {{ doctor.user.id }}</h1>
      <div
        class="d-flex doctor-info flex-column mt-4 w-50 mx-auto justify-content-center align-items-left"
      >
        <div class="d-flex flex-row">
          <div class="mr-2">Birth Date:</div>
          <div class="font-weight-bold">
            {{ doctor.doctor.date_of_birth }}
          </div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">IIN:</div>
          <div class="font-weight-bold">{{ doctor.doctor.iin }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">First name:</div>
          <div class="font-weight-bold">{{ doctor.user.first_name }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Last name:</div>
          <div class="font-weight-bold">{{ doctor.user.last_name }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Email:</div>
          <div class="font-weight-bold">{{ doctor.user.email }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Category:</div>
          <div class="font-weight-bold">{{ doctor.doctor.category }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Contact number:</div>
          <div class="font-weight-bold">
            {{ doctor.doctor.contact_number }}
          </div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Education degree:</div>
          <div class="font-weight-bold">
            {{ doctor.doctor.education_degree }}
          </div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Address:</div>
          <div class="font-weight-bold">{{ doctor.doctor.address }}</div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Appointment price:</div>
          <div class="font-weight-bold">
            {{ doctor.doctor.appointment_price }}KZT
          </div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Department id:</div>
          <div class="font-weight-bold">
            {{ doctor.doctor.department_id }}
          </div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Specialization details id:</div>
          <div class="font-weight-bold">
            {{ doctor.doctor.specialization_details_id }}
          </div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Number of patients:</div>
          <div class="font-weight-bold">
            {{ doctor.doctor.number_of_patients }}
          </div>
        </div>
        <div class="d-flex flex-row">
          <div class="mr-2">Role:</div>
          <div class="font-weight-bold">{{ doctor.user.role }}</div>
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
    await store.dispatch("users/get_doctor", query);
  },
  computed: {
    ...mapGetters({
      doctor: "users/doctor",
    }),
  },
};
</script>
