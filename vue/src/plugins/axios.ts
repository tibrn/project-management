import Vue from "vue";
import axios from "axios";
import VueAxios from "vue-axios";
// import store from "@/store/";
import router from "@/router";
Vue.use(VueAxios, axios);
import { VMessages } from "vuetify/lib";
// Response interceptor
axios.interceptors.response.use(
  function(response) {
    // console.log('--> axios response interceptor:', response.config.url)
    const { status, data } = response;
    // show response message/Laravel
    // if (typeof data.message !== "undefined") {
    //   VMessages({
    //     dangerouslyUseHTMLString: true,
    //     message: data.message,
    //     type: "success",
    //     duration: 3000,
    //     showClose: true
    //   });
    // }
    return response;
  },
  function(error) {
    const { status, statusText } = error.response;
    // handle 404 errors
    if (status === 404) {
      router.push({ name: "NotFound" });
    }
    return Promise.reject(error);
  }
);
