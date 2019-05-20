import * as types from "@/store/mutation-types";
import "@/types";
import { QueueMessages, Message } from "@/class/QueueMessages";
import { AuthRefresh } from "@/types";
//USED FOR DEBUGGIN
const LOCATION = "store/modules/user.ts";

// Type State User
interface StateUser {
  [index: string]: string | number | object;
  id: number;
  name: string;
  slug: string;
  email: string;
  avatar: string;
  type: number;
  theme: string;

  queueMessages: QueueMessages;
}

// Model State User
const userModel: StateUser = {
  id: -1,
  name: "",
  slug: "",
  email: "",
  avatar: "",
  type: -1,
  theme: localStorage.getItem("theme") || "dark",

  queueMessages: QueueMessages.getInstance(2100)
};

// state
export const state = {
  ...userModel
};

// getters
export const getters = {
  id: (state: StateUser) => state.id,
  name: (state: StateUser) => state.name,
  email: (state: StateUser) => state.email,
  avatar: (state: StateUser) => state.avatar,
  type: (state: StateUser) => state.type,
  theme: (state: StateUser) => state.theme,

  queue: (state: StateUser) => state.queueMessages,
  isDarkTheme: (state: StateUser) => state.theme === "dark",
  isLightTheme: (state: StateUser) => state.theme === "light"
  // notifications_count: state => (state.is_init) ? state.data.count.notifications : 0,
};

// mutations
export const mutations = {
  INIT(state: StateUser, data: AuthRefresh) {
    try {
      state.id = data.id;
      state.name = data.name;
      state.slug = data.slug;
      state.email = data.email;
      state.type = data.type;

      state.avatar =
        (data.settings && data.settings.avatar) || userModel.avatar;
      state.theme = (data.settings && data.settings.theme) || userModel.theme;

      if (state.theme && state.theme !== "") {
        localStorage.setItem("theme", state.theme);
      }
    } catch (e) {
      console.log({ Location: LOCATION, Error: e });
    }
  },

  RESET(state: StateUser) {
    try {
      Object.keys(state).forEach(key => {
        if (key !== "theme") {
          state[key] = userModel[key];
        }
      });
    } catch (e) {
      console.log({ Location: LOCATION, Error: e });
    }
  },

  [types.UPDATE_THEME](state: StateUser, theme: string) {
    try {
      state.theme = theme;
      localStorage.setItem("theme", state.theme);
    } catch (e) {
      console.log({ Location: LOCATION, Error: e });
    }
  }
};

// actions
export const actions = {};
