<template>
  <div>
    <button class="btn btn-sm mx-1"
    :class="{
      'btn-outline-success': !Dirty.keyMapDirty && !Dirty.listeningDirty && !Dirty.sACNdirty ,
      'btn-outline-warning': Dirty.keyMapDirty || Dirty.listeningDirty || Dirty.sACNdirty}"
    @click.prevent="saveConfig">Save Config</button>
    <button class="btn btn-sm btn-outline-danger mx-1" @click.prevent="delConfig">Delete Config</button>
  </div>
</template>

<script>
import { DEL_CONFIG, SAVE_CONFIG } from '@/constants/mutations'

export default {
  name: 'ConfigBtns',
  props: {
    Dirty: Object
  },
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
      }).catch(() => {
        this.$snotify.error('Could not write config file!', 'Error!')
      })
    },
    delConfig () {
      this.$apolloClient.mutate({
        mutation: DEL_CONFIG
      }).then(({data, error}) => {
        console.log(data)
        if (data.deleteConfig) {
          this.$snotify.success('Configuration file was successfully deleted from the server', 'Success!')
        } else {
          this.$snotify.error('Could not delete config file! Maybe it does not exist any more?', 'Error!')
        }
      }).catch(() => {
        this.$snotify.error('Could not delete config file! Maybe it does not exist any more?', 'Error!')
      })
    }
  }
}
</script>
