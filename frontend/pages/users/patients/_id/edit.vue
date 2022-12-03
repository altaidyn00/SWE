<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Edit Patient</h1>
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
        <custom-select
          class="custom-input"
          label="Blood Group"
          placeholder="Select Blood Group"
          v-model="form.blood_group"
          :options="options.blood_group"
          :allow-empty="true"
          :validation="$v.form.blood_group"
        />
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
      form: {
        date_of_birth: null,
        iin: null,
        blood_group: null,
        emergency_contact_number: null,
        contact_number: null,
        address: null,
        marital_status: null,
        registration_date: null,
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
      editPatient: "users/edit_patient",
    }),
    async edit() {
        this.$v.form.$touch();
        if (this.$v.form.$invalid) return;
      var today = new Date();
      var dd = String(today.getDate()).padStart(2, "0");
      var mm = String(today.getMonth() + 1).padStart(2, "0"); //January is 0!
      var yyyy = today.getFullYear();
      today = yyyy + "-" + mm + "-" + dd;
      this.form.registration_date = today;
      const id = this.$route.params.id;
      const payload = this.form;
      const response = await this.editPatient({ id, payload });
      console.log(response);
    },
  },
};
</script>
