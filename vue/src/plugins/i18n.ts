import Vue from "vue";
import store from "@/store/";
import VueI18n from "vue-i18n";

// add plugin to Vue
Vue.use(VueI18n);

// Create VueI18n instance with options
const i18n = new VueI18n({
  locale: "en", // set locale
  messages: {} // set locale messages
});

// async change locale
export async function loadMessages(locale: string) {
  if (Object.keys(i18n.getLocaleMessage(locale)).length === 0) {
    const messages = await import(/* webpackChunkName: "lang/[request]" */ `@/lang/${locale}`);
    i18n.setLocaleMessage(locale, messages);
  }

  if (i18n.locale !== locale) {
    i18n.locale = locale;
  }
}

// local initial setup
(async function() {
  await loadMessages(store.getters["lang/locale"]);
})();

export default i18n;
