<template>
  <q-card class="q-pa-md login__container" style="max-width: 700px">
    <q-form autofocus @submit="onSubmit" class="q-gutter-md">
      <q-card-section>
        <div class="text-h6">Send us your feedback</div>
      </q-card-section>
      <q-separator inset />
      <q-input
        v-model="email"
        label="Email"
        lazy-rules
        :rules="[(val) => (val && val.length > 0) || 'Please enter your email']"
      />

      <q-input
        v-model="message"
        label="Message"
        hint="Please type your message"
        type="textarea"
        lazy-rules
        :rules="[
          (val) => (val && val.length > 0) || 'Please enter your message',
        ]"
      />
      <q-btn label="Send" type="submit" color="primary" />
    </q-form>
  </q-card>
</template>

<script>
import { defineComponent, ref } from "@vue/runtime-core";

export default defineComponent({
  name: "SendFeedBack",
  setup() {
    const email = ref(null);
    const message = ref(null);

    return {
      email,
      message,

      onSubmit() {},
    };
  },
});
</script>

<style>
.login__container {
  width: 100%;
}
</style>
