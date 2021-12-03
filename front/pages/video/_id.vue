<template>
  <div>
    <nuxt-link to="/">to top page</nuxt-link>
    <p>channel: {{ item.snippet.channelTitle }}</p>
    <p>title: {{ item.snippet.title }}</p>
    <p>description: {{ item.snippet.description | omit }}</p>
    <p>publishedAt: {{ item.snippet.publishedAt }}</p>
    <youtube ref="youtube" :video-id="$route.params.id"></youtube>
    <div>関連動画</div>
    <div class="related-wrapper">
      <div
        v-for="rItem in relatedItems"
        :key="rItem.id.videoId"
        class="related-card"
      >
        <div class="related-card-l">
          <img :src="rItem.snippet.thumbnails.default.url" class="is-maxxmax" />
        </div>
        <div class="related-card-r">
          <div class="sm-txt">{{ rItem.snippet.title }}</div>
          <nuxt-link :to="`/video/${rItem.id.videoId}`">詳細</nuxt-link>
        </div>
      </div>
    </div>
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
    await store.dispatch('fetchRelatedVideos', {
      uri: ROUTES.GET.RELATEDS.replace(':id', route.params.id),
    })
  },
  computed: {
    item() {
      return this.$store.getters.getVideo
    },
    relatedItems() {
      return this.$store.getters.getRelatedVideos
    },
  },
})
</script>

<style scoped>
.related-wrapper {
  display: flex;
  width: 100%;
  height: 100px;
}
.related-card {
  height: 100px;
  width: 300px;
  border: 1px solid black;
  margin-right: 5px;
  display: flex;
}
.related-card-l {
  width: 50%;
  height: 100%;
}
.related-card-r {
  width: 50%;
  height: 100%;
}
.is-maxxmax {
  height: 100%;
  width: 100%;
}
.sm-txt {
  font-size: 5px;
}
</style>
