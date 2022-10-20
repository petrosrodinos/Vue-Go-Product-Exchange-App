import { createStore } from "vuex";
import authActions from "./actions/auth";
import productActions from "./actions/product";
import userActions from "./actions/user";
import tradeActions from "./actions/trade";

const store = createStore({
  modules: {
    auth: authActions,
    product: productActions,
    user: userActions,
    trade: tradeActions,
  },
  mutations: {
    autoLoginUser(state) {
      if (localStorage.getItem("store")) {
        this.replaceState(
          Object.assign(state, JSON.parse(localStorage.getItem("store") || ""))
        );
      }
    },
  },
});

store.subscribe((mutation, state) => {
  localStorage.setItem("store", JSON.stringify(state));
});

export default store;
