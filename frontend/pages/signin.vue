<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Log in</h1>
      <div class="d-flex flex-wrap justify-content-between my-4">
        <custom-input
          v-model="form.iin_number"
          class="custom-input"
          label="IIN Number"
          placeholder="Enter IIN Number"
          :validation="$v.form.iin_number"
        />
        <custom-input
          v-model="form.password"
          class="custom-input"
          label="Password"
          placeholder="Enter Password"
          :validation="$v.form.password"
        />
      </div>
      <b-button size="md" class="button my-2 ml-sm-0" @click="signin"
        >sign in</b-button
      >
    </div>
  </div>
</template>

<script>
import { required } from "vuelidate/lib/validators";
import CustomInput from "../components/ui/CustomInput.vue";
import { num, positiveNum } from "~/helpers/validators";

export default {
  components: { CustomInput },
  name: "signin",
  data() {
    return {
      form: {
        iin_number: null,
        password: null,
      },
    };
  },
  validations() {
    return {
      form: {
        iin_number: {
          required,
          num,
          positiveNum,
        },
        password: {
          required,
        },
      },
    };
  },
  methods: {
    async signin() {
      this.$v.form.$touch();
      console.log(this.$v.form);
      if (this.$v.form.$invalid) return;
    },
  },
};
</script>
