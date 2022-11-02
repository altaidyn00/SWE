<template>
    <div class="custom-input__label">
        <label v-if="label">{{ label }}</label>
        <Multiselect
            :value="value"
            :track-by="trackBy"
            :custom-label="customLabel"
            :options="options"
            :multiple="multiple"
            :allow-empty="allowEmpty"
            :show-labels="false"
            :searchable="searchable"
            :disabled="disabled"
            :placeholder="placeholder"
            :class="{
                'is-invalid multiselect__error':
                    validation && validation.$error,
            }"
            :close-on-select="!multiple"
            @input="input"
        >
        </Multiselect>
        <div v-if="requiredError" class="invalid-feedback">
            This is a required field
        </div>
    </div>
</template>

<script>
export default {
    name: "CustomSelect",
    props: {
        label: {
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
            required: true,
        },
        customLabel: {
            // type: Function,
            required: false,
            default: (option) => option,
        },
        trackBy: {
            // type: String,
            default: null,
        },
        options: {
            // type: Array,
            required: true,
        },
        multiple: {
            // type: Boolean,
            default: false,
        },
        allowEmpty: {
            // type: Boolean,
            default: true,
        },
        placeholder: {
            // type: String,
            required: false,
            default: "Select option",
        },
        disabled: {
            // type: String,
            required: false,
            default: false,
        },
        searchable: {
            required: false,
            default: false,
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
                !this.validation.required
            );
        },
    },
    methods: {
        input(value) {
            this.$emit("input", value);
        },
    },
};
</script>
