<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Requests</h1>
    </div>
    <div>
      <b-table
        class="custom-border"
        table-class="table table-centered w-100"
        thead-tr-class="bg-light"
        tbody-tr-class="hover"
        :items="requests"
        :fields="fields"
        :bordered="true"
        :fixed="true"
        responsive="sm"
      >
        <template #cell(doctor_id)="data">
          {{ data.item.doctor_id }}
        </template>
        <template #cell(email)="data">
          {{ data.item.email }}
        </template>
        <template #cell(name)="data">
          {{ data.item.name }}
        </template>
        <template #cell(surname)="data">
          {{ data.item.surname }}
        </template>
        <template #cell(preferred_date)="data">
          {{ data.item.preferred_date }}
        </template>
        <template #cell(preferred_time)="data">
          {{ data.item.preferred_time }}
        </template>
        <template #cell(accept)="data">
          <b-button size="md" class="button my-2 ml-sm-0" @click="accept"
            >accept</b-button
          >
        </template>
      </b-table>
    </div>
  </div>
</template>

<script>
export default {
  name: "patients",
  async asyncData({ store }) {
    const requests = await store.dispatch("users/get_appointments");
    return { requests };
  },
  data() {
    return {
      fields: [
        {
          key: "doctor_id",
          label: "Doctor ID",
        },
        {
          key: "email",
          label: "Email",
        },
        {
          key: "name",
          label: "Name",
        },
        {
          key: "surname",
          label: "Surname",
        },
        {
          key: "preferred_date",
          label: "Pereferred Date",
        },
        {
          key: "preferred_time",
          label: "Pereferred Time",
        },
        {
          key: "accept",
          label: "",
        },
      ],
    };
  },
  methods: {
    accept() {
      this.$bvToast.toast("Request successfully accepted!", {
        title: "Requesdt",
        toaster: "b-toaster-bottom-left",
        variant: "success",
        solid: true,
      });
    },
  },
};
</script>
