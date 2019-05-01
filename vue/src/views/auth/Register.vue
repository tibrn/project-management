<template>
  <div>
    <v-form ref="form">
      <v-text-field
        v-model="form.name"
        :rules="rules.name"
        label="Name"
        required
      ></v-text-field>
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
      <v-text-field
        v-model="form.password_confirmation"
        :rules="rules.password_confirmation"
        :type="'password'"
        label="Password Confirmation"
        required
      ></v-text-field>
      <v-btn
        :loading="isLoading"
        :disabled="isLoading"
        :type="'password'"
        :round="true"
        size="normal"
        @click="register"
      >
        Register
      </v-btn>
    </v-form>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";

@Component({
  name: "Register"
})
export default class extends Vue {
  isLoading = false;
  isError = false;
  form = {
    name: "",
    password: "",
    password_confirmation: "",
    email: ""
  };

  rules = {
    name: [
      (input: string) => !!input || "Name is required",
      (input: string) => /^[a-zA-Z]*$/.test(input) || "Name must be valid"
    ],
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
    ],
    password_confirmation: [
      (input: string) => !!input || "Password is required",
      (input: string) =>
        input.length > 6 || "Password must be greater than 6 characters"
    ]
  };

  async register(event: Event) {
    event.preventDefault();
    this.isLoading = true;

    let form: any = this.$refs["form"];
    if (form.validate()) {
      try {
        let formData = this.form;
        const { data } = await this.axios.post("/api/user", formData);

        // if user_data is available in response
        if (data.data) {
          this.$store.dispatch("auth/LOGIN", data.data);
          this.$router.push({ name: "dashboard" });
        }
      } catch (e) {
        console.log(e)
      }
    }
    this.isLoading = false;
  }
}
</script>
