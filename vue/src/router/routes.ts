import authRoutes from "@/router/parts/auth";
import projectsRoutes from "@/router/parts/projects";
import tasksRoutes from "@/router/parts/tasks";
import accountRoutes from "@/router/parts/account";
import VueRouterMultiguard from "vue-router-multiguard";
import route404 from "@/middleware/modules/404";
import before from "@/middleware/before";
const NotFound = () =>
  import(/* webpackChunkName: "404" */ "@/views/404/404.vue").then(
    m => m.default
  );

export default [
  ...authRoutes,

  {
    path: "/dashboard",
    name: "dashboard",
    meta: {
      layout: "app-layout"
    },
    beforeEnter: VueRouterMultiguard([]),
    children: [...projectsRoutes, ...tasksRoutes, ...accountRoutes]
  },
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
