<template>
  <div></div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
@Component({
  name: "TasksIndex"
})
export default class extends Vue {
  isLoading = false;
  tasks = [];
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
      let { data } = await this.axios.get("/api/tasks");

      console.log(data);
    } catch (e) {
      console.log("TaskIndex", e);
    } finally {
      this.isLoading = false;
    }
  }
}
</script>
