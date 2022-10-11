/* eslint-disable */
import VueRouter, { Route } from "vue-router";
import { store } from "./types/store";
import { Store } from "vuex";

declare module "@vue/runtime-core" {
  interface ComponentCustomProperties {
    $store: Store<store>;
    $router: VueRouter | any;
    $refs: any;
  }
}
