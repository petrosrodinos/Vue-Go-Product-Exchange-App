<template>
  <q-scroll-area class="fit">
    <q-list>
      <template v-for="(menuItem, index) in menuList" :key="index">
        <q-item
          v-if="menuItem.auth === !!token || menuItem.auth === false"
          :to="menuItem.to"
          clickable
          :active="menuItem.label === 'Outbox'"
          v-ripple
        >
          <q-item-section avatar>
            <q-icon :name="menuItem.icon" />
          </q-item-section>
          <q-item-section>
            {{ menuItem.label }}
          </q-item-section>
        </q-item>
        <q-separator :key="'sep' + index" v-if="menuItem.separator" />
      </template>
    </q-list>
  </q-scroll-area>
</template>

<script lang="ts">
import { Link } from "../../types/links";
import { defineComponent } from "vue";
import { mapState } from "vuex";

const menuList: Link[] = [
  {
    icon: "o_home",
    label: "Home",
    to: "/",
    auth: false,
  },
  {
    icon: "o_change_circle_outlined",
    label: "Trades",
    to: "/profile/trades",
    auth: true,
  },
  {
    icon: "o_inventory_2_outlined_icon",
    label: "Your Products",
    to: "/profile/products",
    auth: true,
  },
  {
    icon: "o_add_circle_outline_outlined_icon",
    label: "Add Product",
    to: "/profile/products/add",
    auth: true,
  },
  {
    icon: "o_settings",
    label: "Settings",
    to: "/profile/settings",
    separator: true,
    auth: true,
  },
  {
    icon: "o_feedback",
    label: "Send Feedback",
    to: "/feedback",
    auth: false,
  },
  {
    icon: "o_info",
    label: "About",
    to: "/about",
    auth: false,
  },
];

export default defineComponent({
  name: "Links",
  setup() {
    return {
      menuList,
    };
  },
  computed: {
    ...mapState({
      token: (state: any) => state.auth.token,
    }),
  },
  methods: {
    redirect() {
      this.$router.push({ name: "Login" });
      // this.$router.go(-1);
    },
  },
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
