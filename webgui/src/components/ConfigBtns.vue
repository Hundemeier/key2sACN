<template>
  <div>
    <button class="btn btn-sm btn-outline-success mx-1" @click.prevent="saveConfig">Save Config</button>
    <button class="btn btn-sm btn-outline-danger mx-1" @click.prevent="delConfig">Delete Config</button>
  </div>
</template>

<script>
import { DEL_CONFIG, SAVE_CONFIG } from '@/constants/mutations'

export default {
  name: 'ConfigBtns',
  methods: {
    saveConfig () {
      this.$apolloClient.mutate({
        mutation: SAVE_CONFIG
      }).then(({data}) => {
        if (data.writeConfig) {
          this.$snotify.success('Configuration file was successfully writen on the server', 'Success!')
        } else {
          this.$snotify.error('Could not write config file!', 'Error!')
        }
      })
    },
    delConfig () {
      this.$apolloClient.mutate({
        mutation: DEL_CONFIG
      })
    }
  }
}
</script>
