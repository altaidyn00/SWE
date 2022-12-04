<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Patients</h1>
    </div>
    <div>
      <b-table
        class="custom-border"
        table-class="table table-centered w-100"
        thead-tr-class="bg-light"
        tbody-tr-class="hover"
        :items="patients"
        :fields="fields"
        :bordered="true"
        :fixed="true"
        responsive="sm"
      >
        <template #cell(id)="data">
          {{ data.item.user.id }}
        </template>
        <template #cell(first_name)="data">
          {{ data.item.user.first_name }}
        </template>
        <template #cell(last_name)="data">
          {{ data.item.user.last_name }}
        </template>
        <template #cell(contact_number)="data">
          {{ data.item.patient.contact_number }}
        </template>
        <template #cell(iin)="data">
          {{ data.item.patient.iin }}
        </template>
        <template #cell(date_of_birth)="data">
          {{ data.item.patient.date_of_birth }}
        </template>
        <template #cell(action1)="data">
          <div
            class="application__link font-weight-bold"
            @click="toDetails(data.item)"
          >
            view
          </div>
        </template>
        <template #cell(action2)="data">
          <div
            class="application__link font-weight-bold"
            @click="toModify(data.item)"
          >
            edit
          </div>
        </template>
      </b-table>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
export default {
  // middleware: ["loggedin", "admin"],
  name: "patients",
  async asyncData({ store }) {
    await store.dispatch("users/get_patients");
  },
  data() {
    return {
      fields: [
        {
          key: "id",
          label: "Id",
        },
        {
          key: "first_name",
          label: "First Name",
        },
        {
          key: "last_name",
          label: "Last Name",
        },
        {
          key: "contact_number",
          label: "Contact Number",
        },
        {
          key: "iin",
          label: "IIN",
        },
        {
          key: "date_of_birth",
          label: "Birth Date",
        },
        {
          key: "action1",
          label: "",
        },
        {
          key: "action2",
          label: "",
        },
      ],
    };
  },
  computed: {
    ...mapGetters({
      patients: "users/patients",
    }),
    items() {
      return this.patiens;
    },
  },
  methods: {
    toDetails(item) {
      this.$router.push(`/users/patients/${item.patient.id}`);
    },
    toModify(item) {
      this.$router.push(`/users/patients/${item.patient.id}/edit`);
    },
  },
};
</script>

<style scoped lang="scss">
.application__link {
  color: blue;
}

.application__link:hover {
  cursor: pointer;
}
</style>
