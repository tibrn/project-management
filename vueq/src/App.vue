<template>
  <div id="q-app">
    <component :is="layout"/>
  </div>
</template>

<script lang="ts">

import { defineComponent, computed } from '@vue/composition-api'
import SimpleLayout from 'src/layouts/SimpleLayout.vue'
import AppLayout from 'src/layouts/AppLayout.vue'
export default defineComponent({
  name: 'App',
  components: {
    'simple-layout': SimpleLayout,
    'app-layout': AppLayout,
  },
  setup(props, ctx) {
    const layout = computed(() => (ctx.root.$route.meta.layout || 'simple') + '-layout')

    return {
      layout,
    }
  }

})
</script>
