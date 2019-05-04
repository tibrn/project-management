<template>
  <div>
    <v-toolbar app></v-toolbar>
    <Sidebar />
    <v-content>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-content>
    <v-footer app></v-footer>
  </div>
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
