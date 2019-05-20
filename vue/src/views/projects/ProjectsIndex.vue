<template>
  <div></div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { Getter } from "vuex-class";
@Component({
  name: "ProjectsIndex"
})
export default class extends Vue {
  @Getter("/user/id") UserID!: string;
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
      let { data } = await this.axios.get("/api/projects");

      console.log(data);
    } catch (e) {
      console.log("ProjectsIndex", e);
    } finally {
      this.isLoading = false;
    }
  }
}
</script>
