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
  middleware: ["notloggedin"],
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
        const access_token = response.data.Acces_token.value;
        const user = response.data.user;
        if (access_token) {
          this.$auth.setUserToken(access_token, null);
          localStorage.setItem("user", JSON.stringify(user));
          this.$auth.setUser(user);
        }
      } catch (e) {}
    },
  },
};
</script>
