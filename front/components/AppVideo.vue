<template>
  <div class="card">
    <div class="card-l">
      <img
        class="is-maxxmax"
        :src="item.snippet.thumbnails.medium.url"
        alt=""
      />
    </div>
    <div class="card-r">
      <small>{{ item.snippet.title }}</small>
      <hr />
      <small>{{ item.snippet.description | omit }}</small>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  filters: {
    omit: (value: any) => {
      if (!value) {
        return ''
      }
      if (value.length > 120) {
        return value.substr(0, 120) + '...'
      }
      return value
    },
  },
  props: {
    item: {
      type: Object,
      required: true,
    },
    videoId: {
      type: String,
      required: true,
    },
  },
}
</script>

<style scoped>
.card {
  display: flex;
  width: 600px;
  height: 180px;
  border: 1px solid gray;
}
.card-l {
  width: 40%;
}
.card-r {
  width: 60%;
}
.is-maxxmax {
  width: 100%;
  height: 100%;
}
</style>
