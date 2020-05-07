<template>
  <q-card
    class="Card"
    bordered
  >

    <q-card-section>

      <div class="text-h5 q-mt-sm q-mb-xs">{{project.name}}</div>
      <div class="Card__description text-caption text-grey q-pa-sm">
        {{ project.description}}
      </div>
    </q-card-section>

    <q-card-actions class="Card__actions">
      <q-btn
        flat
        color="dark"
        label="Edit"
        @click="$router.push({name:'project-edit',params:{ id: project.id}})"
      />
      <q-btn
        flat
        color="primary"
        label="Add Task"
        @click="$router.push({name:'project-show',params:{ id: project.id}})"
      />

      <q-btn
        class="float-right"
        flat
        color="negative"
        label="Delete"
        @click="$emit('delete',project)"
      />

      <q-space />

    </q-card-actions>

    <!-- <q-slide-transition>
      <div v-show="expanded">
        <q-separator />
        <q-card-section class="text-subitle2">
          {{ lorem }}
        </q-card-section>
      </div>
    </q-slide-transition> -->
  </q-card>
</template>

<script lang="ts">
import { defineComponent, PropType } from '@vue/composition-api'
import { Project } from 'src/types/models/project'
export default defineComponent({
  name: 'ProjectCard',
  props: {
    project: {
      type: Object as PropType<Project>,
      required: true,
    },
  },
})
</script>
<style lang="scss" scoped>
.Card {
  width: 100%;
  max-width: 350px;
  height: 200px;
  word-wrap: break-word;
  position: relative;
  &__description {
    overflow: hidden;
    text-overflow: ellipsis;
    max-height: 70px;
  }

  &__actions {
    position: absolute;
    bottom: 0;
    left: 0;
    width: 100%;
    display: block;
  }
}
</style>
