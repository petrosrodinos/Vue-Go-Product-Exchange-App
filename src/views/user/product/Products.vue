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
          <q-tab name="listed" label="Listed" />
          <q-tab name="unlisted" label="Unlisted" />
        </q-tabs>

        <q-separator />

        <q-tab-panels v-model="tab" animated>
          <q-tab-panel name="listed">
            <ListedProducts />
          </q-tab-panel>

          <q-tab-panel name="unlisted">
            <UnlistedProducts />
          </q-tab-panel>
        </q-tab-panels>
      </q-card>
    </div>
  </div>
</template>

<script>
import { ref, defineComponent } from "vue";
import ListedProducts from "@/components/Products/ListedProducts.vue";
import UnlistedProducts from "@/components/Products/UnlistedProducts.vue";

export default defineComponent({
  name: "Products",
  components: {
    ListedProducts,
    UnlistedProducts,
  },
  setup() {
    return {
      tab: ref("listed"),
    };
  },
});
</script>
