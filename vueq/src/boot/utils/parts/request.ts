import { RequestUtil } from "src/types/boot/utils"

export default async function ({
  vm, loading, errors, call, debug
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

  if (typeof vm[loading] === 'boolean') {
    vm[loading] = true
  }

  try {
    await call()
  } catch (e) {
    if (debug) {
      console.error(e)
    }
    if (e.response && e.response.data && e.response.data.errors && typeof vm[errors] === 'object') {
      Object.assign(
        vm[errors],
        ...Object.keys(e.response.data.errors).map(key => ({
          [key]: e.response.data.errors[key][0]
        }))
      )
    }
  }
  if (typeof vm[loading] === 'boolean') {
    vm[loading] = false
  }
}
