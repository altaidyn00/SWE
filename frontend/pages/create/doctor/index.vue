<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Create Doctor</h1>
      <div class="d-flex flex-wrap justify-content-between my-4">
        <custom-date-input
          v-model="form.dateofbirth"
          class="custom-input"
          label="Date of Birth"
          :validation="$v.form.dateofbirth"
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
          v-model="form.fullname"
          class="custom-input"
          label="Fullname"
          placeholder="Enter Fullname"
          :validation="$v.form.fullname"
        />
        <custom-input
          v-model="form.contactnumber"
          class="custom-input"
          label="Contact Number"
          placeholder="Enter Contact Number"
          :validation="$v.form.contactnumber"
        />
        <custom-input
          v-model="form.address"
          class="custom-input"
          label="Address"
          placeholder="Enter Address"
          :validation="$v.form.address"
        />
        <custom-input
          v-model="form.depID"
          class="custom-input"
          label="Department ID"
          placeholder="Enter Department ID"
          :validation="$v.form.depID"
        />
        <custom-input
          v-model="form.specID"
          class="custom-input"
          label="Specialization Details ID"
          placeholder="Enter Specialization Details ID"
          :validation="$v.form.specID"
        />
        <custom-input
          v-model="form.expirience"
          class="custom-input"
          label="Experience in Years"
          placeholder="Enter Experience in Years"
          :validation="$v.form.expirience"
        />
        <custom-file-input
          class="custom-input"
          label="Doctor's Photo"
          placeholder=" Upload Doctor's Photos"
          v-model="form.photo"
          :validation="$v.form.photo"
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
          v-model="form.degree"
          :options="options.degree"
          :allow-empty="true"
          :validation="$v.form.degree"
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
import { required } from "vuelidate/lib/validators";
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
      role: "patient",
      form: {
        dateofbirth: null,
        iin: null,
        id: null,
        fullname: null,
        contactnumber: null,
        address: null,
        depID: null,
        specID: null,
        expirience: null,
        photo: null,
        category: null,
        degree: null,
        rating: null,
      },
    };
  },
  validations() {
    return {
      form: {
        dateofbirth: {
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
        fullname: {
          required,
        },
        contactnumber: {
          required,
          phoneNum,
        },
        address: {
          required,
        },
        depID: {
          required,
          num,
          positiveNum,
        },
        specID: {
          required,
          num,
          positiveNum,
        },
        experience: {
          required,
          num,
          positiveNum,
        },
        photo: {
          required,
        },
        category: {
          required,
        },
        degree: {
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
      const response = await this.createDoctor(this.form);
      console.log(response);
    },
  },
};
</script>
