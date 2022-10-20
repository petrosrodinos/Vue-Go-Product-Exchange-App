import axios from "axios";
import { displayErrorMessage } from "@/utils/errorMessage";
const API_URL = process.env.VUE_APP_API_URL;

export default {
  state: {
    loading: false,
    products: [],
  },
  mutations: {
    setLoading(state, payload) {
      state.loading = payload;
    },
    setProducts(state, payload) {
      state.products = payload;
    },
  },
  actions: {
    async getListedProductsByUserId({ commit }, payload) {
      commit("setLoading", true);
      try {
        const res = await axios.get(
          `${API_URL}user/${payload}/products/listed`
        );
        if (res.status === 200) {
          commit("setLoading", false);
          return res.data.products;
        }
        commit("setLoading", false);
        displayErrorMessage("Could not fetch products");
        return null;
      } catch (error) {
        displayErrorMessage("Could not fetch products");
        return null;
      }
    },
    async getProductsByUserId({ rootState, commit }) {
      try {
        const res = await axios.get(
          `${API_URL}user/${rootState.auth.id}/products`
        );
        if (res.status === 200) {
          commit("setProducts", res.data.products);
          return res.data.products;
        }
        return null;
      } catch (error) {
        return null;
      }
    },
    async getUserById(state, payload) {
      try {
        const res = await axios.get(`${API_URL}user/${payload}`);
        if (res.status === 200) {
          return res.data.user;
        }
      } catch (error) {
        return null;
      }
    },
  },
  getters: {
    getUserListedProducts(state) {
      return state.products.filter((product) => product.listed);
    },
    getUserUnlistedProducts(state) {
      return state.products.filter((product) => !product.listed);
    },
    getUserProductsNames(state) {
      return state.products.map((product) => product.name);
    },
    getUserProductByName: (state) => (name) => {
      return state.products.find((product) => product.name === name);
    },
  },
};
