<template>
  <q-layout view="hHh lpR fFf">

    <q-header elevated class="bg-primary text-white">
      <q-toolbar>
        <q-btn dense flat round icon="menu" @click="left = !left" />

        <q-toolbar-title>
          <q-avatar>
            <!-- <img src="https://cdn.quasar.dev/logo/svg/quasar-logo.svg"> -->
          </q-avatar>
          Project Management
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="left" side="left" overlay elevated>
      <!-- drawer content -->
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>

  </q-layout>
</template>

<script lang="ts">
import { createComponent, reactive, ref } from '@vue/composition-api'
// import { createNamespacedHelpers } from 'vuex'
// const { mapState } = createNamespacedHelpers('app')

export default createComponent({
  name: 'AppLayout',

  setup({ ctx }) {
    return {
      left:ref(true),
    }
  }
})
</script>
