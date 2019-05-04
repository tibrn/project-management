import VueRouterMultiguard from "vue-router-multiguard";
const Login = () =>
  import(/* webpackChunkName: "auth" */ "@/views/auth/Login.vue");
const Register = () =>
  import(/* webpackChunkName: "auth" */ "@/views/auth/Register.vue");
import guest from "@/middleware/modules/guest";
export default [
  {
    path: "/login",
    name: "login",
    component: Login,
    meta: {
      layout: "auth-layout"
    },
    beforeEnter: VueRouterMultiguard([guest])
  },
  {
    path: "/register",
    name: "register",
    component: Register,
    meta: {
      layout: "auth-layout"
    },
    beforeEnter: VueRouterMultiguard([guest])
  }
];
