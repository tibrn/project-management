<template>
  <div>
    <v-container fluid>
      <router-view />
    </v-container>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
@Component({
  name: "Main"
})
export default class extends Vue {
  created() {
    console.log("Dashboard");
  }
}
</script>
