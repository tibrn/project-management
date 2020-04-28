import { RouteContext } from 'vue-router-middleware-plugin/build/types/VueTypes'
import { Store } from 'vuex'
import { RootState } from '../store/index'

export interface Context extends RouteContext{
  store: Store<RootState>;
}

export type MiddlewareCallback = (context: Context) => void
