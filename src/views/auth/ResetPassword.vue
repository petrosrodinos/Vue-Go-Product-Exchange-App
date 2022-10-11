<template>
  <q-card class="q-pa-md max-width" style="max-width: 700px">
    <q-form autofocus @submit="onSubmit" class="q-gutter-md">
      <q-card-section>
        <div class="text-h6">Reset your password</div>
      </q-card-section>
      <q-separator inset />

      <q-input
        v-if="!emailSent"
        label="Email"
        v-model="email"
        :rules="[(val) => validateEmail(val) || 'Please enter a valid email.']"
      ></q-input>

      <q-input
        v-if="emailSent"
        label="Enter verification code"
        v-model="otp"
        lazy-rules
        :rules="[(val) => (val && val.length > 0) || 'Please enter your code']"
      ></q-input>

      <q-input
        v-if="emailSent"
        type="password"
        v-model="password"
        label="New Password"
        lazy-rules
        :rules="[
          (val) => (val && val.length > 0) || 'Please enter your new password',
        ]"
      />

      <q-btn label="Reset Password" type="submit" color="primary" />
    </q-form>
  </q-card>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { validateEmail } from "@/utils/validation";

export default defineComponent({
  name: "ResetPassword",
  setup() {
    const email = ref(null);
    const otp = ref(null);
    const password = ref(null);
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
