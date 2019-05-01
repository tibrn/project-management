import authRoutes from "@/router/parts/auth";
const NotFound = () =>
  import(/* webpackChunkName: "404" */ "@/views/404/404.vue");

export default [
  ...authRoutes,
  // 404
  {
    path: "*",
    name: "NotFound",
    component: NotFound
  }
];
