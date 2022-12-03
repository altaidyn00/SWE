<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Create Patient</h1>
      <div class="d-flex flex-wrap justify-content-between my-4">
        <custom-date-input
          v-model="form.date_of_birth"
          class="custom-input"
          label="Date of Birth"
          :validation="$v.form.date_of_birth"
        />
        <custom-input
          v-model="form.iin"
          class="custom-input"
          label="IIN Number"
          placeholder="Enter IIN Number"
          :validation="$v.form.iin"
        />
        <custom-input
          v-model="form.id"
          class="custom-input"
          label="ID Number"
          placeholder="Enter ID Number"
          :validation="$v.form.id"
        />
        <custom-input
          v-model="form.first_name"
          class="custom-input"
          label="First Name"
          placeholder="Enter First Name"
          :validation="$v.form.first_name"
        />
        <custom-input
          v-model="form.last_name"
          class="custom-input"
          label="Last Name"
          placeholder="Enter Last Name"
          :validation="$v.form.last_name"
        />
        <custom-input
          v-model="form.password"
          class="custom-input"
          label="Password"
          placeholder="Enter Password"
          :validation="$v.form.password"
        />
        <custom-select
          class="custom-input"
          label="Blood Group"
          placeholder="Select Blood Group"
          v-model="form.blood_group"
          :options="options.blood_group"
          :allow-empty="true"
          :validation="$v.form.blood_group"
        />
        <!-- <custom-input
          v-model="form.blood_group"
          class="custom-input"
          label="Blood Group"
          placeholder="Enter Blood Group"
          :validation="$v.form.blood_group"
        /> -->
        <custom-input
          v-model="form.emergency_contact_number"
          class="custom-input"
          label="Emergency Contact"
          placeholder="Enter Emergency Contact Number"
          :validation="$v.form.emergency_contact_number"
        />
        <custom-input
          v-model="form.contact_number"
          class="custom-input"
          label="Contact Number"
          placeholder="Enter Contact Number"
          :validation="$v.form.contact_number"
        />
        <custom-input
          v-model="form.email"
          class="custom-input"
          label="Email (Optional)"
          placeholder="Enter Email"
          :validation="$v.form.email"
        />
        <custom-input
          v-model="form.address"
          class="custom-input"
          label="Address"
          placeholder="Enter Address"
          :validation="$v.form.address"
        />
        <custom-select
          class="custom-input"
          label="Marital Status"
          placeholder="Select Marital Status"
          v-model="form.marital_status"
          :options="options.marital_status"
          :allow-empty="true"
          :validation="$v.form.marital_status"
        />
        <!-- <custom-input
          v-model="form.marital_status"
          class="custom-input"
          label="Marital Status"
          placeholder="Enter Marital Status"
          :validation="$v.form.marital_status"
        /> -->
      </div>
      <b-button size="md" class="button my-2 ml-sm-0" @click="create"
        >create</b-button
      >
    </div>
  </div>
</template>

<script>
import { required, email } from "vuelidate/lib/validators";
import { mapActions } from "vuex";
import CustomInput from "../../../components/ui/CustomInput.vue";
import CustomSelect from "../../../components/ui/CustomSelect.vue";
import CustomFileInput from "../../../components/ui/CustomFileInput.vue";
import CustomDateInput from "../../../components/ui/CustomDateInput.vue";
import options from "~/helpers/options";
import { num, positiveNum, phoneNum } from "~/helpers/validators";

export default {
  components: { CustomInput, CustomSelect, CustomFileInput, CustomDateInput },
  name: "signup",
  data() {
    return {
      options,
      form: {
        date_of_birth: null,
        iin: null,
        id: null,
        first_name: null,
        last_name: null,
        blood_group: null,
        emergency_contact_number: null,
        contact_number: null,
        email: null,
        address: null,
        marital_status: null,
        registration_date: null,
        role: "patient",
        password: null,
      },
    };
  },
  validations() {
    return {
      form: {
        date_of_birth: {
          required,
        },
        iin: {
          required,
          num,
          positiveNum,
        },
        id: {
          required,
          num,
          positiveNum,
        },
        first_name: {
          required,
        },
        last_name: {
          required,
        },
        password: {
          required,
        },
        email: {
          required,
          email,
        },
        blood_group: {
          required,
        },
        emergency_contact_number: {
          required,
          phoneNum,
        },
        contact_number: {
          required,
          phoneNum,
        },
        address: {
          required,
        },
        marital_status: {
          required,
        },
      },
    };
  },
  methods: {
    ...mapActions({
      createPatient: "users/create_patient",
    }),
    async create() {
      this.$v.form.$touch();
      if (this.$v.form.$invalid) return;
      var today = new Date();
      var dd = String(today.getDate()).padStart(2, "0");
      var mm = String(today.getMonth() + 1).padStart(2, "0"); //January is 0!
      var yyyy = today.getFullYear();
      today = yyyy + "-" + mm + "-" + dd;
      this.form.registration_date = today;
      this.form.marital_status = String(this.form.marital_status);
      this.form.blood_group = String(this.form.blood_group);
      const user = {};
      user.id = +this.form.id;
      user.email = this.form.email;
      user.role = this.form.role;
      user.first_name = this.form.first_name;
      user.last_name = this.form.last_name;
      user.password = this.form.password;
      const patient = {};
      patient.date_of_birth = this.form.date_of_birth;
      patient.iin = this.form.iin;
      patient.id = +this.form.id;
      patient.blood_group = this.form.blood_group;
      patient.marital_status = this.form.marital_status;
      patient.registration_date = this.form.registration_date;
      patient.contact_number = this.form.contact_number;
      patient.emergency_contact_number = this.form.emergency_contact_number;
      patient.address = this.form.address;
      const response = await this.createPatient({ user, patient });
      console.log(response);
    },
  },
};
</script>
