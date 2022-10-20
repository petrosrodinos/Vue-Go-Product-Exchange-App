<template>
  <AppBar />
</template>

<script>
import AppBar from "./components/AppBar";

export default {
  name: "App",
  beforeCreate() {
    this.$store.commit("autoLoginUser");
  },
  components: {
    AppBar,
  },

  setup() {
    return {};
  },
};
</script>
