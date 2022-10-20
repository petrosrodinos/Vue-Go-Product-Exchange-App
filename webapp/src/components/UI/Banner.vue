<template>
  <q-spinner v-if="loading" color="primary" size="3em" :thickness="2" />

  <q-banner
    v-if="show"
    :style="{ width: '100%' }"
    rounded
    class="bg-orange text-white"
  >
    {{ message }}
  </q-banner>
</template>

<script>
export default {
  name: "Banner",
  props: ["message", "show", "loading"],
};
</script>

<style></style>
