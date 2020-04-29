<template>
    <q-list>
        <q-item-label
          header
          class="text-grey-8"
        >
          Essential Links
        </q-item-label>
        <EssentialLink
          v-for="link in essentialLinks"
          :key="link.title"
          v-bind="link"
        />
        <q-item
          clickable
          tag="a"
          target="_blank"
          :href="link"
        >
          <q-item-section
            v-if="icon"
            avatar
          >
            <q-icon :name="icon" />
          </q-item-section>

          <q-item-section>
            <q-item-label>{{ title }}</q-item-label>
            <q-item-label caption>
              {{ caption }}
            </q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
</template>

<script>

import {createComponent} from '@vue/composition-api'
export default createComponent({
  name: 'Sidebar',
  props: {
    title: {
      type: String,
      required: true
    },

    caption: {
      type: String,
      default: ''
    },

    link: {
      type: String,
      default: '#'
    },

    icon: {
      type: String,
      default: ''
    }
  },

  setup(props,ctx){
    const links = {}

    return {
      links,
    }
  }
})
</script>
