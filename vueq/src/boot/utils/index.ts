import { boot } from 'quasar/wrappers'
import { ModuleStatic } from 'src/types/boot/static'

const requireContext = (require as any).context('src/boot/utils/parts', false, /.*\.ts$/)

export const utils = requireContext.keys()
  .map((file: string) => [file.replace(/(^.\/)|(\.ts$)/g, ''), requireContext(file)])
  .reduce((modules: Array<ModuleStatic>, [name, module]: [string, ModuleStatic]) => {
    if (typeof module.default === 'undefined' || name === 'index') {
      return { ...modules }
    }
    return { ...modules, [name]: module.default }
  }, {})

Object.freeze(utils)

export const callback = boot(({ Vue }) => {
  Vue.prototype.$utils = utils
})
