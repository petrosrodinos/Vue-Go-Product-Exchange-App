<template>
  <Banner
    :show="!products && !loading"
    :loading="loading"
    message="There are no products to show"
  />
  <div v-if="products" class="row items-start">
    <div class="row">
      <div
        class="q-gutter-md my-card col-xs-12 col-sm-11 col-md-6 col-lg-4 mr-5 mb-5"
        v-for="product in products"
        :key="product?.name"
      >
        <Product :product="product" :fromProfile="false" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, computed } from "vue";
import Product from "@/components/UI/Product.vue";
import Banner from "@/components/UI/Banner.vue";
import { useStore } from "vuex";
import { ProductType } from "@/types/product";

export default defineComponent({
  name: "Home",
  components: {
    Product,
    Banner,
  },
  setup() {
    const store = useStore();
    const products = ref<ProductType[] | null>(null);

    onMounted(async () => {
      await store.dispatch("getProducts").then((res) => {
        if (res) {
          products.value = res;
        }
      });
    });

    return {
      products,
      loading: computed(() => store.state.product.loading),
    };
  },
});
</script>

<style>
.my-card {
  width: 100%;
  max-width: 250px;
}
</style>
