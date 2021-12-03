import { createRequestClient } from './request-client'
export const state = () => ({
  items: [],
  item: {},
  relatedItems: [],
  meta: {},
})

export const actions = {
  async fetchPopularVideos({ commit }, payload) {
    const client = createRequestClient(this.$axios)
    const res = await client.get(payload.uri, payload.params)
    commit('mutatePopularVideos', res)
  },
  async findVideo({ commit }, payload) {
    const client = createRequestClient(this.$axios)
    const res = await client.get(payload.uri, payload.params)
    const params = {
      ...res.video_list,
    }
    commit('mutateVideo', params)
  },
  async fetchRelatedVideos({ commit }, payload) {
    const client = createRequestClient(this.$axios)
    const res = await client.get(payload.uri, payload.params)
    commit('mutateRelatedVideos', res)
  },
}

export const mutations = {
  mutatePopularVideos(state, payload) {
    state.items = payload.items ? state.items.concat(payload.items) : []
    state.meta = payload
  },
  mutateVideo(state, payload) {
    const params =
      payload.items && payload.items.length > 0 ? payload.items[0] : {}
    state.item = params
  },
  mutateRelatedVideos(state, payload) {
    state.relatedItems = []
    payload.items.forEach((item) => {
      if (item.snippet) {
        state.relatedItems.push(item)
      }
    })
  },
}

export const getters = {
  getPopularVideos(state) {
    return state.items
  },
  getMeta(state) {
    return state.meta
  },
  getVideo(state) {
    return state.item
  },
  getRelatedVideos(state) {
    return state.relatedItems
  },
}
