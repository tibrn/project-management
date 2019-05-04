import Vue from "vue";
import axios from "axios";
import VueAxios from "vue-axios";
import store from "@/store/";
import router from "@/router";
import { QueueMessages, Message } from "@/class/QueueMessages";
Vue.use(VueAxios, axios);
// Response interceptor

axios.interceptors.response.use(
  function(response) {
    // console.log('--> axios response interceptor:', response.config.url)
    const { status, data } = response;
    // show response message/Golang

    if (typeof data.message !== "undefined" && data.message) {
      const message = {
        text: data.message,
        "message-type": data["message-type"]
      };
      QueueMessages.getInstance().sendMessage(message);
    }

    return response;
  },
  function(error) {
    const { status, data } = error.response;
    // handle 404 errors
    if (status === 404) {
      router.push({ name: "NotFound" });
    }

    if (typeof data.message !== "undefined" && data.message) {
      setTimeout(function() {
        QueueMessages.getInstance().sendMessage({
          text: data.message,
          "message-type": data["message-type"]
        });
      }, 0);
    }
    return Promise.reject(error);
  }
);
