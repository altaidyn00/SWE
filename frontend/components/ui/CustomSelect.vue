<template>
  <div class="d-flex flex-column">
    <div class="w-100">
      <span>{{ label }}</span>
      <Multiselect
        :value="value"
        :options="options"
        :multiple="multiple"
        :allow-empty="allowEmpty"
        :show-labels="false"
        :searchable="false"
        :disabled="disabled"
        :placeholder="placeholder"
        :class="{
          'is-invalid multiselect__error': validation && validation.$error,
        }"
        :close-on-select="!multiple"
        @input="input"
      >
      </Multiselect>
      <div v-if="requiredError" class="invalid-feedback">
        This is a required field
      </div>
    </div>
  </div>
</template>

<script>
import Multiselect from "vue-multiselect";

export default {
  components: { Multiselect },
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
