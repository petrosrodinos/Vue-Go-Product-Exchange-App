<template>
  <Banner
    :show="!products && !loading"
    :loading="loading"
    message="There are not products to show"
  />

  <div class="row items-start">
    <div
      class="max-width col col-sm-9 col-xs-12 col-md-9 col-lg-9"
      v-if="products"
    >
      <div class="row">
        <div
          class="q-gutter-md my-card col col-sm-6 col-xs-12 col-md-6 col-lg-4 mr-5 mb-5"
          v-for="product in products"
          :key="product.name"
        >
          <Product :product="product" :fromProfile="true" />
        </div>
      </div>
    </div>
    <div v-if="products" class="col col-sm-3 col-xs-12 col-md-3 col-lg-3">
      <UserProfile v-if="user" :user="user" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, computed } from "vue";
import Product from "@/components/UI/Product.vue";
import Banner from "@/components/UI/Banner.vue";
import UserProfile from "@/components/UI/UserProfile.vue";
import { useStore } from "vuex";
import { ProductType } from "@/types/product";

export default defineComponent({
  name: "Profile",
  props: ["id"],
  components: {
    Product,
    UserProfile,
    Banner,
  },
  setup(props) {
    const store = useStore();
    const products = ref<ProductType[] | null>(null);
    const user = ref();

    onMounted(async () => {
      await store
        .dispatch("getListedProductsByUserId", props.id)
        .then((res: any) => {
          if (res) {
            products.value = res;
          }
        });
      await store.dispatch("getUserById", props.id).then((res: any) => {
        user.value = res;
      });
    });

    return {
      products,
      user,
      loading: computed(() => store.state.product.loading),
    };
  },
});
</script>
