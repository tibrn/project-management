import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import store from "@/store";
Vue.use(Vuex);
const LOCATION = "store/index";
// Load store modules dynamically.
const requireContext = require.context("@/store/modules/", false, /.*\.ts$/);

const modules = requireContext
  .keys()
  .map(file => [file.replace(/(^.\/)|(\.ts$)/g, ""), requireContext(file)])
  .reduce((modules, [name, module]) => {
    if (module.namespaced === undefined) {
      module.namespaced = true;
    }

    return { ...modules, [name]: module };
  }, {});

export default new Vuex.Store({
  state: {
    is_init: false,
    is_serve_mode: false
  },

  getters: {
    is_init: state => state.is_init,
    is_serve_mode: state => state.is_init
  },

  mutations: {
    INIT: state => {
      state.is_init = true;
    },
    IS_SERVE_MODE: state => {
      state.is_serve_mode = true;
    }
  },

  actions: {
    // init store
    async INIT_STORE() {
      let initData = false;
      // get state hydration data from inline script or via API call
      let windowData = <any>window;
      if (windowData.user_data) {
        initData = windowData.user_data;
      } else {
        try {
          const { data } = await axios.get("/api/auth/refresh");
          if (data) {
            initData = data;
          }
        } catch (e) {
          console.log(LOCATION, e);
        } finally {
          store.commit("IS_SERVE_MODE");
        }
      }
      // if user data is available, init state modules
      if (initData) {
        store.dispatch("INIT_STORE_DATA", initData);
      }
      // mark state as initialized
      store.commit("INIT");
    },
    // init store data
    INIT_STORE_DATA({ commit }, storeData) {
      commit("auth/INIT", storeData, { root: true });
      commit("user/INIT", storeData, { root: true });
    },

    // clean/remove store data
    RESET_STORE_DATA({ commit }) {
      commit("auth/RESET", { root: true });
      commit("user/RESET", { root: true });
    }
  },

  // store modules
  modules
});
