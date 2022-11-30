<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Log in</h1>
      <div class="d-flex flex-wrap justify-content-between my-4">
        <custom-input
          v-model="form.username"
          class="custom-input"
          label="Username"
          placeholder="Enter Username"
          :validation="$v.form.username"
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
import { mapActions, mapGetters } from "vuex";

export default {
  components: { CustomInput },
  name: "signin",
  data() {
    return {
      form: {
        username: null,
        password: null,
      },
    };
  },
  validations() {
    return {
      form: {
        username: {
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
    ...mapActions({
      login: "users/login",
    }),
    async signin() {
      try {
        const response = await this.$auth.loginWith("local", {
          data: this.form,
        });
        const access_token = response.data.Value;
        console.log(access_token, "access_token");
        const token = await this.$auth.setUserToken(access_token, null);
        const user = await this.$store.dispatch("users/get_current_user", {
          payload: response.data,
        });
        console.log(user);
      } catch (e) {}
    },
  },
};
</script>
