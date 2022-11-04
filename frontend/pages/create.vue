<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Create User</h1>
      <div class="mt-4">
        <h4>Choose role</h4>
        <div class="d-flex justify-content-between mt-2">
          <div
            class="registration-tab"
            :class="{
              'registration-tab__active': role === 'patient',
            }"
            @click="role = 'patient'"
          >
            Patient
          </div>
          <div
            class="registration-tab"
            :class="{
              'registration-tab__active': role === 'doctor',
            }"
            @click="role = 'doctor'"
          >
            Doctor
          </div>
        </div>
      </div>
      <div class="d-flex flex-wrap justify-content-between my-4">
        <div class="custom-input">
          <span>Date of Birth</span>
          <b-form-datepicker v-model="form.birth_date"></b-form-datepicker>
        </div>
        <custom-input
          v-model="form.iin_number"
          class="custom-input"
          label="IIN Number"
          placeholder="Enter IIN Number"
          :validation="$v.form.iin_number"
        />
        <custom-input
          v-model="form.id_number"
          class="custom-input"
          label="ID Number"
          placeholder="Enter ID Number"
          :validation="$v.form.id_number"
        />
        <custom-input
          v-model="form.name"
          class="custom-input"
          label="Name"
          placeholder="Enter Name"
          :validation="$v.form.name"
        />
        <custom-input
          v-model="form.surname"
          class="custom-input"
          label="Surname"
          placeholder="Enter Surname"
          :validation="$v.form.surname"
        />
        <custom-input
          v-model="form.middlename"
          class="custom-input"
          label="Middlename"
          placeholder="Enter Middlename"
          :validation="$v.form.middlename"
        />
        <custom-select
          v-if="isPatient"
          class="custom-input"
          label="Blood Group"
          placeholder="Select Blood Group"
          v-model="form.blood_group"
          :options="options.blood_group"
          :allow-empty="true"
          :validation="$v.form.blood_group"
        />
        <custom-input
          v-if="isPatient"
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
          v-if="isPatient"
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
          v-if="isPatient"
          class="custom-input"
          label="Marital Status"
          placeholder="Select Marital Status"
          v-model="form.marital_status"
          :options="options.marital_status"
          :allow-empty="true"
          :validation="$v.form.marital_status"
        />
        <custom-input
          v-if="isDoctor"
          v-model="form.department_id"
          class="custom-input"
          label="Department ID"
          placeholder="Enter Department ID"
          :validation="$v.form.department_id"
        />
        <custom-input
          v-if="isDoctor"
          v-model="form.specialization_details_id"
          class="custom-input"
          label="Specialization Details ID"
          placeholder="Enter Specialization Details ID"
          :validation="$v.form.specialization_details_id"
        />
        <custom-input
          v-if="isDoctor"
          v-model="form.experience_in_years"
          class="custom-input"
          label="Experience in Years"
          placeholder="Enter Experience in Years"
          :validation="$v.form.experience_in_years"
        />
        <custom-file-input
          v-if="isDoctor"
          class="custom-input"
          label="Doctor's Photo"
          placeholder=" Upload Doctor's Photos"
          v-model="form.photo_of_doctor"
          :options="options.photo_of_doctor"
          :allow-empty="true"
          :validation="$v.form.photo_of_doctor"
        />
        <custom-select
          v-if="isDoctor"
          class="custom-input"
          label="Category"
          placeholder="Select Category"
          v-model="form.category"
          :options="options.category"
          :allow-empty="true"
          :validation="$v.form.category"
        />
        <custom-input
          v-if="isDoctor"
          v-model="form.price_of_the_appointment"
          class="custom-input"
          label="Price (KZT)"
          placeholder="Enter Price (KZT)"
          :validation="$v.form.price_of_the_appointment"
        />
        <custom-input
          v-if="isDoctor"
          v-model="form.schedule_details"
          class="custom-input"
          label="Schedule Details"
          placeholder="Enter Schedule Details"
          :validation="$v.form.schedule_details"
        />
        <custom-select
          v-if="isDoctor"
          class="custom-input"
          label="Degreee"
          placeholder="Select Degree"
          v-model="form.degree"
          :options="options.degree"
          :allow-empty="true"
          :validation="$v.form.degree"
        />
        <custom-select
          v-if="isDoctor"
          class="custom-input"
          label="Rating"
          placeholder="Select Rating"
          v-model="form.rating"
          :options="options.rating"
          :allow-empty="true"
          :validation="$v.form.rating"
        />
        <custom-input
          v-if="isDoctor"
          v-model="form.homepage_url"
          class="custom-input"
          label="Homepage Url (Optional)"
          placeholder="Enter Homepage Url"
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
import CustomInput from "../components/ui/CustomInput.vue";
import CustomSelect from "../components/ui/CustomSelect.vue";
import CustomFileInput from "../components/ui/CustomFileInput.vue";
import options from "~/helpers/options";
import { validateIf, num, positiveNum, phoneNum } from "~/helpers/validators";

