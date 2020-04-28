import { RouteConfig } from 'vue-router'
import { VNode } from 'vue'
import AuthMiddleware from '../middlewares/auth'
const ProjectsList = () => import(/* webpackChunkName: "projects" */ 'src/pages/projects/ProjectsList.vue')
const ProjectsEdit = () => import(/* webpackChunkName: "projects" */ 'src/pages/projects/ProjectsEdit.vue')
const ProjectsNew = () => import(/* webpackChunkName: "projects" */ 'src/pages/projects/ProjectsNew.vue')
export default ([
  {
    name: 'dashboard',
    path: '/',
    component: { render (h): VNode { return h('router-view') } },
    meta: {
      layout: 'app',
      middleware: {
        attach: [AuthMiddleware]
      }
    },
    children: [
      {
        name: 'dashboard',
        path: '/projects',
        component: ProjectsList,
        meta: {
          layout: 'app',
          middleware: {
            attach: [AuthMiddleware]
          }
        },
      },
      {
        name: 'dashboard',
        path: '/project/:id',
        component: ProjectsEdit,
        meta: {
          layout: 'app',
          middleware: {
            attach: [AuthMiddleware]
          }
        },
      },
      {
        name: 'dashboard',
        path: '/project/new',
        component: ProjectsNew,
        meta: {
          layout: 'app',
          middleware: {
            attach: [AuthMiddleware]
          }
        },
      }
    ]
  },

]) as Array<RouteConfig>
