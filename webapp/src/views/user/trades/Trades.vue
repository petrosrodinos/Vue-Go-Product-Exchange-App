<template>
  <div :style="{ maxWidth: '100%' }" class="q-pa-md max-width">
    <div class="q-gutter-y-md">
      <q-card>
        <q-tabs
          v-model="tab"
          dense
          class="text-grey"
          active-color="primary"
          indicator-color="primary"
          align="justify"
          narrow-indicator
        >
          <q-tab name="pending" label="Pending" />
          <q-tab name="exchanged" label="Completed" />
          <q-tab name="cancelled" label="Rejected" />
        </q-tabs>

        <q-separator />

        <q-tab-panels v-model="tab" animated>
          <q-tab-panel name="pending">
            <PendingProducts />
          </q-tab-panel>
          <q-tab-panel name="exchanged">
            <ExchangedProducts />
          </q-tab-panel>

          <q-tab-panel name="cancelled">
            <CancelledProducts />
          </q-tab-panel>
        </q-tab-panels>
      </q-card>
    </div>
  </div>
</template>

<script>
import { ref, defineComponent } from "vue";
import ExchangedProducts from "@/components/Trades/ExchangedProducts.vue";
import PendingProducts from "@/components/Trades/PendingProducts.vue";
import CancelledProducts from "@/components/Trades/CancelledProducts.vue";

export default defineComponent({
  name: "Trades",
  components: {
    ExchangedProducts,
    PendingProducts,
    CancelledProducts,
  },
  setup() {
    return {
      tab: ref("pending"),
    };
  },
});
</script>
