<template>
  <div class="d-flex flex-column">
    <div class="w-100">
      <span>{{ label }}</span>
      <b-form-datepicker
        :value="value"
        :class="{
          'is-invalid': validation && validation.$error,
        }"
        @input="input"
      ></b-form-datepicker>
      <div v-if="requiredError" class="datepicker__error invalid-feedback">
        This is a required field
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "CustomDatepicker",
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
  },
  computed: {
    requiredError() {
      return (
        this.validation &&
        Object.prototype.hasOwnProperty.call(this.validation, "required") &&
        !this.validation.required &&
        this.validation.$error
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
