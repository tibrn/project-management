<template>
  <div>
    <v-navigation-drawer app></v-navigation-drawer>
    <v-toolbar app></v-toolbar>
    <v-content>
      <router-view></router-view>
    </v-content>
    <v-footer app></v-footer>
  </div>
</template>

<script>
import { Vue, Component } from "vue-property-decorator";

@Component({
  name: "AppLayout"
})
export default class extends Vue {}
</script>
