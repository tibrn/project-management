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
  default(state: StateLayout) {
    return state.default;
  },
  // get current layout name
  layout(state: StateLayout) {
    return state.status || state.default;
  }
};

export const mutations = {
  [types.SET_LAYOUT](state: StateLayout, payload: string) {
    // set/change app layout
    state.status = payload;
  }
};
