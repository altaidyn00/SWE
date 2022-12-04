<template>
  <div class="custom-container">
    <div class="text-center">
      <h1>Request</h1>
      <custom-input
        v-model="search"
        class="custom-input w-100"
        label="Search Doctor"
        placeholder="Search"
        @onEnter="filter"
      />
      {{ search }}
      <div class="d-flex flex-column mt-4" v-for="doctor in data">
        <div class="doctor d-flex flex-row justify-content-between">
          <div>{{ doctor.user.id }}</div>
          <div>{{ doctor.user.first_name }}</div>
          <div>{{ doctor.user.last_name }}</div>
          <div>{{ doctor.user.email }}</div>
          <div>{{ doctor.doctor.category }}</div>
          <div>{{ doctor.doctor.rating }}</div>
          <div>{{ doctor.doctor.appointment_price }}KZT</div>
          <div>{{ spec(doctor.doctor.department_id) }}</div>
          <div>{{ dep(doctor.doctor.specialization_id) }}</div>
          <div @click="toDetails(doctor)" class="application__link">
            make appointment
          </div>
        </div>
      </div>
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
          doctor.user.first_name.includes(this.search) ||
          this.spec(doctor.doctor.specialization_id).description.includes(
            this.search
          ) ||
          this.dep(doctor.doctor.department_id).description.includes(
            this.search
          )
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
