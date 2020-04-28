import { instance as store } from './../store/index'
import { route } from 'quasar/wrappers'
import VueRouter from 'vue-router'
import { StoreInterface } from '../store'
import routes from './routes'
import MiddlewarePlugin from 'vue-router-middleware-plugin'

/*
 * If not building with SSR mode, you can
 * directly export the Router instantiation
 */

export let instance: VueRouter
export default route<StoreInterface>(function ({ Vue }) {
  Vue.use(VueRouter)

  const router = new VueRouter({
    scrollBehavior: (): {x: number; y: number} => ({ x: 0, y: 0 }),
    routes,
    // Leave these as they are and change in quasar.conf.js instead!
    // quasar.conf.js -> build -> vueRouterMode
    // quasar.conf.js -> build -> publicPath
    mode: process.env.VUE_ROUTER_MODE,
    base: process.env.VUE_ROUTER_BASE,
  })

  Vue.use(MiddlewarePlugin as any, {
    router,
    context: { store },
  })

  instance = router

  return router
})
