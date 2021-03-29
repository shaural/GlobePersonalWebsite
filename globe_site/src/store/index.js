import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    country: undefined
  },
  getters: {
    country(state) {
      return state.country;
    }
  },
  mutations: {
    setCountry(state, country) {
      state.country = country;
    },
    unsetCountry(state) {
      state.country = undefined;
    }
  },
  actions: {},
  modules: {}
});
