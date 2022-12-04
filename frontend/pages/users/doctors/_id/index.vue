<template>
  <div class="custom-container">
    <div class="text-center d-flex flex-column justify-content-center">
      <h1>Doctor {{ doctor.user.id }}</h1>
      <div
        class="d-flex doctor-info flex-column mt-4 w-50 mx-auto justify-content-center align-items-center"
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
      <h1>Make appointment</h1>
      <custom-input
        v-model="doctor_id"
        class="custom-input"
        label="Doctor ID"
        placeholder="Enter Doctor ID"
        :validation="$v.doctor_id"
      />
      <custom-input
        v-model="preferred_date"
        class="custom-input"
        label="Date"
        placeholder="Enter date"
        :validation="$v.preferred_date"
      />
      <custom-input
        v-model="preferred_time"
        class="custom-input"
        label="Time"
        placeholder="Enter Time"
        :validation="$v.preferred_time"
      />
      <custom-input
        v-model="name"
        class="custom-input"
        label="Name"
        placeholder="Enter Name"
        :validation="$v.name"
      />
      <custom-input
        v-model="surname"
        class="custom-input"
        label="Surname"
        placeholder="Enter Surname"
        :validation="$v.surname"
      />
      <custom-input
        v-model="email"
        class="custom-input"
        label="Email"
        placeholder="Enter Email"
        :validation="$v.email"
      />
      <b-button size="md" class="button my-2 ml-sm-0" @click="make"
        >make</b-button
      >
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { required, email } from "vuelidate/lib/validators";
import CustomInput from "../../../../components/ui/CustomInput.vue";
export default {
  // middleware: ["loggedin", "admin"],
  components: { CustomInput },
  name: "patients",
  async asyncData({ store, params }) {
    const query = {};
    console.log(params);
    query.id = params.id;
    await store.dispatch("users/get_doctor", query);
  },
  data() {
    return {
      doctor_id: null,
      preferred_date: null,
      preferred_time: null,
      name: null,
      surname: null,
      email: null,
    };
  },
  validations() {
    return {
      doctor_id: {
        required,
      },
      preferred_date: {
        required,
      },
      preferred_time: {
        required,
      },
      name: {
        required,
      },
      surname: {
        required,
      },
      email: {
        required,
      },
    };
  },
  mounted() {
    this.doctor_id = +this.doctor.user.id;
    this.name = this.$auth.user.first_name;
    this.surname = this.$auth.user.last_name;
    this.email = this.$auth.user.email;
  },
  computed: {
    ...mapGetters({
      doctor: "users/doctor",
    }),
  },
  methods: {
    async make() {
      this.$v.$touch();
      if (this.$v.$invalid) return;
      const payload = {};
      payload.doctor_id = +this.doctor_id;
      payload.preferred_date = this.preferred_date;
      payload.preferred_time = this.preferred_time;
      payload.name = this.name;
      payload.surname = this.surname;
      payload.email = this.email;
      const response = await this.$store.dispatch(
        "users/make_appointment",
        payload
      );
    },
  },
};
</script>
