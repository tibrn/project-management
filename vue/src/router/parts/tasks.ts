import VueRouterMultiguard from "vue-router-multiguard";
import auth from "@/middleware/modules/auth";
const TasksIndex = () =>
  import(/* webpackChunkName: "tasks" */ "@/views/tasks/TasksIndex.vue");
export default [
  {
    path: "tasks",
    name: "tasks",
    component: TasksIndex,
    meta: {
      layout: "app-layout"
    },
    beforeEnter: VueRouterMultiguard([auth])
  }
];
