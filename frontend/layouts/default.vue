<template>
  <div>
    <div>
      <div>
        <Header />
        <Nuxt @setUser="setUser" />
      </div>
    </div>
    <Footer />
  </div>
</template>

<script>
export default {
  created() {
    if (this.$auth.loggedIn) {
      const user = JSON.parse(localStorage.getItem("user"));
      this.$auth.setUser(user);
    }
  },
};
</script>
