<template>
  <form method="POST">
    <fieldset>
      <legend>Add / Edit sACN</legend>
      <div>
        <label for="universe">Universe</label>
        <input type="number" class="form-control" name="universe" min="0" max="63999" required v-model="universeHolder">
      </div>
      <div class="my-3 custom-control custom-checkbox">
        <input class="custom-control-input" id="mutlicast" type="checkbox" v-model="multicastHolder">
        <label class="custom-control-label" for="mutlicast">Multicast</label>
      </div>
      <div>
        <label for="destinations">Destinations</label>
        <small>A comma seperated list of IP-Addresses</small>
        <input type="text" class="form-control" name="destinations" v-model="destinationsHolder" @input="checkDestinationField">
        <label for="destinations"><small class="text-muted">eg. '192.168.1.2,192.168.3.4'</small></label>
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
      universeHolder: 1, // for holding the information for the gui
      multicastHolder: false,
      destinationsHolder: ''
    }
  },
  props: {
    // the props are only used for the user selecting another item in the table
    universe: {
      type: Number,
      default: 1
    },
    multicast: {
      type: Boolean,
      default: true
    },
    destinations: {
      type: String,
      default: ''
    }
  },
  watch: {
    // watch for props changing
    universe (newVal, oldVal) {
      this.universeHolder = newVal
    },
    destinations (newVal, oldVal) {
      this.destinationsHolder = newVal
    },
    multicast (newVal, oldVal) {
      this.multicastHolder = newVal
    }
  },
  methods: {
    setSacn () {
      this.$apollo.mutate({
        mutation: SET_SACN,
        variables: {
          'universe': this.universeHolder,
          'destinations': this.destinationsHolder.split(','),
          'multicast': this.multicastHolder
        }
      })
    },
    checkDestinationField (event) {
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
