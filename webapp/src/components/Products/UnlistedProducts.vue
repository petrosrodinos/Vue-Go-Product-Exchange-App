<template>
  <q-table
    @row-click="goToProduct"
    title="Your listed Products"
    :rows="reformedProducts"
    :columns="productColumns"
    row-key="name"
  />
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, computed } from "vue";
import { useStore } from "vuex";
import { ProductType } from "@/types/product";
import { useRouter } from "vue-router";
import { productColumns } from "@/constants/product";

export default defineComponent({
  name: "UnlistedProducts",
  setup() {
    const store = useStore();
    const router = useRouter();

    const reformedProducts = computed(() => {
      return store.getters.getUserUnlistedProducts.map(
        (product: ProductType) => {
          return {
            id: product.id,
            name: product.name,
            quantity: product.quantity,
            type: product.type,
            exchangeFor: product.exchangeFor?.name,
            createdAt: product.createdAt,
          };
        }
      );
    });

    return {
      reformedProducts,
      productColumns,
      goToProduct(event: any, row: any) {
        router.push(`/profile/products/${row.id}`);
      },
    };
  },
});
</script>

<style></style>
