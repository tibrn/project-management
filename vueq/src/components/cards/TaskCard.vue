<template>
  <q-card
    class="Card"
    bordered
  >

    <q-card-section>

      <div class="text-h5 q-mt-sm q-mb-xs">{{task.name}}</div>
      <div class="Card__description text-caption text-grey q-pa-sm">
        {{ task.description}}
      </div>
    </q-card-section>

    <q-card-actions class="Card__actions">

    </q-card-actions>

  </q-card>
</template>

<script lang="ts">
import { defineComponent, PropType } from '@vue/composition-api'
import { Task } from '../../types/models/task'
export default defineComponent({
  name: 'TaskCard',
  props: {
    task: {
      type: Object as PropType<Task>,
      required: true,
    },
  },
})
</script>
