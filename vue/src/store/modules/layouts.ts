import * as types from "@/store/mutation-types.ts";

export const state = {
  default: "app-layout",
  status: null
};

interface StateLayout {
  default: string;
  status: null | string;
}
export const getters = {
  // get current layout name
  layout(state: StateLayout) {
    return state.status || state.default;
  }
};

export const mutations = {
  [types.SET_LAYOUT](state: StateLayout, payload: string) {
    // set/change app layout
    state.status = payload;
    console.log("LAYOUT", payload);
    // add layout name as class to <body> tag
    const body = document.querySelector("body");
    if (body) {
      body.classList.remove("app", "auth", "simple");
      body.classList.add(payload.split("-")[0]);
    }
  }
};
