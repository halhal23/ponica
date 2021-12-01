<template>
  <div>index content</div>
</template>

<script lang="ts">
import Vue from 'vue'
import ROUTES from '~/routes/api'

export default Vue.extend({
  // eslint-disable-next-line require-await
  async fetch({ store }) {
    const payload = {
      uri: ROUTES.GET.POPULARS,
    }

    return store.dispatch('fetchPopularVideos', payload)
  },
})
</script>
