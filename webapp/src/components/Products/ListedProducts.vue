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
import { defineComponent, computed } from "vue";
import { useStore } from "vuex";
import { ProductType } from "@/types/product";
import { useRouter } from "vue-router";
import { productColumns } from "@/constants/product";

export default defineComponent({
  name: "ListedProducts",
  setup() {
    const store = useStore();
    const router = useRouter();

    const reformedProducts = computed(() => {
      return store.getters.getUserListedProducts.map((product: ProductType) => {
        return {
          id: product.id,
          name: product.name,
          quantity: product.quantity,
          type: product.type,
          exchangeFor: product.exchangeFor?.name,
          createdAt: product.createdAt,
        };
      });
    });

    return {
      productColumns,
      reformedProducts,
      goToProduct(event: any, row: any) {
        router.push(`/profile/products/${row.id}`);
      },
    };
  },
});
</script>

<style></style>
