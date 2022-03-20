<template>
  <div id="app">
    <div class="information-section">
      <div class="field">
        <label class="label">Username</label>
          <div class="control">
            <input
              class="input is-success"
              type="text"
              placeholder="Text input"
              value="Enter Text"
              v-model="login"
            >
          </div>
      </div>

      <button class="button is-link" @click='onSubmit'>Enter</button>
    </div>

    <hr />

    <div class="activities-section">
      <Loader v-if="isLoading" />
      <Activities v-else :pullRequests="pullRequests" />
    </div>
  </div>
</template>

<script>

import Loader from '@/components/Loader.vue'
import Activities from '@/components/Activities'

export default {
  name: "app",
  components: {
    Loader,
    Activities,
  },
  data () {
    return {
      login: "jawt94",
      pullRequests: []
    }
  },
  computed: {
    isLoading () {
      return this.pullRequests.length == 0
    }
  },
  methods: {
    async onSubmit () {
      this.pullRequests = await window.backend.userPullRequests(this.login) || []
    },
  },
};
</script>

<style scoped>

.information-section {
  padding: 5%;
}

.activities-section {
  height: 100%;
  padding-bottom: 20px;
}

</style>