<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Create Doctor</h1>
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
        <custom-input
          v-model="form.email"
          class="custom-input"
          label="Email (Optional)"
          placeholder="Enter Email"
          :validation="$v.form.email"
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
        <custom-select
          class="custom-input"
          label="Specialization Department ID"
          placeholder="Select Department ID"
          v-model="form.department_id"
          :options="departments"
          :allow-empty="true"
          track-by="description"
          :custom-label="(option) => option.description"
          :validation="$v.form.department_id"
        />
        <custom-select
          class="custom-input"
          label="Specialization Details ID"
          placeholder="Select Specialization Details ID"
          v-model="form.specialization_id"
          :options="specializations"
          :allow-empty="true"
          track-by="description"
          :custom-label="(option) => option.description"
          :validation="$v.form.specialization_id"
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
        <custom-file-input
          label="Photo"
          :value="form.photo"
          class="custom-input"
          :validation="$v.form.photo"
          button="Upload photo"
          @change="setPhoto"
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
import CustomInput from "../../../components/ui/CustomInput.vue";
import CustomSelect from "../../../components/ui/CustomSelect.vue";
import CustomFileInput from "../../../components/ui/CustomFileInput.vue";
import CustomDateInput from "../../../components/ui/CustomDateInput.vue";
import options from "~/helpers/options";
import { num, positiveNum, phoneNum } from "~/helpers/validators";
import objectToFormData from "~/helpers/objectToFormData";

export default {
  middleware: ["loggedin", "admin"],
  components: { CustomInput, CustomSelect, CustomFileInput, CustomDateInput },
  name: "signup",
  async asyncData({ store }) {
    const departments = await store.dispatch("users/get_departments");
    const specializations = await store.dispatch("users/get_specializations");
    return { departments, specializations };
  },
  data() {
    return {
      options,
      form: {
        role: "doctor",
        date_of_birth: null,
        iin: null,
        id: null,
        contact_number: null,
        number_of_patients: null,
        address: null,
        department_id: null,
        specialization_id: null,
        appointment_price: null,
        experience_in_years: null,
        schedule_detaile: null,
        category: null,
        education_degree: null,
        rating: null,
        email: null,
        first_name: null,
        last_name: null,
        password: null,
        photo: null,
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
        appointment_price: {
          required,
        },
        specialization_id: {
          required,
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
        photo: {
          required,
        },
      },
    };
  },
  methods: {
    objectToFormData,
    ...mapActions({
      createDoctor: "users/create_doctor",
    }),
    setPhoto(file) {
      this.form.photo = file;
    },
    async create() {
      this.$v.form.$touch();
      if (this.$v.form.$invalid) return;
      const user = {};
      user.id = +this.form.id;
      user.email = this.form.email;
      user.role = this.form.role;
      user.first_name = this.form.first_name;
      user.last_name = this.form.last_name;
      user.password = this.form.password;
      const doc = {};
      doc.date_of_birth = this.form.date_of_birth;
      doc.iin = this.form.iin;
      doc.id = +this.form.id;
      doc.contact_number = this.form.contact_number;
      doc.number_of_patients = +this.form.number_of_patients;
      doc.category = this.form.category;
      doc.department_id = +this.form.department_id.id;
      doc.specialization_id = +this.form.specialization_id.id;
      doc.experience_in_years = +this.form.experience_in_years;
      doc.appointment_price = +this.form.appointment_price;
      doc.education_degree = this.form.education_degree;
      doc.schedule_detaile = this.form.schedule_detaile;
      doc.rating = +this.form.rating;
      doc.address = this.form.address;
      doc.photo = this.form.photo;
      const doctor = this.objectToFormData(doc);
      const response = await this.createDoctor({ user, doctor });
      console.log(response);
    },
  },
};
</script>
