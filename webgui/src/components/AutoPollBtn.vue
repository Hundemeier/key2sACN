<template>
  <button class="btn" :class="{'btn-success': polling, 'btn-secondary': !polling}" @click.prevent="togglePolling">Auto Refresh</button>
</template>

<script>
import { POLLING } from '@/constants/constants'

export default {
  name: 'AutoPollBtn',
  data () {
    return {
      polling: true // the default is, that all the qeueries are polling.
      // Unfortunately there is no way to check wether the graphql
      // queries are polling or not at the moment
    }
  },
  methods: {
    // Unfinished! TODO
    togglePolling () {
      if (this.polling) {
        this.polling = false
        this.$apollo.queries.sACN.stopPolling()
      } else {
        this.polling = true
        this.$apollo.queries.sACN.startPolling(POLLING)
      }
    }
  }
}
</script>
