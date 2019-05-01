<template>
  <v-app>
    <v-navigation-drawer app></v-navigation-drawer>
    <v-toolbar app></v-toolbar>
    <v-content>
      <router-view></router-view>
    </v-content>
    <v-footer app></v-footer>
    <v-snackbar
      v-if="queue.message"
      v-model="queue.isMessage"
      :top="true"
      :color="queue.message['message-type']"
      :timeout="1000"
    >
      {{ queue.message.text }}
      <v-btn flat @click="queue.closeMessage">
        Close
      </v-btn>
    </v-snackbar>
  </v-app>
</template>

<script lang="ts">
// const AuthLayout = () =>
//   import(/* webpackChunkName: "auth" */ "@/layouts/AuthLayout.vue");
// const AppLayout = () => import("@/layouts/AppLayout.vue");
// const SimpleLayout = () => import("@/layouts/SimpleLayout.vue");
import { Getter } from "vuex-class";
import { Component, Vue } from "vue-property-decorator";
import Vuetify from "vuetify/lib";
import { QueueMessages } from "@/class/QueueMessages";
console.log(Vuetify);
//TODO: Tree Shaking dupa ce ajung intr-un punct in care sa imi dau seama de
//TODO: ce componente am nevoie
Vue.use(Vuetify);
@Component({
  name: "App"
  // components: {
  //   "auth-layout": AuthLayout,
  //   "app-layout": AppLayout,
  //   "simple-layout": SimpleLayout
  // }
})
export default class App extends Vue {
  @Getter("layouts/layout") layout!: string;
  @Getter("user/queue") queue!: QueueMessages;

  created() {
    console.log(this.$store);
    console.log(this.queue);
  }
}
</script>
<style lang="scss"></style>
