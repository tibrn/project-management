import VueRouterMultiguard from "vue-router-multiguard";
import auth from "@/middleware/modules/auth";
const ProjectsIndex = () =>
  import(/* webpackChunkName: "projects" */ "@/views/projects/ProjectsIndex.vue");
export default [
  {
    path: "projects",
    name: "projects",
    component: ProjectsIndex,
    meta: {
      layout: "app-layout"
    },
    beforeEnter: VueRouterMultiguard([auth])
  }
];
