<template>
  <q-page padding>
    <h4> {{isNew ? 'New' : 'Edit'}} Project</h4>
  </q-page>
</template>

<script lang="ts">
import {
  defineComponent,
  computed,
  reactive,
  toRefs,
  onMounted,
  ref,
  Ref,
} from '@vue/composition-api'
import { QForm } from 'quasar'

export default defineComponent({
  name: 'ProjectsEdit',
  setup(props, ctx) {
    const { $utils, axios, $route } = ctx.root
    const state = reactive({
      isLoading: false,
      isEditing: false,
      form: {},
    })
    const isNew = computed(() => $route.name === 'project-new')

    const ID = computed(() => $route.params.id)

    const form: Ref<null | QForm> = ref(null)

    const save = () => {
      $utils.request({
        vm: state,
        call: async () => {
          if (!form.value) return
          if (!(await form.value.validate())) return
          await axios.post(`/api/projects`, state.form)
        },
        loading: 'isEditing',
      })
    }

    const loadProject = () => {
      $utils.request({
        vm: state,
        call: async () => {
          const { data } = await axios.get(`/api/projects/${ID.value}`)

          state.form = data.data
        },
      })
    }

    onMounted(() => {
      if (!isNew.value) {
        loadProject()
      }
    })

    return {
      isNew,
      ...toRefs(state),
      save,
    }
  },
})
</script>
