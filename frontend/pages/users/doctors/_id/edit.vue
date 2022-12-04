<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Edit Doctor</h1>
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
          v-model="form.contact_number"
          class="custom-input"
          label="Contact Number"
          placeholder="Enter Contact Number"
          :validation="$v.form.contact_number"
        />
        <custom-input
          v-model="form.number_of_patients"
          class="custom-input"
          label="Number of Patients"
          placeholder="Enter Number of Patients"
          :validation="$v.form.number_of_patients"
        />
        <custom-input
          v-model="form.address"
          class="custom-input"
          label="Address"
          placeholder="Enter Address"
          :validation="$v.form.address"
        />
        <custom-input
          v-model="form.department_id"
          class="custom-input"
          label="Department ID"
          placeholder="Enter Department ID"
          :validation="$v.form.department_id"
        />
        <custom-input
          v-model="form.specialization_details_id"
          class="custom-input"
          label="Specialization Details ID"
          placeholder="Enter Specialization Details ID"
          :validation="$v.form.specialization_details_id"
        />
        <custom-input
          v-model="form.experience_in_years"
          class="custom-input"
          label="Experience in Years"
          placeholder="Enter experience in Years"
          :validation="$v.form.experience_in_years"
        />
        <custom-input
          v-model="form.schedule_detaile"
          class="custom-input"
          label="Schedule Detaile"
          placeholder="Enter schedule detaile"
          :validation="$v.form.schedule_detaile"
        />
        <custom-input
          v-model="form.appointment_price"
          class="custom-input"
          label="Appointment Price"
          placeholder="Enter appointment price"
          :validation="$v.form.appointment_price"
        />
        <custom-select
          class="custom-input"
          label="Category"
          placeholder="Select Category"
          v-model="form.category"
          :options="options.category"
          :allow-empty="true"
          :validation="$v.form.category"
        />
        <custom-select
          class="custom-input"
          label="Degreee"
          placeholder="Select Degree"
          v-model="form.education_degree"
          :options="options.education_degree"
          :allow-empty="true"
          :validation="$v.form.education_degree"
        />
        <custom-select
          class="custom-input"
          label="Rating"
          placeholder="Select Rating"
          v-model="form.rating"
          :options="options.rating"
          :allow-empty="true"
          :validation="$v.form.rating"
        />
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
import CustomInput from "../../../../components/ui/CustomInput.vue";
import CustomSelect from "../../../../components/ui/CustomSelect.vue";
import CustomFileInput from "../../../../components/ui/CustomFileInput.vue";
import CustomDateInput from "../../../../components/ui/CustomDateInput.vue";
import options from "~/helpers/options";
import { num, positiveNum, phoneNum } from "~/helpers/validators";

export default {
  middleware: ["loggedin", "admin"],
  components: { CustomInput, CustomSelect, CustomFileInput, CustomDateInput },
  name: "signup",
  data() {
    return {
      options,
      form: {
        role: "doctor",
        date_of_birth: null,
        iin: null,
        contact_number: null,
        number_of_patients: null,
        address: null,
        department_id: null,
        specialization_details_id: null,
        appointment_price: null,
        experience_in_years: null,
        schedule_detaile: null,
        category: null,
        education_degree: null,
        rating: null,
        photolocation: "adasd",
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
        contact_number: {
          required,
          phoneNum,
        },
        number_of_patients: {
          required,
        },
        address: {
          required,
        },
        department_id: {
          required,
          num,
          positiveNum,
        },
        appointment_price: {
          required,
        },
        specialization_details_id: {
          required,
          num,
          positiveNum,
        },
        experience_in_years: {
          required,
          num,
          positiveNum,
        },
        schedule_detaile: {
          required,
        },
        category: {
          required,
        },
        education_degree: {
          required,
        },
        rating: {
          required,
        },
      },
    };
  },
  methods: {
    ...mapActions({
      createDoctor: "users/create_doctor",
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
      await this.createDoctor({
        id: this.$route.params.id,
        payload: this.form,
      });
    },
  },
};
</script>
