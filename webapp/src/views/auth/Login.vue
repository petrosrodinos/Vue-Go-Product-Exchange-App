<template>
  <q-card class="q-pa-md max-width" style="max-width: 700px">
    <q-form autofocus @submit="onSubmit" class="q-gutter-md">
      <q-card-section>
        <div class="text-h6">Log in to your account</div>
      </q-card-section>
      <q-separator inset />
      <q-input
        v-model="phone"
        label="Phone number"
        lazy-rules
        :rules="[
          (val) => (val && val.length > 0) || 'Please enter your phone number',
        ]"
      />

      <q-input
        type="password"
        v-model="password"
        label="Password"
        lazy-rules
        :rules="[
          (val) => (val && val.length > 0) || 'Please enter your password',
        ]"
      />

      <q-btn flat color="primary" to="/auth/reset-password">
        I forgot my password</q-btn
      >

      <div>
        <q-btn :loading="loading" label="Login" type="submit" color="primary" />
      </div>
    </q-form>
  </q-card>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { displayErrorMessage } from "../../utils/errorMessage";

export default defineComponent({
  name: "Login",
  setup() {
    const phone = ref(null);
    const password = ref(null);
    const loading = ref(false);
    const store = useStore();
    const router = useRouter();

    function onSubmit() {
      loading.value = true;
      store
        .dispatch("loginUser", {
          phoneNumber: phone.value,
          password: password.value,
        })
        .then((res) => {
          if (res) {
            router.push("/");
          }
        })
        .catch((err) => {
          displayErrorMessage("Invalid Credentials");
        });
      loading.value = false;
    }

    return {
      phone,
      password,
      onSubmit,
      loading,
    };
  },
});
</script>
