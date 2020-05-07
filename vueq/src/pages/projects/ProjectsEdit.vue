<template>
  <q-page padding>

    <h4 class="Title"> {{isNew ? 'New' : 'Edit'}} Project</h4>

    <div class="q-px-xl">

      <div class="row">
        <q-form
          ref="refForm"
          class="col-6 offset-3 q-px-xl"
          greedy
        >
          <q-input
            v-model="form.name"
            class="q-py-sm"
            name="name"
            color="teal"
            outlined
            clearable
            label="Name"
            style="max-width:300px;"
            :error="Boolean(formErrors.name).valueOf()"
            :error-message="formErrors.name"
            :rules="validation.name"
          />
          <q-input
            v-model="form.description"
            class="q-py-md"
            name="description"
            color="teal"
            outlined
            type="textarea"
            label="Description"
            :rules="validation.description"
            :error="Boolean(formErrors.description).valueOf()"
            :error-message="formErrors.description"
          />

          <!-- ... -->
        </q-form>

        <div class="col-12 text-center q-px-lg">
          <q-btn
            :loading="isDoingWork"
            :disabled="isDoingWork"
            @click="save"
          >
            Save
          </q-btn>
        </div>
      </div>

    </div>
    <q-inner-loading :showing="isLoading">
      <q-spinner-gears
        size="50px"
        color="primary"
      />
    </q-inner-loading>
  </q-page>
</template>

<script lang="ts">
import {
  defineComponent,
  computed,
  reactive,
  toRefs,
  onMounted,
  ref,
  Ref,
} from '@vue/composition-api'
import { QForm } from 'quasar'

const validation = {
  name: [
    (value: string) =>
      (value && value.trim().length >= 4) || 'Minim length of name is 4',
  ],
  description: [
    (value: string) =>
      (value && value.trim().length >= 100) ||
      'Minim length of description is 100',
    (value: string) =>
      (value && value.trim().length <= 255) ||
      'Maxim length of description is 255',
  ],
}
const model = {
  name: '',
  description: '',
}
export default defineComponent({
  name: 'ProjectsEdit',
  setup(props, ctx) {
    const { $utils, axios, $route } = ctx.root
    const state = reactive({
      isLoading: false,
      isEditing: false,
      form: { ...model },
      formErrors: { ...model },
      validation,
    })
    const isNew = computed(() => $route.name === 'project-new')

    const ID = computed(() => $route.params.id)

    const isDoingWork = computed(() => state.isLoading || state.isEditing)

    const refForm: Ref<null | QForm> = ref(null)

    const save = () => {
      $utils.request({
        vm: state,
        call: async () => {
          if (!refForm.value) return

          if (!(await refForm.value.validate())) return

          if (isNew.value) {
            await axios.post(`/api/projects`, state.form)

            state.form = { ...model }

            refForm.value.reset()
          } else {
            await axios.put(`/api/projects/${ID.value}`, state.form)
          }
        },
        loading: 'isEditing',
      })
    }

    const loadProject = () => {
      $utils.request({
        vm: state,
        call: async () => {
          const { data } = await axios.get(`/api/projects/${ID.value}`)

          state.form.name = data.data.name
          state.form.description = data.data.description
        },
      })
    }

    onMounted(() => {
      if (!isNew.value) {
        loadProject()
      }
    })

    return {
      isNew,
      save,
      isDoingWork,
      ...toRefs(state),
      refForm,
    }
  },
})
</script>
