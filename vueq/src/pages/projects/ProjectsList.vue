<template>
  <q-page padding>

    <div class="q-pa-md">

      <div class="row ">
        <div class="col-3">
          <h4 class="Title"> Projects</h4>
        </div>
        <div class="col-8 text-right">
          <q-btn @click="$router.push({name:'project-new'})">
            New Project
          </q-btn>
        </div>
      </div>

      <div
        v-if="projects"
        class="row q-py-lg"
      >
        <div
          v-for="project in projects"
          class="col-3 q-px-sm"
          :key="project.id"
        >

          <ProjectCard :project="project" />

        </div>

      </div>

    </div>
    <q-inner-loading :showing="isLoading">
      <q-spinner-gears
        size="50px"
        color="primary"
      />
    </q-inner-loading>
  </q-page>
</template>

<script>
import { defineComponent, reactive, toRefs, onMounted } from '@vue/composition-api'

const ProjectCard = () => import(/* webpackChunkName: "projects" */ 'src/components/cards/ProjectCard')
export default defineComponent({
  name: 'ProjectsList',
  components: {
    ProjectCard,
  },
  setup (props, ctx) {
    const state = reactive({
      projects: null,
      isLoading: false,
    })

    const { $utils, axios } = ctx.root

    onMounted(() => {
      $utils.request({
        vm: state,
        call: async () => {
          const { data } = await axios.get("/api/projects")

          state.projects = data.data
        }
      })
    })

    return {
      ...toRefs(state),

    }
  },
})
</script>
<style lang="scss">
.Title {
  margin: 0px;
}
</style>
