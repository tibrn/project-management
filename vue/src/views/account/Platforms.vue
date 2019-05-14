<template>
  <v-layout align-center justify-center>
    <ButtonPlatforms />
  </v-layout>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
const ButtonPlatforms = () =>
  import(/* webpackChunkName: "account" */ "@/components/ButtonPlatforms.vue");
@Component({
  name: "Platforms",
  components: {
    ButtonPlatforms
  }
})
export default class extends Vue {}
</script>
