import axios from "axios";
import { displayErrorMessage } from "@/utils/errorMessage";

const API_URL = process.env.VUE_APP_API_URL;

const getHeaders = (token) => {
  return {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };
};

export default {
  state() {
    return {
      loading: false,
    };
  },
  mutations: {
    setLoading(state, payload) {
      state.loading = payload;
    },
  },
  actions: {
    async uploadPhoto({ commit }, payload) {
      try {
        commit("setLoading", true);
        let image = new FormData();
        image.append("file", payload);
        const res = await axios.post(
          `${API_URL}products/file/formdata`,
          image,
          {
            headers: { "Content-Type": "multipart/form-data" },
          }
        );
        if (res.data.statusCode === 200) {
          commit("setLoading", false);
          return res.data.data.data;
        }
        displayErrorMessage("Could not upload image");
        commit("setLoading", false);

        return false;
      } catch (error) {
        displayErrorMessage("Please choose a valid image");
        commit("setLoading", false);

        return false;
      }
    },
    async createProduct({ commit, rootState, dispatch }, payload) {
      try {
        commit("setLoading", true);
        const res = await axios.post(
          `${API_URL}products`,
          {
            ...payload,
            userId: rootState.auth.id,
          },
          { ...getHeaders(rootState.auth.token) }
        );
        if (res.status === 200) {
          commit("setLoading", false);
          dispatch("getProductsByUserId");
          return true;
        }
        displayErrorMessage(
          "Could not create product,please check your inputs"
        );
        commit("setLoading", false);
        return false;
      } catch (error) {
        displayErrorMessage("Could not create product,please try again later");
        commit("setLoading", false);
        return false;
      }
    },
    async getProducts({ commit }) {
      try {
        commit("setLoading", true);
        const res = await axios.get(`${API_URL}products`);
        if (res.status === 200) {
          commit("setLoading", false);
          return res.data.products;
        }
        return null;
      } catch (error) {
        if (error.response.status === 404) {
          displayErrorMessage(`There are no products yet`, "info");
          commit("setLoading", false);
          return;
        }
        displayErrorMessage("Could not fetch products,try again later");
        commit("setLoading", false);
        return null;
      }
    },
  },
};
