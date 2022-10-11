<template>
  <q-card class="q-pa-md max-width">
    <q-form autofocus @submit="onSubmit" class="q-gutter-md">
      <q-card-section>
        <div class="text-h6">Change your password</div>
      </q-card-section>
      <q-separator inset />
      <q-input
        type="password"
        v-model="oldPassword"
        label="Old Password"
        lazy-rules
        :rules="[
          (val) => (val && val.length > 0) || 'Please enter your old password',
        ]"
      />

      <q-input
        type="password"
        v-model="newPassword"
        label="New Password"
        lazy-rules
        :rules="[(val) => validatePassword(val) || 'Password is not valid']"
      />

      <q-input
        lazy-rules
        label="Confirm Password"
        v-model="confirmPassword"
        type="password"
        :rules="[
          (val) => (val && val === newPassword) || 'Passwords should match',
        ]"
      >
      </q-input>

      <q-btn label="Change password" type="submit" color="primary" />
    </q-form>
  </q-card>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { validatePassword } from "@/utils/validation";

export default defineComponent({
  name: "ChangePassword",
  setup() {
    const oldPassword = ref(null);
    const newPassword = ref(null);
    const confirmPassword = ref(null);

    return {
      oldPassword,
      newPassword,
      confirmPassword,
      validatePassword,
      onSubmit() {
        console.log("submit");
      },
    };
  },
});
</script>
