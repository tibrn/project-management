import Cookies from "js-cookie";
import * as types from "@/store/mutation-types";
import store from "@/store";
interface StateLocale {
  locale: string;
  locales: object;
}
const locale = "en";
const locales = {
  en: "en",
  ro: "ro"
};

// state
export const state = {
  locale: Cookies.get("locale") || locale,
  locales: locales
};

// getters
export const getters = {
  locale: (state: StateLocale) => state.locale,
  locales: (state: StateLocale) => state.locales
};

// mutations
export const mutations = {
  [types.SET_LOCALE]: (state: StateLocale, locale: string) => {
    state.locale = locale;
  }
};

// actions
export const actions = {
  setLocale(state: StateLocale, locale: string) {
    store.commit(types.SET_LOCALE, locale);

    Cookies.set("locale", locale, { expires: 365 });
  }
};
