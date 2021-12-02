<template>
  <div>
    <div v-for="item in items" :key="item.id">
      <app-video :item="item" :video-id="item.id" />
      <br />
    </div>
    <a href.prevent="#" @click="loadMore">more</a>
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
  async fetch({ store }) {
    const payload = {
      uri: ROUTES.GET.POPULARS,
    }
    if (
      store.getters.getPopularVideos &&
      store.getters.getPopularVideos.length > 0
    ) {
      return
    }

    await store.dispatch('fetchPopularVideos', payload)
  },
  computed: {
    items() {
      return this.$store.getters.getPopularVideos
    },
    nextPageToken() {
      return this.$store.getters.getMeta.nextPageToken
    },
  },
  methods: {
    loadMore() {
      const payload = {
        uri: ROUTES.GET.POPULARS,
        params: {
          pageToken: this.nextPageToken,
        },
      }
      this.$store.dispatch('fetchPopularVideos', payload)
    },
  },
})
</script>

<style scoped>
.test {
  background: tomato;
}
</style>
