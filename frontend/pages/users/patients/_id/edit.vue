<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Edit Patient</h1>
      <div class="d-flex flex-wrap justify-content-between my-4">
        <custom-date-input
          v-model="form.dateofbirth"
          class="custom-input"
          label="Date of Birth"
        />
        <custom-input
          v-model="form.iin"
          class="custom-input"
          label="IIN Number"
          placeholder="Enter IIN Number"
        />
        <custom-input
          v-model="form.id"
          class="custom-input"
          label="ID Number"
          placeholder="Enter ID Number"
        />
        <custom-input
          v-model="form.fullname"
          class="custom-input"
          label="Fullname"
          placeholder="Enter Fullname"
        />
        <custom-select
          class="custom-input"
          label="Blood Group"
          placeholder="Select Blood Group"
          v-model="form.bloodgroup"
          :options="options.blood_group"
          :allow-empty="true"
        />
        <custom-input
          v-model="form.emergencynumber"
          class="custom-input"
          label="Emergency Contact"
          placeholder="Enter Emergency Contact Number"
        />
        <custom-input
          v-model="form.contactnumber"
          class="custom-input"
          label="Contact Number"
          placeholder="Enter Contact Number"
        />
        <custom-input
          v-model="form.email"
          class="custom-input"
          label="Email (Optional)"
          placeholder="Enter Email"
        />
        <custom-input
          v-model="form.address"
          class="custom-input"
          label="Address"
          placeholder="Enter Address"
        />
        <custom-select
          class="custom-input"
          label="Marital Status"
          placeholder="Select Marital Status"
          v-model="form.maritalstatus"
          :options="options.marital_status"
          :allow-empty="true"
        />
      </div>
      <b-button size="md" class="button my-2 ml-sm-0" @click="edit"
        >edit</b-button
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
      role: "patient",
      form: {
        dateofbirth: null,
        iin: null,
        id: null,
        fullname: null,
        bloodgroup: null,
        emergencynumber: null,
        contactnumber: null,
        email: null,
        address: null,
        maritalstatus: null,
        registrationdate: "sadasdas",
      },
    };
  },
  //   validations() {
  //     return {
  //       form: {
  //         dateofbirth: {
  //           required,
  //         },
  //         iin: {
  //           required,
  //           num,
  //           positiveNum,
  //         },
  //         id: {
  //           required,
  //           num,
  //           positiveNum,
  //         },
  //         fullname: {
  //           required,
  //         },
  //         email: {
  //           required,
  //           email,
  //         },
  //         bloodgroup: {
  //           required,
  //         },
  //         emergencynumber: {
  //           required,
  //           phoneNum,
  //         },
  //         contactnumber: {
  //           required,
  //           phoneNum,
  //         },
  //         address: {
  //           required,
  //         },
  //         maritalstatus: {
  //           required,
  //         },
  //       },
  //     };
  //   },
  methods: {
    ...mapActions({
      editPatient: "users/edit_patient",
    }),
    async edit() {
      //   this.$v.form.$touch();
      //   if (this.$v.form.$invalid) return;
      const id = this.$route.params.id;
      const payload = this.form;
      const response = await this.editPatient({ id, payload });
      console.log(response);
    },
  },
};
</script>
