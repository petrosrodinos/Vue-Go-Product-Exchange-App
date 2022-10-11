<template>
  <q-table
    title="Your Exchanged Products"
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
  name: "ExchangedProducts",
  setup() {
    const store = useStore();
    const products = ref([]);

    onMounted(() => {
      store
        .dispatch("getTrades", "completed")
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
