import Vue from "vue";
import Vuetify from "vuetify/lib";
import "vuetify/src/stylus/app.styl";
import en from "@/lang/en.json";
import ro from "@/lang/ro.json";
Vue.use(Vuetify, {
  theme: {
    primary: "#2098D1",
    secondary: "#424242",
    accent: "#82B1FF",
    error: "#FF5252",
    info: "#2196F3",
    success: "#4CAF50",
    warning: "#FFC107"
  },
  options: {
    customProperties: true
  },
  iconfont: "mdi",
  lang: {
    locales: { en, ro },
    current: "en"
  }
});
