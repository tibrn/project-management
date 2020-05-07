<template>
  <q-page padding>
    <h4
      v-if="project"
      class="Title"
    > {{project.name}}</h4>

    <div class="q-pa-md">

      <div class="row">
      </div>
    </div>
  </q-page>
</template>

<script lang="ts">
import {
  defineComponent,
  computed,
  reactive,
  toRefs,
  onMounted,
} from '@vue/composition-api'
export default defineComponent({
  name: 'ProjectsShow',
  setup(props, ctx) {
    const { $utils, axios, $route } = ctx.root

    const state = reactive({
      project: null,
      tasks: [],
      isLoading: false,
    })

    const ID = computed(() => $route.params.id)

    const loadProject = () => {
      $utils.request({
        vm: state,
        call: async () => {
          const [projectRequest, tasksRequest] = await Promise.all([
            axios.get(`/api/projects/${ID.value}`),
            axios.get(`/api/tasks`, {
              params: {
                project_id: ID.value,
              },
            }),
          ])

          state.project = projectRequest.data.data
          state.tasks = tasksRequest.data.data
        },
      })
    }

    onMounted(loadProject)

    return {
      ID,
      ...toRefs(state),
    }
  },
})
</script>
