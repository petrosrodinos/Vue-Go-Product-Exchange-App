import axios from "axios";
import { displayErrorMessage } from "@/utils/errorMessage";

const API_URL = process.env.VUE_APP_API_URL;

export default {
  state: {
    token: null,
    id: null,
  },
  mutations: {
    setAuthState(state, payload) {
      state.token = payload.token;
      state.id = payload.id;
    },
    logoutUser(state) {
      state.token = null;
      state.id = null;
    },
    // autoLoginUser(state) {
    //   const store = JSON.parse(localStorage.getItem("store"));
    //   if (store) {
    //     state.token = store.token;
    //     state.id = store.id;
    //   }
    // },
  },
  actions: {
    async loginUser({ commit, dispatch }, payload) {
      try {
        const res = await axios.post(`${API_URL}auth/login`, payload);
        commit("setAuthState", { id: res.data.id, token: res.data.token });
        if (res.status === 200) {
          dispatch("getProductsByUserId");
          return true;
        }
        displayErrorMessage("Invalid Credentials");
        return false;
      } catch (error) {
        displayErrorMessage("Could not login please try later");
        return false;
      }
    },
    async registerUser({ commit }, payload) {
      try {
        const res = await axios.post(`${API_URL}auth/register`, payload);
        commit("setAuthState", { id: res.data.id, token: res.data.token });
        if (res.status === 200) {
          return true;
        }
        displayErrorMessage("Could not register user");

        return false;
      } catch (error) {
        displayErrorMessage("Could not create account please try later");
        return false;
      }
    },
  },
  getters: {
    getToken(state) {
      return state.token;
    },
  },
};
