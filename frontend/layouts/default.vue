<template>
  <div>
    <div>
      <div>
        <Header />
        <Nuxt />
      </div>
    </div>
    <Footer />
  </div>
</template>

<script>
export default {
  computed: {
    isLoggedIn() {
      return this.$auth.loggedIn;
    },
  },
  mounted() {
    if (this.isLoggedIn) {
      const user = JSON.parse(localStorage.getItem("user"));
      this.$auth.setUser(user);
    }
  },
};
</script>
