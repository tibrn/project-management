import Vue from "vue";
import Meta from "vue-meta";
import Router from "vue-router";
import { sync } from "vuex-router-sync";
import routes from "./routes";
import store from "@/store/";
import { Route } from "vue-router";
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

// global guards
// router.beforeEach((to: Route, from: Route, next) => {});

// router.afterEach((to: Route, from: Route, next) => {});

export default router;
