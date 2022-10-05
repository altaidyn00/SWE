<template>
  <div id="app">
    <b-table :items="items">
      <template v-slot:custom-foot="data">
        <b-tr>
          <b-td>Name</b-td>
          <b-td>Surname</b-td>
        </b-tr>
      </template>
    </b-table>
  </div>
</template>
<script>
export default {
  name: "App",
  data() {
    return {
      items: [
        { firstName: "alex", lastName: "green" },
        {
          firstName: "may",
          lastName: "smith",
        },
        { firstName: "james", lastName: "jones" },
      ],
    };
  },
};
</script>
