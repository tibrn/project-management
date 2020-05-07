import { RootState } from './../types/store/index.d'
import Vue from 'vue'
import { store } from 'quasar/wrappers'
import Vuex, { mapState, mapGetters, mapMutations, mapActions, NamespacedMappers } from 'vuex'
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
Vue.use(Vuex)
export const instance = new Vuex.Store<RootState>({
  modules,

  // enable strict mode (adds overhead!)
  // for dev mode only
  strict: !!process.env.DEV
})

export default store(() => {
  return instance
})
const binder = { $store: instance }

const bindStore = (fun: Function): any => {
  return function (args: Array<string>) {
    const object = fun(args)
    Object.keys(object).map(key => {
      object[key] = object[key].bind(binder)
    })

    return object
  }
}

// export const { createNamespacedHelpers } = wrapStore(instance)
export const createNamespacedHelpers = (namespace: string): NamespacedMappers => {
  return {
    mapState: bindStore(mapState.bind(null, namespace)),
    mapGetters: bindStore(mapGetters.bind(null, namespace)),
    mapMutations: bindStore(mapMutations.bind(null, namespace)),
    mapActions: bindStore(mapActions.bind(null, namespace))
  }
}
