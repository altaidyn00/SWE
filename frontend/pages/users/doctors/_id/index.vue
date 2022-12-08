<template>
  <div class="custom-container">
    <div class="text-center d-flex">
      <div class="d-flex flex-column w-50" :class="{ 'mx-auto': isAdmin }">
        <h1>Doctor {{ doctor.user.id }}</h1>
        <div
          class="d-flex patient-info mt-4 w-100 justify-content-between align-items-center px-4"
          :class="{ 'mr-4': !isAdmin, 'mx-auto': isAdmin }"
        >
          <div class="d-flex flex-column w-50">
            <div class="d-flex flex-row">
              <div class="mr-2">Birth Date:</div>
              <div class="font-weight-bold">
                {{ doctor.doctor.date_of_birth }}
              </div>
            </div>
            <div class="d-flex flex-row">
              <div class="mr-2">IIN:</div>
              <div class="font-weight-bold">{{ doctor.doctor.iin }}</div>
            </div>
            <div class="d-flex flex-row">
              <div class="mr-2">First name:</div>
              <div class="font-weight-bold">{{ doctor.user.first_name }}</div>
            </div>
            <div class="d-flex flex-row">
              <div class="mr-2">Last name:</div>
              <div class="font-weight-bold">{{ doctor.user.last_name }}</div>
            </div>
            <div class="d-flex flex-row">
              <div class="mr-2">Email:</div>
              <div class="font-weight-bold">{{ doctor.user.email }}</div>
            </div>
            <div class="d-flex flex-row">
              <div class="mr-2">Category:</div>
              <div class="font-weight-bold">{{ doctor.doctor.category }}</div>
            </div>
            <div class="d-flex flex-row">
              <div class="mr-2">Contact number:</div>
              <div class="font-weight-bold">
                {{ doctor.doctor.contact_number }}
              </div>
            </div>
            <div class="d-flex flex-row">
              <div class="mr-2">Education degree:</div>
              <div class="font-weight-bold">
                {{ doctor.doctor.education_degree }}
              </div>
            </div>
            <div class="d-flex flex-row">
              <div class="mr-2">Address:</div>
              <div class="font-weight-bold">{{ doctor.doctor.address }}</div>
            </div>
            <div class="d-flex flex-row">
              <div class="mr-2">Appointment price:</div>
              <div class="font-weight-bold">
                {{ doctor.doctor.appointment_price }}KZT
              </div>
            </div>
            <div class="d-flex flex-row">
              <div class="mr-2">Department id:</div>
              <div class="font-weight-bold">
                {{ doctor.doctor.department_id }}
              </div>
            </div>
            <div class="d-flex flex-row">
              <div class="mr-2">Specialization details id:</div>
              <div class="font-weight-bold">
                {{ doctor.doctor.specialization_details_id }}
              </div>
            </div>
            <div class="d-flex flex-row">
              <div class="mr-2">Number of patients:</div>
              <div class="font-weight-bold">
                {{ doctor.doctor.number_of_patients }}
              </div>
            </div>
            <div class="d-flex flex-row">
              <div class="mr-2">Role:</div>
              <div class="font-weight-bold">{{ doctor.user.role }}</div>
            </div>
          </div>
          <div class="image" :class="{ 'image-hover': isAdmin }">
            <label for="avatar-input" class="image-hover">
              <img
                v-if="doctor.doctor.photolocation.length > 4"
                class="img-avatar w-100 h-100"
                :src="pht"
              />
              <img
                v-else
                class="img-avatar w-100"
                src="~/assets/default-avatar.png"
              />
            </label>
            <input
              id="avatar-input"
              class="d-none"
              accept="image/png, image/gif, image/jpeg, image/jpg"
              type="file"
              @change="setAvatar"
            />
          </div>
        </div>
      </div>
      <div
        class="d-flex flex-column w-50 justify-content-center align-items-left ml-4"
        v-if="!isAdmin"
      >
        <h1>Make appointment</h1>
        <custom-input
          v-model="doctor_id"
          class="custom-input w-1000"
          label="Doctor ID"
          placeholder="Enter Doctor ID"
          :validation="$v.doctor_id"
        />
        <custom-input
          v-model="preferred_date"
          class="custom-input w-1000"
          label="Date"
          placeholder="Enter date"
          :validation="$v.preferred_date"
        />
        <custom-input
          v-model="preferred_time"
          class="custom-input w-1000"
          label="Time"
          placeholder="Enter Time"
          :validation="$v.preferred_time"
        />
        <custom-input
          v-model="name"
          class="custom-input w-1000"
          label="Name"
          placeholder="Enter Name"
          :validation="$v.name"
        />
        <custom-input
          v-model="surname"
          class="custom-input w-1000"
          label="Surname"
          placeholder="Enter Surname"
          :validation="$v.surname"
        />
        <custom-input
          v-model="email"
          class="custom-input w-1000"
          label="Email"
          placeholder="Enter Email"
          :validation="$v.email"
        />
        <b-button size="md" class="button my-2 ml-sm-0" @click="make"
          >make</b-button
        >
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { required, email } from "vuelidate/lib/validators";
import CustomInput from "../../../../components/ui/CustomInput.vue";
import objectToFormData from "~/helpers/objectToFormData";

export default {
  // middleware: ["loggedin", "admin"],
  components: { CustomInput },
  name: "patients",
  async asyncData({ store, params }) {
    const query = {};
    console.log(params);
    query.id = params.id;
    await store.dispatch("users/get_doctor", query);
  },
  data() {
    return {
      doctor_id: null,
      preferred_date: null,
      preferred_time: null,
      name: null,
      surname: null,
      email: null,
      avatar: null,
    };
  },
  validations() {
    return {
      doctor_id: {
        required,
      },
      preferred_date: {
        required,
      },
      preferred_time: {
        required,
      },
      name: {
        required,
      },
      surname: {
        required,
      },
      email: {
        required,
      },
    };
  },
  mounted() {
    this.doctor_id = +this.doctor.user.id;
    this.name = this.$auth.user.first_name;
    this.surname = this.$auth.user.last_name;
    this.email = this.$auth.user.email;
  },
  computed: {
    ...mapGetters({
      doctor: "users/doctor",
    }),
    isAdmin() {
      return this.$auth.user.role === "admin";
    },
    pht() {
      return require("~/" +
        this.doctor?.doctor?.photolocation.substring(
          13,
          this.doctor?.doctor?.photolocation.length
        ));
    },
  },
  methods: {
    objectToFormData,
    async setAvatar(event) {
      if (!event) return;
      this.avatar = event.target.files[0];
      const asd = {};
      asd.id = +this.doctor_id;
      asd.photo = this.avatar;
      const payload = this.objectToFormData(asd);
      const response = await this.$store.dispatch(
        "users/upload_photo",
        payload
      );
    },
    async make() {
      this.$v.$touch();
      if (this.$v.$invalid) return;
      const payload = {};
      payload.doctor_id = +this.doctor_id;
      payload.preferred_date = this.preferred_date;
      payload.preferred_time = this.preferred_time;
      payload.name = this.name;
      payload.surname = this.surname;
      payload.email = this.email;
      const response = await this.$store.dispatch(
        "users/make_appointment",
        payload
      );
    },
  },
};
</script>

<style lang="scss">
.w-1000 {
  width: 100% !important;
}

.image {
  width: 300px;
  height: auto;
  border-radius: 12px;
}

.image-hover:hover {
  cursor: pointer;
}

.img-avatar {
  border-radius: 12px;
}
</style>
