<template>
  <q-table
    title="Your Cancelled Products"
    :columns="CancelledColumns"
    :rows="products"
    row-key="name"
  />
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from "vue";
import { useStore } from "vuex";
import { CancelledColumns } from "@/constants/product";

export default defineComponent({
  name: "CancelledProducts",
  setup() {
    const store = useStore();
    const products = ref([]);

    onMounted(() => {
      store
        .dispatch("getTrades", "cancelled")
        .then((res) => {
          if (res) {
            products.value = res;
          }
        })
        .catch((err) => {
          console.log(err);
        });
    });

    return {
      CancelledColumns,
      products,
    };
  },
});
</script>

<style></style>
