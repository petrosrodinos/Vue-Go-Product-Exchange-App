<template>
  <q-card class="q-pa-md max-width">
    <q-form autofocus @submit="onSubmit" class="q-gutter-md">
      <q-card-section>
        <div class="text-h6">Change your Email</div>
      </q-card-section>
      <q-separator inset />

      <div v-if="!emailSent">
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
          label="New Email"
          v-model="email"
          :rules="[
            (val) => validateEmail(val) || 'Please enter a valid email.',
          ]"
        ></q-input>
      </div>

      <q-input
        v-if="emailSent"
        label="Enter verification code"
        v-model="otp"
        lazy-rules
        :rules="[(val) => (val && val.length > 0) || 'Please enter your code']"
      ></q-input>

      <q-btn label="Change Email" type="submit" color="primary" />
    </q-form>
  </q-card>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { validateEmail } from "@/utils/validation";

export default defineComponent({
  name: "ChangeEmail",
  setup() {
    const password = ref(null);
    const email = ref(null);
    const otp = ref(null);
    const emailSent = ref(false);

    return {
      password,
      email,
      otp,
      emailSent,
      validateEmail,
      onSubmit() {
        emailSent.value = true;
      },
    };
  },
});
</script>
