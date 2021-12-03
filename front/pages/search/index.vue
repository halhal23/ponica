<template>
  <div>
    <nuxt-link to="/">トップページへ</nuxt-link>
    <div v-for="item in items" :key="item.etag">
      <app-video :item="item" :video-id="item.id.videoId" />
      <br />
    </div>
    <a href.prevent @click="loadMore">More</a>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import ROUTES from '~/routes/api'
import AppVideo from '~/components/AppVideo.vue'
export default Vue.extend({
  components: {
    AppVideo,
  },
  async fetch({ store, query }) {
    const q = encodeURIComponent(query.q?.toString()) || ''
    if (store.getters.getSearchedItems.length > 0) {
      return
    }
    const payload = {
      uri: ROUTES.GET.SEARCH,
      params: {
        q,
      },
    }
    await store.dispatch('fetchSearchedItems', payload)
  },
  computed: {
    items() {
      return this.$store.getters.getSearchedItems
    },
    nextPageToken() {
      return this.$store.getters.getSearchedMeta.nextPageToken
    },
  },
  methods: {
    async loadMore() {
      const q = encodeURIComponent(this.$route.query.q?.toString())
      const payload = {
        uri: ROUTES.GET.SEARCH,
        params: {
          q,
          pageToken: this.nextPageToken,
        },
      }
      await this.$store.dispatch('fetchSearchedItems', payload)
    },
  },
})
</script>
