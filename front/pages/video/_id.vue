<template>
  <div>
    <nuxt-link to="/">to top page</nuxt-link>
    <p>channel: {{ item.snippet.channelTitle }}</p>
    <p>title: {{ item.snippet.title }}</p>
    <p>description: {{ item.snippet.description | omit }}</p>
    <p>publishedAt: {{ item.snippet.publishedAt }}</p>
    <youtube ref="youtube" :video-id="$route.params.id"></youtube>
  </div>
</template>
<script lang="ts">
import Vue from 'vue'
import ROUTES from '~/routes/api'

export default Vue.extend({
  filters: {
    omit: (value: any) => {
      if (!value) {
        return ''
      }
      if (value.length > 120) {
        value = value.substr(0, 120) + '...'
      }
      return value
    },
  },
  async fetch({ store, route }) {
    const payload = {
      uri: ROUTES.GET.VIDEO.replace(':id', route.params.id),
    }
    await store.dispatch('findVideo', payload)
  },
  computed: {
    item() {
      return this.$store.getters.getVideo
    },
  },
})
</script>
