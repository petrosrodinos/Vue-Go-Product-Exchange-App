<template>
  <q-input
    label="Search..."
    dark
    dense
    standout
    v-model="text"
    class="q-ml-md mr-5"
  >
    <template v-slot:append>
      <q-icon v-if="text === ''" name="o_search" />
      <q-icon v-else name="o_clear" class="cursor-pointer" @click="text = ''" />
    </template>
  </q-input>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
export default defineComponent({
  name: "SearchBar",
  setup() {
    const text = ref("");
    return {
      text,
    };
  },
});
</script>

<style></style>
