<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Request</h1>
      <custom-input
        v-model="search"
        class="w-100 my-4"
        label=""
        placeholder="Search"
        @onEnter="filter"
      />
      <b-table
        class="custom-border"
        table-class="table table-centered w-100"
        thead-tr-class="bg-light"
        tbody-tr-class="hover"
        :items="data"
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
        <template #cell(category)="data">
          {{ data.item.doctor.category }}
        </template>
        <template #cell(rating)="data">
          {{ data.item.doctor.rating }}
        </template>
        <template #cell(appointment_price)="data">
          {{ data.item.doctor.appointment_price }} KZT
        </template>
        <template #cell(department_id)="data">
          {{ dep(data.item.doctor.department_id).description }}
        </template>
        <template #cell(specialization_id)="data">
          {{ spec(data.item.doctor.specialization_id).description }}
        </template>
        <template #cell(action1)="data">
          <div
            class="application__link font-weight-bold"
            @click="toDetails(data.item)"
          >
            appointment
          </div>
        </template>
      </b-table>
    </div>
  </div>
</template>

<script>
import { required, email } from "vuelidate/lib/validators";
import { mapGetters } from "vuex";
import CustomInput from "../components/ui/CustomInput.vue";
import options from "~/helpers/options";
import { num, positiveNum, phoneNum } from "~/helpers/validators";
export default {
  middleware: ["loggedin", "notadmin"],
  components: { CustomInput },
  name: "request",
  async asyncData({ store }) {
    await store.dispatch("users/get_doctors");
    const departments = await store.dispatch("users/get_departments");
    const specializations = await store.dispatch("users/get_specializations");
    return { departments, specializations };
  },
  data() {
    return {
      search: null,
      data: null,
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
          key: "category",
          label: "Category",
        },
        {
          key: "rating",
          label: "Rating",
        },
        {
          key: "appointment_price",
          label: "Price",
        },
        {
          key: "department_id",
          label: "Department",
        },
        {
          key: "specialization_id",
          label: "Specialization",
        },
        {
          key: "action1",
          label: "",
        },
      ],
    };
  },
  computed: {
    ...mapGetters({
      doctors: "users/doctors",
    }),
  },
  watch: {
    search() {
      this.filter();
      if (this.search === null || this.search === "") this.data = this.doctors;
    },
  },
  mounted() {
    this.data = this.doctors;
  },
  methods: {
    spec(data) {
      return this.specializations?.find((x) => x?.id === data);
    },
    dep(data) {
      return this.departments?.find((x) => x?.id === data);
    },
    filter() {
      this.data = this.doctors.filter(
        (doctor) =>
          doctor.user.first_name
            .toUpperCase()
            .includes(this.search.toUpperCase()) ||
          this.spec(doctor.doctor.specialization_id)
            .description.toUpperCase()
            .includes(this.search.toUpperCase()) ||
          this.dep(doctor.doctor.department_id)
            .description.toUpperCase()
            .includes(this.search.toUpperCase())
      );
    },
    toDetails(item) {
      this.$router.push(`/users/doctors/${item.user.id}`);
    },
  },
};
</script>

<style scoped lang="scss">
.doctor {
  padding: 12px;
  border: 1px solid black;
  border-radius: 12px;
}

.application__link {
  color: blue;
}

.application__link:hover {
  cursor: pointer;
}
</style>
