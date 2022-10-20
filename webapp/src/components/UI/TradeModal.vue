<template>
  <q-dialog v-model="prompt">
    <q-card class="trade-modal__container">
      <q-card-section>
        <div class="text-h6">Select what you want to trade with</div>
      </q-card-section>
      <q-form autofocus class="q-gutter-md" @submit="onSubmit">
        <q-card-section class="q-pt-none">
          <q-input
            min="1"
            type="number"
            v-model="tradeForm.quantity1"
            label="Quantity you want to exchange"
            lazy-rules
            :rules="[
              (val) =>
                (val && val.length > 0) || 'Please enter product  quantity',
            ]"
          />
          <q-select
            @update:model-value="getProductId"
            v-model="tempProduct"
            :options="userProducts"
            label="Your product you want to exchange with"
            lazy-rules
            :rules="[
              (val) => (val && val.length > 0) || 'Please select a product',
            ]"
          />
          <q-input
            min="1"
            type="number"
            v-model="tradeForm.quantity2"
            label="Quantity of your product you want to exchange with"
            lazy-rules
            :rules="[
              (val) =>
                (val && val.length > 0) || 'Please enter product quantity',
            ]"
          />
        </q-card-section>

        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="send Trade request" type="submit" />
          <q-btn flat label="Cancel" @click="closeModal" color="red" />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
</template>

<script>
import { computed, reactive, ref } from "vue";
import { useStore } from "vuex";
import { displayErrorMessage } from "@/utils/errorMessage";

export default {
  name: "TradeModal",
  props: ["product", "show", "closeModal"],
  setup(props) {
    const store = useStore();
    const tempProduct = ref("");
    const tradeForm = reactive({
      traderId: props.product.userId,
      productId1: props.product.id,
      productId2: "",
      quantity1: "1",
      quantity2: "1",
    });

    async function onSubmit() {
      if (!store.state.auth.token) {
        displayErrorMessage("You need to be logged in to trade");
        return;
      }
      await store
        .dispatch("sendTradeRequest", tradeForm)
        .then((res) => {
          if (res) {
            props.closeModal();
            displayErrorMessage("Trade request sent", "green");
          }
        })
        .catch((err) => {
          displayErrorMessage("Could not make trade request");
        });
    }

    function getUserProducts() {
      let prods = store.getters.getUserProductsNames;
      if (prods.length > 0) {
        return prods;
      }
      displayErrorMessage("You have no products to trade with");

      return [];
    }

    function getProductId(val) {
      const product = store.getters.getUserProductByName(val);
      tradeForm.productId2 = product.id;
    }

    return {
      prompt: computed(() => props.show),
      userProducts: computed(getUserProducts),
      tradeForm,
      onSubmit,
      getProductId,
      tempProduct,
    };
  },
};
</script>

<style>
.trade-modal__container {
  width: 450px;
  max-width: 100%;
  width: 100%;
}
</style>
