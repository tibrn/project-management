<template>
  <q-card
    class="my-card"
    flat
    bordered
  >

    <q-card-section>

      <div class="text-h5 q-mt-sm q-mb-xs">{{project.name}}</div>
      <div class="text-caption text-grey">
        {{project.description}}
      </div>
    </q-card-section>

    <q-card-actions>
      <q-btn
        flat
        color="dark"
        label="Edit"
      />
      <q-btn
        flat
        color="primary"
        label="Add Task"
      />

      <q-space />

      <!-- <q-btn
        color="grey"
        round
        flat
        dense
        :icon="expanded ? 'keyboard_arrow_up' : 'keyboard_arrow_down'"
        @click="expanded = !expanded"
      /> -->
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
  // setup(props, ctx) {

  // },
})
</script>
