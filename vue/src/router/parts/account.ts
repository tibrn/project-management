import VueRouterMultiguard from "vue-router-multiguard";
const Account = () =>
  import(/* webpackChunkName: "account" */ "@/views/account/Account.vue");
const AccountSettings = () =>
  import(/* webpackChunkName: "account" */ "@/views/account/Settings.vue");
const AccountPlatforms = () =>
  import(/* webpackChunkName: "account" */ "@/views/account/Platforms.vue");
import auth from "@/middleware/modules/auth";
export default [
  {
    path: "account",
    name: "account",
    component: Account,
    meta: {
      layout: "app-layout"
    },
    beforeEnter: VueRouterMultiguard([auth]),
    children: [
      {
        path: "settings",
        name: "account-settings",
        component: AccountSettings,
        meta: {
          layout: "app-layout"
        },
        beforeEnter: VueRouterMultiguard([auth])
      },
      {
        path: "platforms",
        name: "account-platforms",
        component: AccountPlatforms,
        meta: {
          layout: "app-layout"
        },
        beforeEnter: VueRouterMultiguard([auth])
      }
    ]
  }
];
