

export default (arg: any, target :object) :object => {
  Object.keys(arg).map(key => {
    Object.defineProperty(target, key, {
      get() {
        return arg[key]
      },
      set(val: any) {
        arg[key] = val
      }
    })
  })
  return target
}
