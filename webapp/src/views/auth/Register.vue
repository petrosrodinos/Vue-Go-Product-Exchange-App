<template>
  <q-card class="q-pa-md max-width" style="max-width: 700px">
    <q-form autofocus @submit="onSubmit" class="q-gutter-md">
      <q-card-section>
        <div class="text-h6">{{ edit ? "Edit" : "Register" }}</div>
        <div v-if="!edit" class="text-subtitle1">
          Fill out the following form to create your account.
        </div>
      </q-card-section>
      <q-separator inset />

      <div class="row">
        <div class="col mr-5">
          <q-input
            v-model="formState.firstName"
            label="First name"
            lazy-rules
            :rules="[
              (val) =>
                (val && val.length > 0) || 'Please enter your first name',
            ]"
          />
        </div>
        <div class="col">
          <q-input
            v-model="formState.lastName"
            label="Last name"
            lazy-rules
            :rules="[
              (val) => (val && val.length > 0) || 'Please enter your last name',
            ]"
          />
        </div>
      </div>

      <q-input
        v-if="!edit"
        type="number"
        v-model="formState.phoneNumber"
        label="Phone number"
        lazy-rules
        :rules="[
          (val) => (val && val.length > 0) || 'Please enter your phone number',
        ]"
      />

      <q-input
        v-if="!edit"
        type="email"
        label="Email"
        v-model="formState.email"
        :rules="[(val) => validateEmail(val) || 'Please enter a valid email.']"
      ></q-input>

      <div class="row">
        <div class="col mr-5">
          <q-input
            v-model="formState.address.city"
            label="City"
            lazy-rules
            :rules="[
              (val) => (val && val.length > 0) || 'Please enter your city',
            ]"
          />
        </div>

        <div class="col">
          <q-input
            v-model="formState.address.area"
            label="Area"
            lazy-rules
            :rules="[
              (val) => (val && val.length > 0) || 'Please enter your area',
            ]"
          />
        </div>
      </div>

      <div v-if="!edit" class="row">
        <div class="col mr-5">
          <q-input
            label="Password"
            v-model="formState.password"
            type="password"
            lazy-rules
            :rules="[(val) => validatePassword(val) || 'Password is not valid']"
          >
          </q-input>
        </div>
        <div class="col">
          <q-input
            label="Confirm Password"
            v-model="formState.confirmPassword"
            type="password"
            lazy-rules
            :rules="[
              (val) =>
                (val && val === formState.password) || 'Passwords should match',
            ]"
          >
          </q-input>
        </div>
      </div>

      <div>
        <q-btn
          :loading="requestInfo.isLoading"
          :label="edit ? 'Save' : 'Register'"
          type="submit"
          color="primary"
        />
        <q-btn
          label="Reset"
          type="reset"
          color="primary"
          flat
          class="q-ml-sm"
        />
      </div>
    </q-form>
  </q-card>
</template>

<script lang="ts">
import { reactive, defineComponent } from "vue";
import { RegisterState } from "@/types/auth";
import { validateEmail, validatePassword } from "@/utils/validation";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import { displayErrorMessage } from "@/utils/errorMessage";

export default defineComponent({
  name: "Register",
  props: ["edit"],

  setup() {
    const router = useRouter();
    const store = useStore();
    const requestInfo = reactive({ isLoading: false });

    const formState = reactive<RegisterState>({
      firstName: "",
      lastName: "",
      phoneNumber: "",
      email: "",
      address: { city: "", area: "" },
      password: "",
      confirmPassword: "",
    });

    function onSubmit() {
      store
        .dispatch("registerUser", formState)
        .then((res) => {
          if (res) {
            router.push("/");
          }
        })
        .catch((err) => {
          displayErrorMessage("Could not register user");
        });
      requestInfo.isLoading = false;
    }

    return {
      formState,
      validateEmail,
      validatePassword,
      onSubmit,
      requestInfo,
    };
  },
});
</script>
