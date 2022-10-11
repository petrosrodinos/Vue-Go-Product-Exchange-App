<template>
  <q-card class="q-pa-md max-width">
    <q-form autofocus @submit="onSubmit" class="q-gutter-md">
      <q-card-section>
        <div class="text-h6">Change your Phone number</div>
      </q-card-section>
      <q-separator inset />

      <div v-if="!codeSent">
        <q-input
          type="password"
          v-model="password"
          label="Password"
          lazy-rules
          :rules="[
            (val) => (val && val.length > 0) || 'Please enter your password',
          ]"
        />

        <q-input
          type="number"
          label="New Phone number"
          v-model="phone"
          :rules="[
            (val) =>
              (val && val.length > 0) || 'Please enter your new phone number',
          ]"
        ></q-input>
      </div>

      <q-input
        v-if="codeSent"
        label="Enter verification code"
        v-model="otp"
        lazy-rules
        :rules="[(val) => (val && val.length > 0) || 'Please enter your code']"
      ></q-input>

      <q-btn label="Change Phone" type="submit" color="primary" />
    </q-form>
  </q-card>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";

export default defineComponent({
  name: "ChangePhone",
  setup() {
    const password = ref(null);
    const phone = ref(null);
    const otp = ref(null);
    const codeSent = ref(false);

    return {
      password,
      phone,
      otp,
      codeSent,
      onSubmit() {
        codeSent.value = true;
      },
    };
  },
});
</script>
