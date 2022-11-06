<template>
  <div class="d-flex flex-column">
    <div class="w-100">
      <span>{{ label }}</span>
      <b-form-file
        :multiple="multiple"
        :value="value"
        :accept="accept"
        @input="change"
        :class="{
          'fileinput__error is-invalid': validation && validation.$error,
        }"
      ></b-form-file>
    </div>
    <div v-if="requiredError" class="input-error">This is a required field</div>
  </div>
</template>
<script>
export default {
  name: "CustomFileInput",
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
    multiple: {
      // type: Boolean,
      default: false,
    },
    accept: {
      // type: String,
      required: false,
      default: "image/jpeg, image/png, image/jpg",
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
    change(value) {
      this.$emit("change", value);
    },
  },
};
</script>

<style lang="scss" scoped></style>
