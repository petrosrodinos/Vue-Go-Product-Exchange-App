<template>
  <div class="q-pa-md max-width" style="max-width: 800px">
    <div class="q-gutter-y-md">
      <q-card>
        <q-tabs
          v-model="tab"
          dense
          class="text-grey"
          active-color="primary"
          indicator-color="primary"
          align="justify"
          narrow-indicator
        >
          <q-tab name="information" label="Information" />
          <q-tab name="password" label="Password" />
          <q-tab name="email" label="Email" />
          <q-tab name="phone" label="Phone number" />
          <q-tab name="deleteprofile" label="Delete profile" />
        </q-tabs>

        <q-separator />

        <q-tab-panels v-model="tab" animated>
          <q-tab-panel name="information">
            <Register :edit="true" />
          </q-tab-panel>

          <q-tab-panel name="password">
            <ChangePassword />
          </q-tab-panel>

          <q-tab-panel name="email">
            <ChangeEmail />
          </q-tab-panel>

          <q-tab-panel name="phone">
            <ChangePhone />
          </q-tab-panel>

          <q-tab-panel name="deleteprofile">
            <DeleteProfile />
          </q-tab-panel>
        </q-tab-panels>
      </q-card>
    </div>
  </div>
</template>

<script>
import { ref, defineComponent } from "vue";
import Register from "../auth/Register.vue";
import ChangeEmail from "../../components/Settings/ChangeEmail.vue";
import ChangePassword from "../../components/Settings/ChangePassword.vue";
import ChangePhone from "../../components/Settings/ChangePhone.vue";
import DeleteProfile from "../../components/Settings/DeleteProfile.vue";

export default defineComponent({
  name: "Settings",
  components: {
    Register,
    ChangeEmail,
    ChangePassword,
    ChangePhone,
    DeleteProfile,
  },
  setup() {
    return {
      tab: ref("information"),
    };
  },
});
</script>
