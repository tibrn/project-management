import Vue from "vue";
import Meta from "vue-meta";
import Router, { Route } from "vue-router";
import { sync } from "vuex-router-sync";
import routes from "./routes";
import store from "@/store/";
// init vue-meta, vue-router
Vue.use(Meta);
Vue.use(Router);
// create router
const router = new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: routes
});

const NOT_INIT = ["login", "register"];
// add router to store
sync(store, router);
router.beforeEach(async (to: Route, from: Route, next) => {
  if (!store.getters["is_init"]) {
    await store.dispatch("INIT_STORE");
  }

  if (to.meta.layout && to.meta.layout.length) {
    // update layout only if it's a different layout
    if (to.meta.layout !== store.getters["layouts/layout"]) {
      store.commit("layouts/SET_LAYOUT", to.meta.layout);
    }
  } else {
    store.commit("layouts/SET_LAYOUT", store.getters["layouts/default"]);
  }
  //DON'T FORGET NEXT
  return next();
});

export default router;