export default {
  components: { CustomInput, CustomSelect, CustomFileInput },
  name: "signup",
  data() {
    return {
      options,
      role: "patient",
      form: {
        birth_date: null,
        iin_number: null,
        id_number: null,
        name: null,
        surname: null,
        middlename: null,
        blood_group: null,
        emergency_contact_number: null,
        contact_number: null,
        email: null,
        address: null,
        marital_status: null,
        registration_date: null,
        department_id: null,
        specialization_details_id: null,
        experience_in_years: null,
        photo_of_doctor: null,
        category: null,
        price_of_the_appointment: null,
        schedule_details: null,
        degree: null,
        rating: null,
        homepage_url: null,
      },
    };
  },
  validations() {
    return {
      form: {
        birth_date: {
          required,
        },
        iin_number: {
          required,
          num,
          positiveNum,
        },
        id_number: {
          required,
          num,
          positiveNum,
        },
        name: {
          required,
        },
        surname: {
          required,
        },
        middlename: {
          required,
        },
        email: {
          ...validateIf(this.isPatient, {
            email,
          }),
        },
        blood_group: {
          ...validateIf(this.isPatient, {
            required,
          }),
        },
        emergency_contact_number: {
          ...validateIf(this.isPatient, {
            required,
            phoneNum
          }),
        },
        contact_number: {
          required,
          phoneNum
        },
        address: {
          required,
        },
        marital_status: {
          ...validateIf(this.isPatient, {
            required,
          }),
        },
        department_id: {
          ...validateIf(this.isDoctor, {
            required,
            num,
            positiveNum,
          }),
        },
        specialization_details_id: {
          ...validateIf(this.isDoctor, {
            required,
            num,
            positiveNum,
          }),
        },
        experience_in_years: {
          ...validateIf(this.isDoctor, {
            required,
            num,
            positiveNum,
          }),
        },
        photo_of_doctor: {
          ...validateIf(this.isDoctor, {
            required,
          }),
        },
        category: {
          ...validateIf(this.isDoctor, {
            required,
          }),
        },
        price_of_the_appointment: {
          ...validateIf(this.isDoctor, {
            required,
            num,
            positiveNum,
          }),
        },
        schedule_details: {
          ...validateIf(this.isDoctor, {
            required,
          }),
        },
        degree: {
          ...validateIf(this.isDoctor, {
            required,
          }),
        },
        rating: {
          ...validateIf(this.isDoctor, {
            required,
          }),
        },
      },
    };
  },
  computed: {
    isPatient() {
      return this.role === "patient";
    },
    isDoctor() {
      return this.role === "doctor";
    },
  },
  methods: {
    async create() {
      this.$v.form.$touch();
      console.log(this.$v.form);
      if (this.$v.form.$invalid) return;
    },
    makeToast() {
      this.$bvToast.toast("Error occured", {
        title: "Error",
        toaster: "b-toaster-bottom-left",
        variant: "danger",
        solid: true,
      });
    },
  },
};
</script>
