import { RequestUtil } from "src/types/boot/utils"

export default async function ({
  vm, loading, errors, call
}: RequestUtil) {
  if (typeof call === 'undefined') {
    return
  }

  if (!loading) {
    loading = 'isLoading'
  }
  if (!errors) {
    errors = 'formErrors'
  }

  if (vm[loading]) {
    return
  }

  vm[loading] = true

  try {
    await call.apply(vm)
  } catch (e) {
    if (e.response && e.response.data && e.response.data.errors) {
      Object.assign(
        vm[errors],
        ...Object.keys(e.response.data.errors).map(key => ({
          [key]: e.response.data.errors[key][0]
        }))
      )
    }
  }
  vm[loading] = false
}
