<template>
  <div class="column items-center justify-center centered__container">
    <slot name="centered"></slot>
  </div>
</template>

<script>
export default {
  name: "Container",
};
</script>

<style>
.centered__container {
  padding: 10px;
  width: 100%;
}
</style>
