<template>
  <form>
    <fieldset>
      <legend>Add / Edit sACN</legend>
      <div>
        <label for="universe">Universe</label>
        <input type="number" class="form-control" name="universe" min="0" max="63999" value="1">
      </div>
      <div class="my-3 custom-control custom-checkbox">
        <input class="custom-control-input" id="mutlicast" checked="" type="checkbox">
        <label class="custom-control-label" for="mutlicast">Multicast</label>
      </div>
      <div>
        <label for="destinations">Destinations</label>
        <small>A comma seperated list of IP-Addresses</small>
        <input type="text" class="form-control" name="destinations" placeholder="eg. '192.168.1.2,192.168.3.4'">
      </div>
    </fieldset>
  </form>
</template>

<script>
import { SET_SACN } from '@/constants/mutations'

export default {
  name: 'EditSacn',
  methods: {
    setSacn () {
      this.$apollo.mutate({
        mutation: SET_SACN,
        variables: {
          'universe': 1,
          'destinations': ['192.168.1.6'],
          'multicast': true
        }
      }).then(() => {
        this.$apollo.queries.sACN.refetch()
      })
    }
  }
}
</script>
