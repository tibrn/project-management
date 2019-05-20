<template>
  <v-app>
    <v-toolbar></v-toolbar>
    <Sidebar />
    <v-content>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-content>
    <v-footer></v-footer>
  </v-app>
</template>

<script>
import { Vue, Component } from "vue-property-decorator";
const Sidebar = () =>
  import(/* webpackChunkName: "sidebar" */ "@/components/Sidebar.vue");
@Component({
  name: "AppLayout",
  components: {
    Sidebar
  }
})
export default class extends Vue {
  created() {
    console.log("App-Layout");
  }
}
</script>
