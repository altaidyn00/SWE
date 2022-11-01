<template>
    <div class="d-flex flex-column">
        <div class="w-100">
            <span>{{ label }}</span>
            <b-form-input
                :value="value"
                :placeholder="placeholder"
                :class="{
                    'is-invalid': validation && validation.$error,
                }"
                @input="input"
                @keyup.enter="onEnter"
            ></b-form-input>
        </div>
        <div v-if="requiredError" class="input-error">
            This is a required field
        </div>
    </div>
</template>
<script>
export default {
    name: "CustomInput",
    props: {
        label: {
            // type: String,
            required: false,
            default: "",
        },
        placeholder: {
            // type: String,
            required: false,
            default: "",
        },
        validation: {
            // type: Object,
            required: false,
            default: null,
        },
        value: {
            required: false,
            default: "",
        },
    },
    computed: {
        requiredError() {
            return (
                this.validation &&
                Object.prototype.hasOwnProperty.call(
                    this.validation,
                    "required"
                ) &&
                !this.validation.required &&
                this.validation.$error
            );
        },
    },
    methods: {
        input(value) {
            this.$emit("input", value);
        },
        onEnter() {
            this.$emit("onEnter");
        },
    },
};
</script>

<style lang="scss" scoped></style>
