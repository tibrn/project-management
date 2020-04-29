<template>
  <q-page
    class="window-height window-width row justify-center items-center"
    style="background: linear-gradient(#8274C5, #5A4A9F);"
  >
    <div class="column q-pa-lg">
      <div class="row">
        <q-card
          square
          class="shadow-24"
          style="width:300px;height:650px;"
        >
          <q-card-section class="bg-deep-purple-7">
            <h4 class="text-h5 text-white q-my-md">Company &amp; Co</h4>
            <div
              class="absolute-bottom-right q-pr-md"
              style="transform: translateY(50%);"
            >
              <q-btn
                fab
                icon="add"
                color="purple-4"
              />
            </div>
          </q-card-section>
          <q-card-section>
            <q-form ref="elFormLogin" greedy class="q-px-sm q-pt-xl">
              <q-input
                square
                clearable
                v-model="formLogin.email"
                type="email"
                label="Email"
                :error="Boolean(formLoginErrors.email).valueOf()"
                :error-message="formLoginErrors.email"
                :rules="formLoginRules.email"
              >
                <template v-slot:prepend>
                  <q-icon name="email" />
                </template>
              </q-input>
              <q-input
                square
                clearable
                v-model="formLogin.password"
                type="password"
                label="Password"
                :error="Boolean(formLoginErrors.password).valueOf()"
                :error-message="formLoginErrors.password"
                :rules="formLoginRules.password"
              >
                <template v-slot:prepend>
                  <q-icon name="lock" />
                </template>
              </q-input>
            </q-form>
          </q-card-section>
          <q-card-section>
            <div class="text-center q-pa-md q-gutter-md">
              <q-btn
                round
                color="indigo-7"
                :disabled="isLoading"
              >
                <q-icon
                  name="fab fa-facebook-f"
                  size="1.2rem"
                />
              </q-btn>
              <q-btn
                round
                color="red-8"
                :disabled="isLoading"
              >
                <q-icon
                  name="fab fa-google"
                  size="1.2rem"
                />
              </q-btn>
            </div>
          </q-card-section>
          <q-card-actions class="q-px-lg">
            <q-btn
              unelevated
              size="lg"
              color="purple-4"
              class="full-width text-white"
              label="Sign In"
              :loading="isLoading"
              :disabled="isLoading"
              @click="login"
            />
          </q-card-actions>
          <q-card-section class="text-center q-pa-sm">
            <p class="text-grey-6">Forgot your password?</p>
          </q-card-section>
        </q-card>
      </div>
    </div>
    <div class="column q-pa-lg">
      <div class="row">
        <q-card
          square
          class="shadow-24"
          style="width:300px;height:650px;"
        >
          <q-card-section class="bg-deep-purple-7">
            <h4 class="text-h5 text-white q-my-md">Registration</h4>
            <div
              class="absolute-bottom-right q-pr-md"
              style="transform: translateY(50%);"
            >
              <q-btn
                fab
                icon="close"
                color="purple-4"
              />
            </div>
          </q-card-section>
          <q-card-section>
            <q-form ref="elForm" greedy class="q-px-sm q-pt-xl q-pb-lg">
              <q-input
                square
                clearable
                v-model="form.email"
                type="email"
                label="Email"
                :error="Boolean(formErrors.email).valueOf()"
                :error-message="formErrors.email"
                :rules="formRules.email"
              >
                <template v-slot:prepend>
                  <q-icon name="email" />
                </template>
              </q-input>
              <q-input
                square
                clearable
                v-model="form.name"
                type="name"
                label="Name"
                :error="Boolean(formErrors.name).valueOf()"
                :error-message="formErrors.name"
                :rules="formRules.name"
              >
                <template v-slot:prepend>
                  <q-icon name="person" />
                </template>
              </q-input>
              <q-input
                square
                clearable
                v-model="form.surname"
                type="surname"
                label="Surname"
                :error="Boolean(formErrors.surname).valueOf()"
                :error-message="formErrors.surname"
                :rules="formRules.surname"
              >
                <template v-slot:prepend>
                  <q-icon name="person" />
                </template>
              </q-input>
              <q-input
                square
                clearable
                v-model="form.password"
                type="password"
                label="Password"
                :error="Boolean(formErrors.password).valueOf()"
                :error-message="formErrors.password"
                :rules="formRules.password"
              >
                <template v-slot:prepend>
                  <q-icon name="lock" />
                </template>
              </q-input>
              <q-input
                square
                clearable
                v-model="form.password_confirmation"
                type="password"
                label="Password Confirmation"
                :error="Boolean(formErrors.password_confirmation).valueOf()"
                :error-message="formErrors.password_confirmation"
                :rules="formRules.password_confirmation(form,'password')"
              >
                <template v-slot:prepend>
                  <q-icon name="lock" />
                </template>
              </q-input>
            </q-form>
          </q-card-section>
          <q-card-actions class="q-px-lg">
            <q-btn
              unelevated
              size="lg"
              color="purple-4"
              class="full-width text-white"
              label="Get Started"
              :loading="isLoading"
              :disabled="isLoading"
              @click="register"
            />
          </q-card-actions>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<script lang="ts">

