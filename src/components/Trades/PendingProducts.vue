<template>
  <div class="q-pa-md">
    <q-table
      title="Treats"
      :rows="products"
      :columns="TradeColumns"
      row-key="name"
      :loading="loading"
    >
      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td key="name1" :props="props">
            {{ props.row.name1 }}
          </q-td>
          <q-td key="name2" :props="props">
            {{ props.row.name2 }}
          </q-td>
          <q-td key="name2" :props="props">
            {{ props.row.quantity1 }}
          </q-td>
          <q-td key="name2" :props="props">
            {{ props.row.quantity2 }}
          </q-td>
          <q-td key="name2" :props="props">
            {{ props.row.traderName }}
          </q-td>
          <q-td key="name2" :props="props">
            {{ props.row.traderPhone }}
          </q-td>
          <q-td key="name2" :props="props">
            <div v-if="props.row.traderId === userId">
              <q-btn
                v-if="
                  props.row.status === 'cancelled' ||
                  props.row.status === 'pending'
                "
                @click="changeTradeStatus(props.row, 'completed')"
                :loading="loading"
                class="mr-5"
                round
                outline
                color="green"
                icon="o_check_outlined"
              />
              <q-btn
                :loading="loading"
                round
                outline
                color="red"
                icon="o_clear_outlined"
                @click="changeTradeStatus(props.row, 'cancelled')"
                v-if="
                  props.row.status === 'completed' ||
                  props.row.status === 'pending'
                "
              />
            </div>
            <div v-else>
              <q-btn
                :loading="loading"
                @click="cancellTrade(props.row)"
                class="mr-5"
                round
                outline
                color="red"
                icon="o_clear_outlined"
              />
            </div>
          </q-td>
        </q-tr>
      </template>
    </q-table>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from "vue";
import { useStore } from "vuex";
import { TradeColumns } from "@/constants/product";
import { Trade } from "@/types/trade";
import { computed } from "@vue/reactivity";

export default defineComponent({
  name: "PendingProducts",
  setup() {
    const store = useStore();
    const products = ref<Trade[]>([]);

    onMounted(() => {
      store
        .dispatch("getTrades", "pending")
        .then((res) => {
          if (res) {
            products.value = res;
          }
        })
        .catch((err) => {
          console.log(err);
        });
    });

    function changeTradeStatus(
      row: Trade,
      state: "cancelled" | "completed" | "pending"
    ) {
      store.dispatch("changeTradeStatus", { ...row, status: state });
      row.status = state;
      let index = products.value.findIndex((item: Trade) => item.id === row.id);
      products.value.splice(index, 1, row);
    }

    function cancellTrade(row: Trade) {
      store.dispatch("cancellTrade", row.id);
      let index = products.value.findIndex((item: Trade) => item.id === row.id);
      products.value.splice(index, 1);
    }

    return {
      TradeColumns,
      products,
      changeTradeStatus,
      cancellTrade,
      userId: store.state.auth.id,
      loading: computed(() => store.state.trade.loading),
    };
  },
});
</script>

<style></style>
