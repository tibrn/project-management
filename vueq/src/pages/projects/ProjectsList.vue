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

          <ProjectCard
            :project="project"
            @delete="confirmDelete"
          />

        </div>

      </div>

      <div
        v-if="pagination && pagination.total_pages > 1"
        class="row"
      >
        <q-pagination
          v-model="pagination.page"
          :max="10"
          :max-pages="pagination.total_pages"
          :boundary-numbers="true"
        >
        </q-pagination>
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

<script lang="ts">
import {
  defineComponent,
  reactive,
  toRefs,
  onMounted,
} from '@vue/composition-api'

const ProjectCard = () =>
  import(
    /* webpackChunkName: "projects" */ 'src/components/cards/ProjectCard.vue'
  )

import { createPagination, createDelete } from 'src/utils'

import { Project } from '../../types/models/project'
export default defineComponent({
  name: 'ProjectsList',
  components: {
    ProjectCard,
  },
  setup(props, ctx) {
    const state = reactive({
      projects: [],
      isLoading: false,
    })

    const { $utils, axios } = ctx.root

    const { pagination, nextPage } = createPagination(9)

    onMounted(() => {
      $utils.request({
        vm: state,
        call: async () => {
          const { data } = await axios.get('/api/projects', {
            params: { ...nextPage() },
          })

          state.projects = data.data
          pagination.value = data.pagination
        },
      })
    })

    const { confirmDelete } = createDelete({
      deleteRoute: (project: Project) => `/api/projects/${project.id}`,
      message: (project: Project) =>
        `Are you sure that you want to delete the project named ${project.name} ?`,
      title: (project: Project) => `Delete ${project.name}`,
      afterDelete: (project: Project) => {
        state.projects = state.projects.filter((prj: Project) => {
          return prj.id !== project.id
        })
      },
    })

    return {
      ...toRefs(state),
      pagination,
      confirmDelete,
    }
  },
})
</script>
<style lang="scss">
</style>
