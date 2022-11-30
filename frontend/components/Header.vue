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
          <b-nav-item v-if="isLoggedIn" href="/request">Request</b-nav-item>
          <b-nav-item-dropdown
            v-if="isLoggedIn"
            id="my-nav-dropdown"
            text="Create"
            toggle-class="nav-link-custom"
            right
          >
            <b-dropdown-item @click="goToCreateDoctor">Doctor</b-dropdown-item>
            <b-dropdown-item @click="goToCreatePatient"
              >Patient</b-dropdown-item
            >
          </b-nav-item-dropdown>
          <b-nav-item-dropdown
            v-if="isLoggedIn"
            id="my-nav-dropdown"
            text="Users"
            toggle-class="nav-link-custom"
            right
          >
            <b-dropdown-item @click="goToDoctors">Doctors</b-dropdown-item>
            <b-dropdown-item @click="goToPatients">Patients</b-dropdown-item>
          </b-nav-item-dropdown>
        </b-navbar-nav>

        <b-navbar-nav class="ml-auto">
          <b-nav-form>
            <b-button v-if="isLoggedIn" size="sm" class="button my-2 ml-sm-0" @click="goToSignin"
              >signin</b-button
            >
            </b-nav-form
          >
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>

<script>
export default {
  name: "Header",
  computed: {
    isLoggedIn() {
      return this.$auth.loggedIn;
    },
  },
  methods: {
    goToSignin() {
      this.$router.push("/signin");
    },
    goToPatients() {
      this.$router.push("/users/patients");
    },
    goToDoctors() {
      this.$router.push("/users/doctors");
    },
    goToCreatePatient() {
      this.$router.push("/create/patient");
    },
    goToCreateDoctor() {
      this.$router.push("/create/doctor");
    },
  },
};
</script>
