<template>
  <AddProduct :id="id" />
</template>

<script>
import AddProduct from "@/views/user/product/AddProduct.vue";
export default {
  components: {
    AddProduct,
  },
  name: "EditProduct",
  props: ["id"],
};
</script>
