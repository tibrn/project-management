<template>
  <v-layout align-center justify-center>
    <v-flex xs12 sm8 md4>
      <v-card class="elevation-12">
        <v-toolbar dark color="primary">
          <v-toolbar-title>Settings</v-toolbar-title>
        </v-toolbar>
        <v-card-text>
          <v-form ref="form">
            <v-select
              v-model="settings.theme"
              :items="themesTypes"
              label="Theme"
              solo
            ></v-select>
            <v-btn
              :loading="isLoading"
              :disabled="isLoading"
              :type="'password'"
              :round="true"
              size="normal"
              @click="updateSettings"
            >
              Update
            </v-btn>
          </v-form>
        </v-card-text>
      </v-card>
    </v-flex>
  </v-layout>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { Mutation, Getter } from "vuex-class";
@Component({
  name: "AccountSettings"
})
export default class extends Vue {
  @Mutation("user/UPDATE_THEME") updateTheme!: (theme: string) => void;
  @Getter("user/avatar") avatar!: string;
  @Getter("user/theme") theme!: string;

  isLoading = false;
  settings = {
    avatar: "",
    theme: ""
  };

  themesTypes = ["dark", "light"];

  mounted() {
    this.settings.theme = this.theme;
    this.settings.avatar = this.avatar;
  }

  async updateSettings() {
    this.isLoading = true;
    try {
      let { data } = await this.axios.post("/api/user/settings", this.settings);

      if (data) {
        console.log(data);
        this.updateTheme(data.theme);
      }
    } catch (e) {
      console.log(e);
    }

    this.isLoading = false;
  }
}
</script>
