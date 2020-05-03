
import axios from 'axios'
import VueAxios from 'vue-axios'
import { Dialog, Notify } from 'quasar'
import { boot } from 'quasar/wrappers'
import { instance as store } from 'src/store/index'
import { instance as router } from 'src/router/index'
import { TokenStorage } from 'src/services/token'
const AxiosInstance = axios.create({
  baseURL: process.env.API
})

if (TokenStorage.getToken()) {
  axios.defaults.headers.common.Authorization = TokenStorage.getAuthentication().Authorization
}

let countErrors = 0
const transform = (type: string): string => {
  switch (type) {
    case 'error':
      return 'negative'
    case 'success':
      return 'positive'
    case 'warning':
      return 'warning'
    default:
      return 'info'
  }
}
// Response interceptor
AxiosInstance.interceptors.response.use(
  (response) => {
    // console.log('--> axios response interceptor:', response.config.url)
    const { data } = response

    if (data.message) {
      Notify.create({
        message: data.message,
        timeout: 2500,
        progress: true,
        position: 'top',
        color: transform(data.type),
        actions: [{ icon: 'close', color: 'white' }]
      })
    }

    return response
  },
  (error) => {
    if (!error.response) {
      return Promise.reject(error)
    }
    const { status, data } = error.response

    let isMessageShown = false

    if (typeof error.response.data.message !== 'undefined' && [200, 422, 400, 403, 401].includes(status)) {
      if (error.response.data.message && !isMessageShown) {
        store.commit('app/SET_ERROR', true)

        countErrors += 1

        const remove = (): void => {
          countErrors -= 1

          if (countErrors === 0) {
            store.commit('app/SET_ERROR', false)
          }
        }

        isMessageShown = true
        Notify.create({
          message: data.message,
          timeout: 2500,
          progress: true,
          position: 'top',
          color: transform(data.type),
          actions: [{ icon: 'close', color: 'white' }],
          onDismiss: remove
        })
      }
    }

    if (error.response.data && status === 401) {
      if (error.response.data.error === "Token is expired") {
        TokenStorage.clear()
        router.push({ name: 'login' })
      }
    }

    // handle server errors
    if (status >= 500 && !isMessageShown) {
      if (data.message) {
        isMessageShown = true

        Dialog.create({
          title: 'Error',
          message: data.message,
          ok: true
        })
      }
    }

    // handle 404 errors
    if (status === 404) {
      router.push({ name: 'NotFound' })
    }

    return Promise.reject(error)
  }
)

export const instance = AxiosInstance
export default boot(({ Vue }) => {
  Vue.use(VueAxios, AxiosInstance)
})
