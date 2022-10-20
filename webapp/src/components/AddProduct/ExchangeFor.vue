<template>
  <q-form autofocus @submit="submitForm" class="q-gutter-md">
    <q-card-section>
      <div class="text-h6">Trade with (Optional)</div>
      <div class="text-subtitle1">Select a product you want to trade with.</div>
    </q-card-section>
    <q-separator inset />

    <q-input
      v-model="formState.exchangeFor.name"
      label="Name"
      lazy-rules
      :rules="[(val) => (val && val.length > 0) || 'Please enter product name']"
    />
    <q-input
      min="1"
      type="number"
      v-model="formState.exchangeFor.quantity"
      label="Quantity"
      lazy-rules
      :rules="[
        (val) => (val && val.length > 0) || 'Please enter product quantity',
      ]"
    />
  </q-form>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "ExchangeFor",
  props: ["formState"],
  setup(props) {
    return {
      submitForm() {
        // console.log("formState", props.formState);
      },
    };
  },
});
</script>
