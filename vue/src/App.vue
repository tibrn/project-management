<template>
  <v-app>
    <Component :is="layout" v-if="layout" />
    <v-snackbar
      v-if="queue.message"
      v-model="queue.isMessage"
      :top="true"
      :color="queue.message['message-type']"
      :timeout="2000"
    >
      {{ queue.message.text }}
      <v-btn flat @click="queue.closeMessage()">
        Close
      </v-btn>
    </v-snackbar>
  </v-app>
</template>

<script lang="ts">
const AuthLayout = () =>
  import(/* webpackChunkName: "layouts" */ "@/layouts/AuthLayout.vue");
const AppLayout = () =>
  import(/* webpackChunkName: "layouts" */ "@/layouts/AppLayout.vue");
const SimpleLayout = () =>
  import(/* webpackChunkName: "layouts" */ "@/layouts/SimpleLayout.vue");
import { Getter } from "vuex-class";
import { Component, Vue } from "vue-property-decorator";
import Vuetify from "vuetify/lib";
import { QueueMessages } from "@/class/QueueMessages";
console.log(Vuetify);
//TODO: Tree Shaking dupa ce ajung intr-un punct in care sa imi dau seama de
//TODO: ce componente am nevoie
Vue.use(Vuetify);
@Component({
  name: "App",
  components: {
    "auth-layout": AuthLayout,
    "app-layout": AppLayout,
    "simple-layout": SimpleLayout
  }
})
export default class App extends Vue {
  @Getter("layouts/layout") layout!: string;
  @Getter("user/queue") queue!: QueueMessages;

  created() {
    this.queue.sendMessage({ "message-type": "success", text: "TEST" });
  }
}
</script>
<style lang="scss"></style>
