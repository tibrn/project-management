<template>
  <div>
    <v-layout row wrap>
      <v-flex xs12>
        <v-card>
          <v-card-title
            class="red darken-3 white--text subheading"
            primary-title
          >
            <span class="title">Oops!</span>
          </v-card-title>
          <v-card-text>
            Looks like you're in uncharted territory. We don't know about this
            page yet. Luckily, we know the way back.
          </v-card-text>
          <v-card-actions>
            <!-- <router-link to="/foo">Go to Foo</router-link> -->
            <v-btn flat class="red darken-3 white--text" to="/dashboard"
              >Back</v-btn
            >
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";

@Component({
  name: "NotFound"
})
export default class extends Vue {
  goBack() {}
}
</script>
