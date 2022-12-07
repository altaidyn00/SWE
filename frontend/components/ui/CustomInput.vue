<template>
  <div class="d-flex flex-column">
    <div class="w-100">
      <span>{{ label }}</span>
      <b-form-input
        :value="value"
        :placeholder="placeholder"
        :type="type"
        :class="{
          'is-invalid': validation && validation.$error,
        }"
        @input="input"
        @keyup.enter="onEnter"
      ></b-form-input>
    </div>
    <div v-if="requiredError" class="input-error">This is a required field</div>
    <div v-else-if="numError" class="input-error">Should be numeric</div>
    <div v-else-if="positiveNumError" class="input-error">
      Should be positive number
    </div>
    <div v-else-if="phoneNumError" class="input-error">
      Should in format of a phone number
    </div>
    <div v-else-if="emailError" class="input-error">
      Should in format of a email
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
    type: {
      // type: String,
      required: false,
      default: "text",
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
    minLength: {
      required: false,
      default: 9,
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
    numError() {
      return (
        this.validation &&
        Object.prototype.hasOwnProperty.call(this.validation, "num") &&
        !this.validation.num &&
        this.validation.$error
      );
    },
    positiveNumError() {
      return (
        this.validation &&
        Object.prototype.hasOwnProperty.call(this.validation, "positiveNum") &&
        !this.validation.positiveNum &&
        this.validation.$error
      );
    },
    phoneNumError() {
      return (
        this.validation &&
        Object.prototype.hasOwnProperty.call(this.validation, "phoneNum") &&
        !this.validation.phoneNum &&
        this.validation.$error
      );
    },
    emailError() {
      return (
        this.validation &&
        Object.prototype.hasOwnProperty.call(this.validation, "email") &&
        !this.validation.email &&
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
