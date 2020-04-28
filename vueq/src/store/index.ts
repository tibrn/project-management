import { RootState } from './../types/store/index.d'

import { store } from 'quasar/wrappers'
import Vuex, { Store } from 'vuex'
import modules from './modules'
// import example from './module-example'
// import exampleState from './module-example/state'

/*
 * If not building with SSR mode, you can
 * directly export the Store instantiation
 */

export interface StoreInterface {
  // Define your own store structure, using submodules if needed
  // example: typeof exampleState;
  example: unknown;
}

export let instance: Store<RootState>

export default store(function ({ Vue }) {
  Vue.use(Vuex)

  const Store = new Vuex.Store<RootState>({
    modules,

    // enable strict mode (adds overhead!)
    // for dev mode only
    strict: !!process.env.DEV
  })

  instance = Store

  return Store
})
