import store from "@/store/";
import { Route } from "vue-router";
export default (to: Route, from: Route) => {
  console.log("AFTER", to, from);
};
