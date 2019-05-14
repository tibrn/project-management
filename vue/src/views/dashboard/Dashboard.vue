<template>
  <div></div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
@Component({
  name: "Dashboard"
})
export default class extends Vue {
  created() {
    console.log("Dashboard");
  }
}
</script>
