<template>
  <div class="Auth">
    <v-content>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-content>
  </div>
</template>

<script>
import { Vue, Component } from "vue-property-decorator";
@Component({
  name: "AuthLayout"
})
export default class extends Vue {
  created() {
    console.log("Auth-Layout");
  }
}
</script>

<style lang="scss"></style>
