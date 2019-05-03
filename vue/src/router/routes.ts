import authRoutes from "@/router/parts/auth";
import VueRouterMultiguard from "vue-router-multiguard";
import route404 from "@/middleware/modules/404";
const NotFound = () =>
  import(/* webpackChunkName: "404" */ "@/views/404/404.vue");

export default [
  ...authRoutes,
  // 404
  {
    path: "*",
    name: "NotFound",
    component: NotFound,
    meta: {
      layout: "simple-layout"
    },
    beforeEnter: VueRouterMultiguard([route404])
  }
];
