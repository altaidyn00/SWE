<template>
  <div class="header">
    <b-navbar toggleable="lg">
      <b-navbar-brand href="/">Medical Center</b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item href="/services">Services</b-nav-item>
          <b-nav-item href="/about-us">About Us</b-nav-item>
          <b-nav-item href="/contacts">Contacts</b-nav-item>
          <!-- <b-nav-item href="/appointment">Appointment</b-nav-item> -->
          <b-nav-item href="/create">Create</b-nav-item>
        </b-navbar-nav>

        <b-navbar-nav class="ml-auto">
          <b-nav-form>
            <b-button size="sm" class="button my-2 ml-sm-0" @click="goToSignin"
              >signin</b-button
            ></b-nav-form
          >
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>

<script>
export default {
  name: "Header",
  methods: {
    goToSignin() {
      this.$router.push("/signin");
    },
  },
};
</script>