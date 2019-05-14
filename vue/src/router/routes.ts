import authRoutes from "@/router/parts/auth";
import projectsRoutes from "@/router/parts/projects";
import tasksRoutes from "@/router/parts/tasks";
import accountRoutes from "@/router/parts/account";
import VueRouterMultiguard from "vue-router-multiguard";
import route404 from "@/middleware/modules/404";
import auth from "@/middleware/modules/auth";
const NotFound = () =>
  import(/* webpackChunkName: "404" */ "@/views/404/404.vue").then(
    m => m.default
  );

const Main = () =>
  import(/* webpackChunkName: "404" */ "@/views/Main.vue").then(m => m.default);

export default [
  ...authRoutes,

  {
    path: "/dashboard",
    name: "dashboard",
    component: Main,
    meta: {
      layout: "app-layout"
    },
    beforeEnter: VueRouterMultiguard([auth]),
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
