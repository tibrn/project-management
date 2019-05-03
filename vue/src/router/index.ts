import Vue from "vue";
import Meta from "vue-meta";
import Router from "vue-router";
import { sync } from "vuex-router-sync";
import routes from "./routes";
import store from "@/store/";
import { Route, NavigationGuard } from "vue-router";
import after from "@/middleware/after";
// init vue-meta, vue-router
Vue.use(Meta);
Vue.use(Router);

// create router
const router = new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: routes
});
// add router to store
sync(store, router);

router.afterEach(async (to: Route, from: Route) => {
  console.log("CEVA");
});

// global guards
router.beforeEach(async (to: Route, from: Route, next) => {
  if (!store.getters["is_init"]) {
    await store.dispatch("INIT_STORE");
  }
});

export default router;
