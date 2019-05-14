<template>
  <div>
    <v-btn
      :loading="isLoading"
      :disabled="isLoading"
      color="info"
      @click="launchLogin('github')"
    >
      Github
      <template v-slot:loader>
        <span class="custom-loader">
          <v-icon light>cached</v-icon>
        </span>
      </template>
    </v-btn>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
@Component({
  name: "ButtonPlatforms"
})
export default class extends Vue {
  isLoading = false;
  popup: null | Window = null;
  popupCheck = -1;
  popupWidth = 600;
  popupHeight = 560;
  created() {
    // listen messages from social popup
    window.addEventListener("message", this.receiveMessage, false);
  }
  destroyed() {
    // remove listener for messages from social popup
    window.removeEventListener("message", this.receiveMessage, false);
  }
  launchLogin(platform: string) {
    this.isLoading = true;

    var left = screen.width / 2 - this.popupWidth / 2;
    var top = screen.height / 2 - this.popupHeight / 2;
    var url =
      (process.env.NODE_ENV === "production" ? "" : process.env.VUE_APP_URL) +
      "/" +
      platform;
    var config =
      "toolbar=no, location=no, status=no, menubar=no, scrollbars=no";
    this.popup = window.open(
      url,
      "Facebook",
      config +
        ", width=" +
        this.popupWidth +
        ", height=" +
        this.popupHeight +
        ", top=" +
        top +
        ", left=" +
        left
    );

    this.popupCheck = setInterval(() => {
      if (this.popup && this.popup.closed) {
        clearInterval(this.popupCheck);
        this.isLoading = false;
      }
    }, 400);
  }

  receiveMessage(event: Event) {
    console.log(event);
  }
}
</script>
