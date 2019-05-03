import store from "@/store/";
import { Route } from "vue-router";
export default (to: Route, from: Route) => {
  console.log("AFTER", to, from);
  if (to.meta.layout && to.meta.layout.length) {
    // update layout only if it's a different layout
    if (to.meta.layout !== store.getters["layouts/layout"]) {
      store.commit("layouts/SET_LAYOUT", to.meta.layout);
    }
  } else {
    store.commit("layouts/SET_LAYOUT", store.getters["layouts/default"]);
  }
};
