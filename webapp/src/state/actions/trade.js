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
    async sendTradeRequest({ commit, rootState }, payload) {
      try {
        commit("setLoading", true);
        const res = await axios.post(
          `${API_URL}trade`,
          {
            ...payload,
            userId: rootState.auth.id,
          },
          { ...getHeaders(rootState.auth.token) }
        );
        if (res.status === 200) {
          commit("setLoading", false);
          return res.data;
        }
        displayErrorMessage("Could create a trade request");
        commit("setLoading", false);

        return false;
      } catch (error) {
        commit("setLoading", false);
        displayErrorMessage("Could create a trade request");
        return false;
      }
    },
    async getTrades({ commit, rootState }, payload) {
      try {
        commit("setLoading", true);
        const res = await axios.get(
          `${API_URL}trade/${rootState.auth.id}/${payload}`,
          { ...getHeaders(rootState.auth.token) }
        );
        if (res.status === 200) {
          commit("setLoading", false);
          return res.data.trades.map((product) => {
            return {
              id: product.id,
              name1: "product.productId1.name",
              name2: "product.product2.name",
              quantity1: product.quantity1,
              quantity2: product.quantity2,
              traderName: "product.traderId.firstName",
              traderPhone: "product.traderId.phoneNumber",
              userId: product.userId,
              traderId: product.traderId,
              status: product.status,
            };
          });
        }
        displayErrorMessage(`Could not get ${payload} trades`);
        commit("setLoading", false);

        return false;
      } catch (error) {
        if (error.response.status !== 404) {
          displayErrorMessage(`Could not get ${payload} trades`);
        }
        commit("setLoading", false);
        return false;
      }
    },
    async changeTradeStatus({ commit }, trade) {
      try {
        commit("setLoading", true);
        const res = await axios.put(
          `${API_URL}trade/${trade.id}/${trade.status}`,
          { ...getHeaders(rootState.auth.token) }
        );
        if (res.status === 200) {
          commit("setLoading", false);
          displayErrorMessage(`Trade set to ${trade.status}`, "green");
          return;
        }
        displayErrorMessage(`Could not get ${payload} trades`, "green");
        commit("setLoading", false);

        return false;
      } catch (error) {
        commit("setLoading", false);
        displayErrorMessage("Something went wrong setting trade status");
        return false;
      }
    },
    async cancellTrade({ commit }, paylaod) {
      try {
        commit("setLoading", true);
        const res = await axios.delete(`${API_URL}trade/${paylaod}`, {
          ...getHeaders(rootState.auth.token),
        });
        if (res.status === 200) {
          commit("setLoading", false);
          displayErrorMessage("Trade request cancelled", "green");
          return;
        }
        displayErrorMessage("Could not cancel trade request");
        commit("setLoading", false);

        return false;
      } catch (error) {
        commit("setLoading", false);
        displayErrorMessage("Could not cancel trade request,please try later");
        return false;
      }
    },
  },
};
