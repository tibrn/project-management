<template>
  <div class="Auth">
    <v-content>
      <router-view></router-view>
    </v-content>
  </div>
</template>

<script>
import { Vue, Component } from "vue-property-decorator";
@Component({
  name: "AuthLayout"
})
export default class extends Vue {}
</script>

<style lang="scss"></style>
