<template>
  <q-layout view="hHh lpR fFf">
    <q-header
      elevated
      class="bg-primary text-white"
    >
      <q-toolbar>
        <q-btn
          dense
          flat
          round
          icon="menu"
          @click="isSidebar = !isSidebar"
        />

        <q-toolbar-title>
          <q-avatar>
            <!-- <img src="https://cdn.quasar.dev/logo/svg/quasar-logo.svg"> -->
          </q-avatar>Project Management
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="isSidebar"
      side="left"
      show-if-above
      bordered
      content-class="bg-grey-2"
      :width="240"
    >
      <Sidebar />
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import { defineComponent, ref } from '@vue/composition-api'
import Sidebar from 'src/components/Sidebar.vue'

export default defineComponent({
  name: 'AppLayout',
  components: {
    Sidebar,
  },
  setup(props, ctx) {
    return {
      isSidebar: ref(true),
    }
  },
})
</script>
