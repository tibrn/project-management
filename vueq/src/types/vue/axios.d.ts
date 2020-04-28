import { AxiosInstance } from 'axios'

declare module 'vue/types/vue' {
  // Global properties can be declared
  // on the `VueConstructor` interface
  interface Vue {
    $axios: Readonly<AxiosInstance>;
    $utils: Readonly<UtilsInterface>;
  }
}
