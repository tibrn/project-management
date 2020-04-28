<template>
  <div id="q-app">
    <router-view />
  </div>
</template>

<script lang="ts">

import { createComponent, computed } from '@vue/composition-api'
export default createComponent({
  name: 'App',
  setup(props, ctx) {
    const layout = computed(() => (ctx.root.$route.meta.layout || 'simple') + '-layout')

    return {
      layout,
    }
  }

})
</script>
