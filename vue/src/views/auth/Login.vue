<template>
  <div>
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
        :type="'password'"
        label="Password"
        required
      ></v-text-field>
      <v-checkbox v-model="form.remember"></v-checkbox>
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
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { VMessages } from "vuetify/lib";
import { Mutation } from "vuex-class";
import { Message } from "@/class/QueueMessages";
@Component({
  name: "Login"
})
export default class extends Vue {
  @Mutation("user/QUEUE_GLOBAL_MESSAGE") queueMessage!: (
    message: Message
  ) => void;
  isLoading = false;
  isError = false;
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
  mounted() {
    this.queueMessage({ text: "TEST", "message-type": "success" });
  }
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
