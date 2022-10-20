<template>
  <q-card class="q-pa-md max-width">
    <q-form autofocus @submit="onSubmit" class="q-gutter-md">
      <q-card-section>
        <div class="text-h6">Delete your profile</div>
      </q-card-section>
      <q-separator inset />

      <q-input
        type="password"
        v-model="password"
        label="Password"
        lazy-rules
        :rules="[
          (val) => (val && val.length > 0) || 'Please enter your password',
        ]"
      />

      <q-btn
        color="red"
        type="submit"
        icon="o_delete"
        label="Delete my profile"
      />
    </q-form>
  </q-card>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";

export default defineComponent({
  name: "DeleteProfile",
  setup() {
    const password = ref(null);

    return {
      password,
      onSubmit() {},
    };
  },
});
</script>
