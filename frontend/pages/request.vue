<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Request</h1>
    </div>
  </div>
</template>

<script>
export default {
  middleware: ["loggedin", "notadmin"],
  name: "request",
  data() {
    return {};
  },
  methods: {},
};
</script>
