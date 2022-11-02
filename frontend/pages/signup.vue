<template>
    <div class="custom-container">
        <div class="text-center">
            <h1>Registration</h1>
            <div class="mt-4">
                <h4>Choose role</h4>
                <div class="d-flex justify-content-between mt-2">
                    <div
                        class="registration-tab"
                        :class="{
                            'registration-tab__active': role === 'patient',
                        }"
                        @click="role = 'patient'"
                    >
                        Patient
                    </div>
                    <div
                        class="registration-tab"
                        :class="{
                            'registration-tab__active': role === 'doctor',
                        }"
                        @click="role = 'doctor'"
                    >
                        Doctor
                    </div>
                </div>
            </div>
            <div class="d-flex flex-wrap justify-content-between my-4">
                <div class="custom-input">
                    <span>Date of Birth</span>
                    <b-form-datepicker
                        v-model="form.birth_date"
                    ></b-form-datepicker>
                </div>
                <custom-input
                    v-model="form.iin_number"
                    class="custom-input"
                    label="IIN Number"
                    placeholder="Enter IIN Number"
                    :validation="$v.form.iin_number"
                />
                <custom-input
                    v-model="form.id_number"
                    class="custom-input"
                    label="ID Number"
                    placeholder="Enter ID Number"
                    :validation="$v.form.id_number"
                />
                <custom-input
                    v-model="form.name"
                    class="custom-input"
                    label="Name"
                    placeholder="Enter Name"
                    :validation="$v.form.name"
                />
                <custom-input
                    v-model="form.surname"
                    class="custom-input"
                    label="Surname"
                    placeholder="Enter Surname"
                    :validation="$v.form.surname"
                />
                <custom-input
                    v-model="form.middlename"
                    class="custom-input"
                    label="Middlename"
                    placeholder="Enter Middlename"
                    :validation="$v.form.middlename"
                />
                <custom-input
                    v-if="isPatient"
                    v-model="form.blood_group"
                    class="custom-input"
                    label="Blood Group"
                    placeholder="Enter Blood Group"
                    :validation="$v.form.blood_group"
                />
                <custom-input
                    v-if="isPatient"
                    v-model="form.emergency_contact_number"
                    class="custom-input"
                    label="Emergency Contact"
                    placeholder="Enter Emergency Contact Number"
                    :validation="$v.form.emergency_contact_number"
                />
                <custom-input
                    v-if="isPatient"
                    v-model="form.email"
                    class="custom-input"
                    label="Email (Optional)"
                    placeholder="Enter Email"
                />
                <custom-input
                    v-model="form.address"
                    class="custom-input"
                    label="Address"
                    placeholder="Enter Address"
                    :validation="$v.form.address"
                />
                <custom-input
                    v-if="isPatient"
                    v-model="form.marital_status"
                    class="custom-input"
                    label="Marital Status"
                    placeholder="Enter Marital Status"
                    :validation="$v.form.marital_status"
                />
                <custom-input
                    v-if="isDoctor"
                    v-model="form.department_id"
                    class="custom-input"
                    label="Department ID"
                    placeholder="Enter Department ID"
                    :validation="$v.form.department_id"
                />
                <custom-input
                    v-if="isDoctor"
                    v-model="form.specialization_details_id"
                    class="custom-input"
                    label="Specialization Details ID"
                    placeholder="Enter Specialization Details ID"
                    :validation="$v.form.specialization_details_id"
                />
                <custom-input
                    v-if="isDoctor"
                    v-model="form.experience_in_years"
                    class="custom-input"
                    label="Experience in Years"
                    placeholder="Enter Experience in Years"
                    :validation="$v.form.experience_in_years"
                />
                <custom-input
                    v-if="isDoctor"
                    v-model="form.category"
                    class="custom-input"
                    label="Category"
                    placeholder="Enter Category"
                    :validation="$v.form.category"
                />
            </div>
            <b-button size="md" class="my-2 ml-sm-0" @click="create"
                >create</b-button
            >
        </div>
    </div>
</template>

<script>
import { required, email, sameAs, minLength } from "vuelidate/lib/validators";
import CustomInput from "../components/ui/CustomInput.vue";

export default {
    components: { CustomInput },
    name: "signup",
    data() {
        return {
            role: "patient",
            form: {
                birth_date: null,
                iin_number: null,
                id_number: null,
                name: null,
                surname: null,
                middlename: null,
                blood_group: null,
                emergency_contact_number: null,
                email: null,
                address: null,
                marital_status: null,
                registration_date: null,
                department_id: null,
                specialization_details_id: null,
                experience_in_years: null,
                photo_of_doctor: null,
                category: null,
                price_of_the_appointment: null,
                schedule_details: null,
                degree: null,
                rating: null,
                homepage_url: null,
            },
        };
    },
    validations() {
        return {
            form: {
                // birth_date: {
                //     required,
                // },
                iin_number: {
                    required,
                },
                id_number: {
                    required,
                },
                name: {
                    required,
                },
                surname: {
                    required,
                },
                middlename: {
                    required,
                },
                blood_group: {
                    required,
                },
                emergency_contact_number: {
                    required,
                },
                address: {
                    required,
                },
                marital_status: {
                    required,
                },
            },
        };
    },
    computed: {
        isPatient() {
            return this.role === "patient";
        },
        isDoctor() {
            return this.role === "doctor";
        },
    },
    methods: {
        async create() {
            this.$v.form.$touch();
            console.log(this.$v);
            if (this.$v.form.$invalid) return;
            console.log("asasd");
        },
    },
};
</script>
