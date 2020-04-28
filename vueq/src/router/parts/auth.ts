import { RouteConfig } from 'vue-router'

const Login = () => import(/* webpackChunkName: "auth" */ 'src/pages/auth/Login.vue')
// const Register = () => import(/* webpackChunkName: "auth" */ 'src/pages/auth/Register.vue');

import NoAuthMiddleware from '../middlewares/no_auth'

export default ([
  {
    name: 'login',
    path: '/login',
    component: Login,
    meta: {
      layout: 'simple',
      middleware: {
        attach: [NoAuthMiddleware]
      }
    },
  },
  // {
  //   name: 'register',
  //   path: '/register',
  //   component: Register,
  //   meta: {
  //     layout: 'simple',
  //     middleware: {
  //       attach: [NoAuthMiddleware]
  //     }
  //   },
  // },
]) as Array<RouteConfig>
