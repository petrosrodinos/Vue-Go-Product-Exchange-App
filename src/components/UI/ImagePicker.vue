<template>
  <q-file
    v-model="image"
    clearable
    label="Image"
    outlined
    accept="image/*"
    @update:model-value="handleUpload()"
  >
    <template v-slot:prepend>
      <q-icon name="o_broken_image_outlined_icon" />
    </template>
  </q-file>
  <br />
  <q-img
    v-if="imageUrl"
    :src="imageUrl"
    spinner-color="white"
    style="height: 140px; max-width: 150px"
  ></q-img>
  <br />
</template>

<script>
import { ref } from "vue";
export default {
  name: "ImagePicker",
  props: ["onImageChange"],
  setup(props) {
    const image = ref(null);
    const imageUrl = ref("");

    const handleUpload = () => {
      if (image.value) {
        imageUrl.value = URL.createObjectURL(image.value);
        props.onImageChange(image.value);
      }
    };

    return {
      image,
      imageUrl,
      handleUpload,
    };
  },
};
</script>

<style></style>
