<template>
  <q-layout view="hHh Lpr fFf">
    <q-header elevated class="bg-primary text-white">
      <q-toolbar>
        <q-btn
          v-if="$q.screen.xs || $q.screen.sm"
          dense
          flat
          round
          icon="o_menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title> Product Exchange </q-toolbar-title>
        <SearchBar />
        <div v-if="!token">
          <q-btn
            :to="{ name: 'Login' }"
            color="primary"
            label="Login"
            class="mr-5"
          />
          <q-btn
            :to="{ name: 'Register' }"
            outline
            color="white"
            label="Register"
          />
        </div>

        <q-btn
          v-if="token"
          @click="logout"
          to="/"
          outline
          color="white"
          label="Logout"
        />
      </q-toolbar>
    </q-header>

    <q-drawer
      persistent
      v-model="leftDrawerOpen"
      show-if-above
      :width="250"
      :breakpoint="700"
      bordered
      class="bg-grey-3"
    >
      <Links />
    </q-drawer>

    <q-page-container>
      <Container>
        <template v-slot:centered>
          <router-view />
        </template>
      </Container>
    </q-page-container>
  </q-layout>
</template>

<script>
import { defineComponent, ref } from "vue";
import Links from "./Links";
import SearchBar from "../UI/SearchBar.vue";
import Container from "../UI/Container.vue";
import { mapState } from "vuex";
import { useStore } from "vuex";

export default defineComponent({
  name: "AppBar",
  components: { Container, Links, SearchBar },
  setup() {
    const leftDrawerOpen = ref(true);
    const store = useStore();
    return {
      leftDrawerOpen,
      logout() {
        store.commit("logoutUser");
        localStorage.removeItem("store");
      },
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value;
      },
    };
  },
  computed: {
    ...mapState({
      token: (state) => state.auth.token,
    }),
  },
});
</script>
