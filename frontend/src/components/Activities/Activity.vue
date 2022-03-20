<template>
	<div id="activity">
    <h4 class='title is-4'>{{ pullRequest.title }}</h4>
    <p>{{ displayDaysAgo() }}</p>
  </div>
</template>

<script>

export default {
  name: "Activity",
  props: {
    pullRequest: {
      type: Object,
      default: () => { return {} },
    }
  },
  methods: {
    displayDaysAgo () {
      return `Open ${this.relativeDay()} ago.`
    },
    relativeDay () {
      const now = new Date().getTime()
      const createdAt = Date.parse(this.pullRequest.createdAt)
      const time = Math.abs(now - createdAt)

      let humanTime, units;

      // If there are years
      if (time > (1000 * 60 * 60 * 24 * 365)) {
        humanTime = parseInt(time / (1000 * 60 * 60 * 24 * 365), 10);
        units = 'years';
      }

      // If there are months
      else if (time > (1000 * 60 * 60 * 24 * 30)) {
        humanTime = parseInt(time / (1000 * 60 * 60 * 24 * 30), 10);
        units = 'months';
      }

      // If there are weeks
      else if (time > (1000 * 60 * 60 * 24 * 7)) {
        humanTime = parseInt(time / (1000 * 60 * 60 * 24 * 7), 10);
        units = 'weeks';
      }

      // If there are days
      else if (time > (1000 * 60 * 60 * 24)) {
        humanTime = parseInt(time / (1000 * 60 * 60 * 24), 10);
        units = 'days';
      }

      // If there are hours
      else if (time > (1000 * 60 * 60)) {
        humanTime = parseInt(time / (1000 * 60 * 60), 10);
        units = 'hours';
      }

      // If there are minutes
      else if (time > (1000 * 60)) {
        humanTime = parseInt(time / (1000 * 60), 10);
        units = 'minutes';
      }

      // Otherwise, use seconds
      else {
        humanTime = parseInt(time / (1000), 10);
        units = 'seconds';
      }

      return humanTime + ' ' + units;


    }
  }
};
</script>

<style scoped>
  #activity {
    display: flex;
    flex-direction: column;
    height: 100%;
  }
</style>