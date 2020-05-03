<template>
  <q-scroll-area class="fit">
    <q-list padding>
      <q-item
        v-for="link in links"
        :key="link.text"
        :to="{name : link.to}"
        v-ripple
        clickable
      >
        <q-item-section avatar>
          <q-icon
            color="grey"
            :name="link.icon"
          />
        </q-item-section>
        <q-item-section>
          <q-item-label>{{ link.text }}</q-item-label>
        </q-item-section>
      </q-item>

    </q-list>
  </q-scroll-area>
</template>

<script>

import { defineComponent } from '@vue/composition-api'

export default defineComponent({
  name: 'Sidebar',

  setup () {
    const links = [
      { icon: 'home', text: 'Home', to: 'home' },
      { icon: 'whatshot', text: 'Projects', to: 'projects' },

    ]

    return {
      links,
    }
  }
})
</script>
