<template>
  <v-app :dark="isDarkTheme" :light="isLightTheme">
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
import { QueueMessages } from "@/class/QueueMessages";
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
  @Getter("user/isDarkTheme") isDarkTheme!: boolean;
  @Getter("user/isLightTheme") isLightTheme!: boolean;

  created() {
    // this.queue.sendMessage({ "message-type": "success", text: "TEST" });
  }
}
</script>
<style lang="scss">
@import "../node_modules/@mdi/font/css/materialdesignicons.css";
@import "../node_modules/roboto-fontface/css/roboto/roboto-fontface.css";
</style>
