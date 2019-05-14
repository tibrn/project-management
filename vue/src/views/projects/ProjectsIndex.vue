<template>
  <div></div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
@Component({
  name: "ProjectsIndex"
})
export default class extends Vue {
  isLoading = false;
  projects = [];
  created() {
    console.log("CEVA");
  }
  mounted() {
    console.log("GET");
    this.getTasks();
  }
  async getTasks() {
    this.isLoading = true;
    try {
      let { data } = await this.axios.get("/api/project");

      console.log(data);
    } catch (e) {
      console.log("ProjectsIndex", e);
    } finally {
      this.isLoading = false;
    }
  }
}
</script>
