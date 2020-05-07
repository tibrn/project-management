import { Pagination } from './../types/models/pagination.d'
import { ref, Ref } from '@vue/composition-api'

export function createPagination (per_page = 20) {
  const pagination: Ref<null | Pagination> = ref(null)

  const nextPage = () => {
    return {
      per_page,
      page: pagination.value ? pagination.value.page + 1 : 1
    }
  }
  return {
    pagination, nextPage
  }
}
