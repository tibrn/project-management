<template>
  <div id="q-app">
    <component :is="layout" />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from '@vue/composition-api'
import SimpleLayout from 'src/layouts/SimpleLayout.vue'
import AppLayout from 'src/layouts/AppLayout.vue'
import { KeepAuth } from './services/keepAuth'
export default defineComponent({
  name: 'App',
  components: {
    'simple-layout': SimpleLayout,
    'app-layout': AppLayout,
  },
  setup(props, ctx) {
    KeepAuth.fire()
    const layout = computed(
      () => (ctx.root.$route.meta.layout || 'simple') + '-layout',
    )

    return {
      layout,
    }
  },
})
</script>
