<template>
  <q-form :ref="formRef" autofocus class="q-gutter-md">
    <q-card-section>
      <div class="text-h6">{{ id ? "Edit Product" : "Create Product" }}</div>
      <div class="text-subtitle1">
        {{
          id
            ? "Edit your product and click save"
            : "Fill out the following form to create and list your product."
        }}
      </div>
    </q-card-section>
    <q-separator inset />

    <q-input
      v-model="formState.name"
      label="Name"
      lazy-rules
      :rules="[(val) => (val && val.length > 0) || 'Please enter product name']"
    />
    <q-input
      type="textarea"
      v-model="formState.description"
      label="Description"
      lazy-rules
      :rules="[
        (val) => (val && val.length > 0) || 'Please enter product description',
      ]"
    />

    <div class="row">
      <div class="col mr-5">
        <q-select
          v-model="formState.type"
          :options="productTypes"
          label="Type"
        />
      </div>
      <div class="col">
        <q-input
          min="1"
          type="number"
          v-model="formState.quantity"
          label="Quantity"
          lazy-rules
          :rules="[
            (val) => (val && val.length > 0) || 'Please enter product quantity',
          ]"
        />
      </div>
    </div>

    <ImagePicker :onImageChange="onImageChange" />

    <q-checkbox
      v-model="formState.listed"
      label="List your product to the market or keep it private for trades"
    />
    <br />
    <q-btn v-if="id" color="red" icon="o_delete" label="Delete product" />
  </q-form>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import ImagePicker from "@/components/UI/ImagePicker.vue";
import { productTypes } from "@/constants/product";

export default defineComponent({
  name: "CreateProduct",
  props: ["formState", "id", "formRef", "onImageChange"],
  components: {
    ImagePicker,
  },
  setup(props) {
    function onImageChange(image: File) {
      props.onImageChange(image);
    }

    return {
      productTypes,
      onImageChange,
    };
  },
});
</script>
