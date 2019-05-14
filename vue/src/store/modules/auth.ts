import axios from "axios";
import * as types from "@/store/mutation-types";
import router from "@/router/";
import { AuthData, AuthRefresh } from "@/types";
import { Dispatch } from "vuex";
import store from "@/store";
//USED FOR DEBUGGIN
const LOCATION = "store/modules/auth.ts";

//Type Auth State
interface AuthState {
  [index: string]: boolean;
  is_auth: boolean;
  is_mail_verified: boolean;
}

interface dispatch {
  dispatch: Dispatch;
}

//Model Auth State
const authModel: AuthState = {
  // auth status
  is_auth: false,

  // email validation status
  is_mail_verified: false
};

// state
export const state = {
  ...authModel
};

// getters
export const getters = {
  is_auth: (state: AuthState) => state.is_auth,

  is_mail_verified: (state: AuthState) => state.is_mail_verified
};

// mutations
export const mutations = {
  INIT(state: AuthState, data: AuthData) {
    try {
      state.is_auth = true;

      state.is_mail_verified = Boolean(
        data.joined_at && data.joined_at < new Date().getTime() / 1000
      );
    } catch (e) {
      console.log({ Location: LOCATION, Error: e });
    }
  },

  // cleanup/reset module state. used at logout
  RESET(state: AuthState) {
    try {
      Object.keys(state).forEach(key => {
        state[key] = authModel[key];
      });

      router.push({ name: "login" });
    } catch (e) {
      console.log({ Location: LOCATION, Error: e });
    }
  }
};

// actions
export const actions = {
  // login
  [types.LOGIN]({ dispatch }: dispatch, user: AuthRefresh) {
    dispatch("INIT_STORE_DATA", user, { root: true });
  },

  // logout
  async [types.LOGOUT]() {
    try {
      let { status } = await axios.delete("/api/auth/logout");
      if (status === 200) {
        store.dispatch("RESET_STORE_DATA", {}, { root: true });
      }
    } catch (e) {
      console.log({ location: LOCATION, error: e });
    }
  },

  // used by page-visibility.js in /src/plugins/
  async [types.AUTH_REFRESH]({ dispatch }: dispatch) {
    try {
      const { status, data } = await axios.get("/api/auth/refresh");
      console.log(status, data);
      if (status === 200 && data.data) {
        store.dispatch("INIT_STORE_DATA", data.data, { root: true });
      }
    } catch (e) {
      console.log({ location: LOCATION, error: e });
    }
  }
};