import { defineComponent, reactive, toRefs, ref, Ref } from '@vue/composition-api'
import { createNamespacedHelpers } from 'vuex'
import { QForm } from 'quasar'
const { mapMutations } = createNamespacedHelpers('user')
const modelForm = {
  email: "",
  name: "",
  surname: "",
  password: "",
  password_confirmation: ""
}
const modelFormLogin = {
  email: "",
  password: ""
}

const validation = {
  email: [(value: string) => /^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(value) || 'The email is invalid'],
  name: [(value: string) => value.length > 3 || 'Name minimum length is 3'],
  surname: [(value: string) => value.length > 3 || 'Surname minimum length is 3'],
  password: [(value: string) => (/^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])(?=.{6,})/.test(value)) || 'The password is not strong enough'],
  password_confirmation: (state: any, prop: string) => {
    return [function(value: string) {
      return value === state[prop] || 'Password does not match confirmation'
    }]
  }
}

export default defineComponent({
  name: 'Login',
  setup(props, ctx) {
    console.log(validation)
    const state = reactive({
      isLoading: false,

      form: { ...modelForm },
      formRules: validation,
      formErrors: { ...modelForm },
      formLogin: { ...modelFormLogin },
      formLoginErrors: { ...modelFormLogin },
      formLoginRules: { email: validation.email, password: validation.password }
    })

    const elForm: Ref<null |QForm > = ref(null)
    const elFormLogin: Ref<null |QForm > = ref(null)

    const { axios, $utils, $router } = ctx.root

    const { SET_USER } = mapMutations(['SET_USER'])

    const register = (): void => {
      $utils.request({
        vm: state,
        call: async () => {
          if (!elForm.value || !await elForm.value.validate()) return

          const { data } = await axios.post("/api/users", state.form)

          if (typeof data !== "undefined") {
            SET_USER(data.data)
          }

          state.formErrors = { ...modelForm }

          $router.push({ path: "/dashboard" })
        }
      })
    }

    const login = (): void => {
      $utils.request({
        vm: state,
        call: async () => {
          if (!elFormLogin.value || !await elFormLogin.value.validate()) return

          const { data } = await axios.post("/api/login", state.formLogin)
          if (typeof data !== "undefined") {
            SET_USER(data.data)
          }

          state.formLoginErrors = { ...modelFormLogin }

          $router.push({ path: "/dashboard" })
        }
      })
    }

    return {
      login,
      register,
      elForm,
      elFormLogin,
      ...toRefs(state),
    }
  }

  // @Mutation("user/SET_USER") setUser?: (user: any) => {};

})
</script>
