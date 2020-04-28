import { isUndefined } from 'util';
import { ModuleStatic } from 'src/types/boot/static';
import { RouteConfig } from 'vue-router';

const requireContext = (<any>require).context('src/router/parts', false, /.*\.ts$/);


const routes : RouteConfig[] = requireContext.keys()
  .map((file : string) => [file.replace(/(^.\/)|(\.ts$)/g, ''), requireContext(file)])
  .reduce((modules : Array<ModuleStatic>, [_, module] : [string, ModuleStatic]) => {
    if (isUndefined(module.default)) {
      return [...modules];
    }
    return [...modules, ...module.default];
  },
  []);


// Always leave this as last one
if (process.env.MODE !== 'ssr') {
  routes.push({
    name: 'NotFound',
    path: '*',
    component: () => import('pages/Error404.vue'),
    meta: {
      layout: 'simple',
    },
  });
}

export default routes;
