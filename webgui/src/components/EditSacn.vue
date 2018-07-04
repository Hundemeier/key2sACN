<template>
  <form>
    <fieldset>
      <legend>Add / Edit sACN</legend>
      <div>
        <label for="universe">Universe</label>
        <input type="number" class="form-control" name="universe" min="0" max="63999" v-model="universe">
      </div>
      <div class="my-3 custom-control custom-checkbox">
        <input class="custom-control-input" id="mutlicast" type="checkbox" v-model="multicast">
        <label class="custom-control-label" for="mutlicast">Multicast</label>
      </div>
      <div>
        <label for="destinations">Destinations</label>
        <small>A comma seperated list of IP-Addresses</small>
        <input type="text" class="form-control" name="destinations" v-model="destinations" placeholder="eg. '192.168.1.2,192.168.3.4'" @input="checkDestinationField">
      </div>
      <button id="submit" class="btn btn-secondary mt-3" @click.prevent="setSacn()">Add / Edit</button>
    </fieldset>
  </form>
</template>

<script>
import { SET_SACN } from '@/constants/mutations'

export default {
  name: 'EditSacn',
  data () {
    return {
      universe: 1,
      destinations: '',
      multicast: false
    }
  },
  methods: {
    setSacn () {
      this.$apollo.mutate({
        mutation: SET_SACN,
        variables: {
          'universe': this.universe,
          'destinations': this.destinations.split(','),
          'multicast': this.multicast
        }
      }).then(() => {
        this.$apollo.queries.sACN.refetch()
      })
    },
    checkDestinationField (event) {
      console.log(event)
      if (event.target.value.match(/[a-z]|\s/i)) {
        // letters found
        event.target.setCustomValidity('Given string is not a valid IP enumeration!')
        document.getElementById('submit').disabled = true
      } else {
        event.target.setCustomValidity('')
        document.getElementById('submit').disabled = false
      }
    }
  }
}
</script>
