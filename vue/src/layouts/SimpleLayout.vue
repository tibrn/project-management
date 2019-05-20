<template>
  <v-app>
    <v-content>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import { Vue, Component } from "vue-property-decorator";
@Component({
  name: "SimpleLayout"
})
export default class extends Vue {
  created() {
    console.log("Simple-Layout");
  }
}
</script>

<style lang="scss"></style>
