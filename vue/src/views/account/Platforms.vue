<template>
  <v-layout row wrap>
    <v-flex xs10 offset-xs1>
      <v-card>
        <v-toolbar color="indigo" dark>
          <v-toolbar-side-icon></v-toolbar-side-icon>

          <v-toolbar-title>Accounts</v-toolbar-title>

          <v-spacer></v-spacer>
        </v-toolbar>
        <v-list>
          <v-list-tile v-for="(account, index) in accounts" :key="index">
            <v-list-tile-action>
              <v-icon v-if="item.icon" color="pink">star</v-icon>
            </v-list-tile-action>

            <v-list-tile-content>
              <v-list-tile-title v-text="account.username"></v-list-tile-title>
            </v-list-tile-content>

            <v-list-tile-avatar>
              <img :src="item.avatar" />
            </v-list-tile-avatar>
          </v-list-tile>
        </v-list>
      </v-card>
    </v-flex>
    <v-flex xs10 offset-xs1>
      <ButtonPlatforms />
    </v-flex>
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
export default class extends Vue {
  accounts = [];
  mounted() {}
  getAccountPlatforms() {}
}
</script>
