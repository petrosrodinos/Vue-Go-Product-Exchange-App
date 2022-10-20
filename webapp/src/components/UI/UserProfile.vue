<template>
  <q-card class="profile-container">
    <q-card-section>
      <div class="text-h5 q-mt-sm q-mb-xs">
        <b>First Name:</b> {{ user.firstName }}
      </div>
      <div class="text-h5 q-mt-sm q-mb-xs">
        <b>Last Name:</b> {{ user.lastName }}
      </div>
    </q-card-section>

    <q-card-section class="q-pt-none">
      <div class="text-h5 q-mt-sm q-mb-xs">
        <b>Phone Number:</b> {{ user.phoneNumber }}
      </div>
      <div class="text-h5 q-mt-sm q-mb-xs"><b>Email:</b> {{ user.email }}</div>
      <div class="text-h5 q-mt-sm q-mb-xs">
        <b>Address:</b> {{ user.address.area }},
        {{ user.address.city }}
      </div>
    </q-card-section>
    <q-separator inset />
  </q-card>
</template>

<script>
export default {
  name: "UserProfile",
  props: ["user"],

  setup() {
    return {};
  },
};
</script>

<style>
.profile-container {
  word-break: break-word;
  width: 100%;
}
</style>
