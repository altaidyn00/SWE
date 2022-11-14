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
          {{ data.item.id }}
        </template>
        <template #cell(fullname)="data">
          {{ data.item.fullname }}
        </template>
        <template #cell(contactnumber)="data">
          {{ data.item.contactnumber }}
        </template>
        <template #cell(iin)="data">
          {{ data.item.iin }}
        </template>
        <template #cell(dateofbirth)="data">
          {{ data.item.dateofbirth }}
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
            @click="toEdit(data.item)"
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
          key: "fullname",
          label: "Fullname",
        },
        {
          key: "contactnumber",
          label: "Contact Number",
        },
        {
          key: "iin",
          label: "IIN",
        },
        {
          key: "dateofbirth",
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
  mounted() {
    console.log(this.patients);
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
      this.$router.push(`/users/patients/${item.id}`);
    },
    toDetails(item) {
      this.$router.push(`/users/patients/edit/${item.id}`);
    },
  },
};
</script>
