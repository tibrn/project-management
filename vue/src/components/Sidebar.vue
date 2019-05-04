<template>
  <v-navigation-drawer v-model="drawer" permanent absolute>
    <v-toolbar v-if="avatar" flat class="transparent">
      <v-list class="pa-0">
        <v-list-tile avatar>
          <v-list-tile-avatar>
            <img :src="avatar" />
          </v-list-tile-avatar>

          <v-list-tile-content>
            <v-list-tile-title>name</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-toolbar>

    <v-list class="pt-0" dense>
      <v-list-group prepend-icon="account_circle" value="true">
        <template v-slot:activator>
          <v-list-tile :to="{ name: 'account' }" exact>
            <v-list-tile-title>Account</v-list-tile-title>
          </v-list-tile>
        </template>
        <v-list-tile
          v-for="(action, i) in account"
          :key="i"
          :to="{ name: action.route }"
          exact
        >
          <v-list-tile-title v-text="action.title"></v-list-tile-title>
          <v-list-tile-action>
            <v-icon v-text="action.icon"></v-icon>
          </v-list-tile-action>
        </v-list-tile>
      </v-list-group>
      <v-list-tile
        v-for="item in items"
        :key="item.title"
        :to="{ name: item.route }"
        exact
      >
        <v-list-tile-action>
          <v-icon>{{ item.icon }}</v-icon>
        </v-list-tile-action>

        <v-list-tile-content>
          <v-list-tile-title>{{ item.title }}</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
  </v-navigation-drawer>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { Getter } from "vuex-class";
@Component({
  name: "Sidebar"
})
export default class extends Vue {
  @Getter("user/avatar") avatar!: string;
  @Getter("user/name") name!: string;
  drawer = true;
  items = [
    { title: "Home", icon: "dashboard", route: "dashboard" },
    { title: "Projects", icon: "question_answer", route: "projects" },
    { title: "Tasks", icon: "question_answer", route: "tasks" }
  ];

  account = [
    { title: "Settings", icon: "settings", route: "account-settings" }
  ];
  right = null;
}
</script>
