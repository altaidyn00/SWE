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
  created() {
    if (this.isLoggedIn) {
      const user = this.localStorage.getItem("user");
      console.log(user, "user");
    }
  },
};
</script>
