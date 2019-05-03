const Login = () =>
  import(/* webpackChunkName: "auth" */ "@/views/auth/Login.vue");
const Register = () =>
  import(/* webpackChunkName: "auth" */ "@/views/auth/Register.vue");

export default [
  {
    path: "/login",
    name: "login",
    component: Login,
    meta: {
      layout: "auth-layout"
      //middleware: [guest]
    }
  },
  {
    path: "/register",
    name: "register",
    component: Register,
    meta: {
      layout: "auth-layout"
      // middleware: [guest]
    }
  }
];
