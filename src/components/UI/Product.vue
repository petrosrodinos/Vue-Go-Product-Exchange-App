<template>
  <q-card class="">
    <TradeModal
      :product="product"
      :show="showTradeModal"
      :closeModal="closeModal"
    />
    <img
      class="card-image"
      :src="
        product.image
          ? product.image
          : 'https://cdn.shopify.com/s/files/1/0996/9166/products/MysteryBox-100.jpg?v=1624306986'
      "
    />

    <q-card-section>
      <div class="text-h5 q-mt-sm q-mb-xs">{{ product.name }}</div>
      <div class="text-caption product-description">
        {{ product.description }}
      </div>
    </q-card-section>

    <q-card-section class="q-pt-none">
      Quantity: {{ product.quantity }}
    </q-card-section>
    <q-separator inset />

    <q-card-actions>
      <q-btn
        @click="showTradeModal = true"
        v-if="product.userId !== id"
        icon="o_change_circle_utlined"
        color="secondary"
        flat
        >Trade</q-btn
      >
      <q-btn
        v-if="!fromProfile"
        :to="{ name: 'Profile', params: { id: product.userId } }"
        icon="o_person2_outlined_icon"
        color="secondary"
        flat
        >Visit Profile</q-btn
      >

      <q-space />

      <q-btn
        color="grey"
        round
        flat
        dense
        :icon="expanded ? 'o_keyboard_arrow_up' : 'o_keyboard_arrow_down'"
        @click="expanded = !expanded"
        v-if="product.exchangeFor?.name"
      />
    </q-card-actions>
    <q-slide-transition v-if="product.exchangeFor?.name">
      <div v-show="expanded">
        <q-separator />
        <q-card-section class="text-subitle2">
          <div class="text-h6 q-mt-sm q-mb-xs">Exchange For</div>
          <q-separator inset />

          <div class="text-h5 q-mt-sm q-mb-xs">
            {{ product.exchangeFor.name }}
          </div>
          <div class="text-caption">
            Quantity: {{ product.exchangeFor.quantity }}
          </div>
        </q-card-section>
      </div>
    </q-slide-transition>
  </q-card>
</template>

<script>
import { ref } from "vue";
import { useStore } from "vuex";
import TradeModal from "./TradeModal.vue";

export default {
  name: "Product",
  props: ["product", "fromProfile"],
  components: {
    TradeModal,
  },
  setup() {
    const store = useStore();
    const showTradeModal = ref(false);

    return {
      showTradeModal,
      expanded: ref(false),
      id: store.state.auth.id,
      closeModal() {
        showTradeModal.value = false;
      },
    };
  },
};
</script>

<style scoped>
.product-description {
  word-break: break-all;
}
.card-image {
  height: 350px;
  object-fit: cover;
  max-width: 100%;
  width: 1000%;
}
</style>
