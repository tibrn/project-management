import { boot } from 'quasar/wrappers'
import { ModuleStatic } from 'src/types/boot/static'

import Vue from 'vue'

const requireContext = (<any>require).context('src/boot/utils/parts', false, /.*\.ts$/)

const utils = requireContext.keys()
  .map((file: string) => [file.replace(/(^.\/)|(\.ts$)/g, ''), requireContext(file)])
  .reduce((modules: Array<ModuleStatic>, [name, module]: [string, ModuleStatic]) => {
    if (typeof module.default === 'undefined' || name === 'index') {
      return { ...modules }
    }
    return { ...modules, [name]: module.default }
  }, {})

Object.freeze(utils)

export default boot(({ Vue }) => {
  Vue.prototype.$utils = utils
})
