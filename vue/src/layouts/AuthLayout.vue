<template>
  <div class="Auth">
    <Transition name="fade" mode="out-in">
      <RouterView :key="$route.name" />
    </Transition>

    <div class="Auth__bg" />
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
