<template>
  <div class="activities">
    <h3 class="title is-3">Pull Requests</h3>

    <div class="activity-section" v-for="pullRequest in pullRequests" :key="pullRequest.id">
      <Activity :pullRequest=pullRequest />

      <hr />
    </div>
  </div>
</template>

<script>

import Activity from './Activity'

export default {
  name: "Activities",
  components: {
    Activity
  },
  props: {
    pullRequests: {
      type: Array,
      default: () => [],
    }
  }
};
</script>

<style scoped>
  .activities {
    display: flex;
    flex-direction: column;
    height: 100%;
    padding-top: 20px;
    padding-left: 30px;
    padding-right: 30px;
  }
</style>