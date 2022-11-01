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
                    <label for="example-datepicker">Date of Birth</label>
                    <b-form-datepicker
                        id="example-datepicker"
                        v-model="form.birth_date"
                        class="mb-2"
                    ></b-form-datepicker>
                </div>
                <div class="custom-input">
                    <label for="example-datepicker">IIN Number</label>
                    <b-form-input
                        v-model="form.iin_number"
                        placeholder="Enter IIN Number"
                    ></b-form-input>
                </div>
                <div class="custom-input">
                    <label for="example-datepicker">ID Number</label>
                    <b-form-input
                        v-model="form.id_number"
                        placeholder="Enter ID Number"
                    ></b-form-input>
                </div>
                <div class="custom-input">
                    <label for="example-datepicker">Name</label>
                    <b-form-input
                        v-model="form.name"
                        placeholder="Enter Name"
                    ></b-form-input>
                </div>
                <div class="custom-input">
                    <label for="example-datepicker">Surname</label>
                    <b-form-input
                        v-model="form.surname"
                        placeholder="Enter Surname"
                    ></b-form-input>
                </div>
                <div class="custom-input">
                    <label for="example-datepicker">Middlename</label>
                    <b-form-input
                        v-model="form.middlename"
                        placeholder="Enter Middlename"
                    ></b-form-input>
                </div>
                <div class="custom-input">
                    <label for="example-datepicker">Blood Group</label>
                    <b-form-input
                        v-model="form.blood_group"
                        placeholder="Enter Blood Group"
                    ></b-form-input>
                </div>
                <div class="custom-input">
                    <label for="example-datepicker"
                        >Emergency Contact Number</label
                    >
                    <b-form-input
                        v-model="form.emergency_contact_number"
                        placeholder="Enter Emergency Contact Number"
                    ></b-form-input>
                </div>
                <div class="custom-input">
                    <label for="example-datepicker">Email (Optional)</label>
                    <b-form-input
                        v-model="form.email"
                        placeholder="Enter Email"
                    ></b-form-input>
                </div>
                <div class="custom-input">
                    <label for="example-datepicker">Address</label>
                    <b-form-input
                        v-model="form.address"
                        placeholder="Enter Address"
                    ></b-form-input>
                </div>
                <div class="custom-input">
                    <label for="example-datepicker">Address</label>
                    <b-form-input
                        v-model="form.address"
                        placeholder="Enter Address"
                    ></b-form-input>
                </div>
                <div class="custom-input">
                    <label for="example-datepicker">Marital Status</label>
                    <b-form-input
                        v-model="form.marital_status"
                        placeholder="Enter Marital Status"
                    ></b-form-input>
                </div>
                <b-button size="md" class="my-2 ml-sm-0" @click="create"
                    >create</b-button
                >
            </div>
        </div>
    </div>
</template>

<script>
export default {
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
            },
        };
    },
    methods: {
        async create() {},
    },
};
</script>
