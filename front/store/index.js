import { createRequestClient } from './request-client'
import firebase from '~/plugins/firebase'
export const state = () => ({
  items: [],
  item: {},
  relatedItems: [],
  searchedItems: [],
  searchedMeta: {},
  meta: {},
  token: '',
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
  async fetchSearchedItems({ commit }, payload) {
    const client = createRequestClient(this.$axios)
    const res = await client.get(payload.uri, payload.params)
    commit('mutateSearchedItems', res)
  },
  async signUp({ commit }, payload) {
    console.log('--------1---------')
    await firebase
      .auth()
      .createUserWithEmailAndPassword(payload.email, payload.password)
    const res = await firebase
      .auth()
      .signInWithEmailAndPassword(payload.email, payload.password)
    const token = await res.user.getIdToken()
    this.$cookies.set('jwt_token', token)
    commit('mutateToken', token)
    this.app.router.push('/')
  },
  setToken({ commit }, payload) {
    commit('mutateToken', payload)
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
  mutateSearchedItems(state, payload) {
    state.searchedItems = payload.items
      ? state.searchedItems.concat(payload.items)
      : []
    state.searchedMeta = payload
  },
  mutateToken(state, payload) {
    state.token = payload
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
  getSearchedItems(state) {
    return state.searchedItems
  },
  getSearchedMeta(state) {
    return state.searchedMeta
  },
  getToken(state) {
    return state.token
  },
}
