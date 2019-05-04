<template>
  <v-layout align-center justify-center>
    <v-flex xs12 sm8 md4>
      <v-card class="elevation-12">
        <v-toolbar dark color="primary">
          <v-toolbar-title>Login</v-toolbar-title>
        </v-toolbar>
        <v-card-text>
          <v-form ref="form">
            <v-text-field
              v-model="form.email"
              :rules="rules.email"
              label="E-mail"
              required
            ></v-text-field>
            <v-text-field
              v-model="form.password"
              :rules="rules.password"
              :append-icon="show ? 'visibility_off' : 'visibility'"
              @click:append="() => (show = !show)"
              :type="show ? 'text' : 'password'"
              label="Password"
              required
            ></v-text-field>
            <v-checkbox v-model="form.remember">
              <span slot="label">
                Remember
              </span>
            </v-checkbox>
            <v-btn
              :loading="isLoading"
              :disabled="isLoading"
              :type="'password'"
              :round="true"
              size="normal"
              @click="login"
            >
              Login
            </v-btn>
          </v-form>

          <p class="">
            You don't have an account?
            <RouterLink :to="{ name: 'register' }">
              Sign up here!
            </RouterLink>
          </p>
        </v-card-text>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { VMessages } from "vuetify/lib";
import { Getter } from "vuex-class";
import { QueueMessages } from "@/class/QueueMessages";
@Component({
  name: "Login"
})
export default class extends Vue {
  @Getter("user/queue") queue!: QueueMessages;
  isLoading = false;
  isError = false;
  show = false;
  form = {
    password: "",
    email: "",
    remember: false
  };

  rules = {
    email: [
      (input: string) => !!input || "E-mail is required",
      (input: string) =>
        /^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$/.test(input) ||
        "E-mail must be valid"
    ],
    password: [
      (input: string) => !!input || "Password is required",
      (input: string) =>
        input.length > 6 || "Password must be greater than 6 characters"
    ]
  };
  mounted() {}
  async login(event: Event) {
    event.preventDefault();
    this.isLoading = true;
    let form: any = this.$refs["form"];
    if (form.validate()) {
      console.log("LOGIN");
      try {
        let formData = this.form;
        const { data } = await this.axios.post("/api/auth/login", formData);

        // if user_data is available in response
        if (data.data) {
          this.$store.dispatch("auth/LOGIN", data.data);
          this.$router.push({ name: "dashboard" });
        }
      } catch (e) {
        console.log(e);
      }
    }
    this.isLoading = false;
  }
}
</script>
