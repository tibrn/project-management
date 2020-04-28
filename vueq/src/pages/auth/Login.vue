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
            <q-form class="q-px-sm q-pt-xl">
              <q-input
                square
                clearable
                v-model="formLogin.email"
                type="email"
                label="Email"
                :error="Boolean(formErrorsLogin.email).valueOf()"
                :error-message="formErrorsLogin.email"
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
                :error="Boolean(formErrorsLogin.password).valueOf()"
                :error-message="formErrorsLogin.password"
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
            <q-form class="q-px-sm q-pt-xl q-pb-lg">
              <q-input
                square
                clearable
                v-model="form.email"
                type="email"
                label="Email"
                :error="Boolean(formErrors.email).valueOf()"
                :error-message="formErrors.email"
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
import Vue from "vue"
import { createComponent, reactive } from '@vue/composition-api'
import { createNamespacedHelpers } from 'vuex'
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

function keepReactive (arg: any, target :object):object{


  Object.keys(arg).map(key => {
    Object.defineProperty(target, key, {
      get() {
        return arg[key]
      },
      set(val: any) {
        arg[key] = val
      }
    })
  })

  return target
}
export default createComponent({
  name:'Login',
  setup(props, ctx) {
    const state = reactive({
      isLoading: false,
      form: { ...modelForm },
      formErrors: { ...modelForm },
      formLogin: { ...modelFormLogin },
      formLoginErrors: { ...modelFormLogin },
    })

    const { axios, $utils, $router } = ctx.root

    const { SET_USER } = mapMutations(['SET_USER'])

    const register = async () => {
      $utils.request({
        vm: this,
        call: async () => {
          const { data } = await axios.post("/api/users", state.form)

          if (typeof data !== "undefined") {
            SET_USER(data.data)
          }

          state.formErrors = { ...modelForm }

          $router.push({ path: "/dashboard" })
        }
      })
    }

    const login = () => {
      $utils.request({
        vm: this,
        call: async () => {
          const { data } = await axios.post("/api/login", state.formLogin)
          if (typeof data !== "undefined") {
            SET_USER(data.data)
          }

          state.formLoginErrors = { ...modelFormLogin }

          $router.push({ path: "/dashboard" })
        }
      })
    }

    return keepReactive(state,{
      register,
      login,
    })
  }

  // @Mutation("user/SET_USER") setUser?: (user: any) => {};

})
</script>
