<template>
  <form method="POST">
    <fieldset>
      <legend>Add / Edit Mapping</legend>
      <div class="row my-2">
        <div class="col-6">
          <label for="universe">Universe</label>
          <input type="number" class="form-control" name="universe" min="0" max="63999" required v-model="internalHolder.universe">
        </div>
        <div class="col-6">
          <label for="channel">Channel</label>
          <input type="number" class="form-control" name="channel" min="0" max="511" required v-model="internalHolder.channel">
        </div>
      </div>
      <div class="row my-2">
        <div class="col-6">
          <label for="keyboardID">Device ID</label>
          <input type="number" class="form-control" name="keyboardID" min="0" max="65535" required v-model="internalHolder.keyboardID">
        </div>
        <div class="col-6">
          <label for="keycode">Key Code <small class="text-primary">{{keyMap[internalHolder.keycode]}}</small></label>
          <input type="number" class="form-control" name="keycode" min="0" max="65535" required v-model="internalHolder.keycode">
        </div>
      </div>
      <button id="submit" class="btn btn-secondary mt-3" @click.prevent="setMapping">Add / Edit</button>
    </fieldset>
  </form>
</template>

<script>
import { SET_MAPPING } from '@/constants/mutations'
import { keyMap } from '@/constants/constants'

export default {
  name: 'EditMapping',
  data () {
    return {
      keyMap: keyMap,
      internalHolder: {
        universe: 1,
        channel: 0,
        keyboardID: 0,
        keycode: 0
      }
    }
  },
  props: ['value'], // value is the prop from v-model
  // why v-model? because parent and child share the same state for the input data.
  // so the parent can manipulate the data that is in the input of this component
  watch: {
    value: {
      handler (newVal, oldVal) {
        this.internalHolder.universe = newVal.universe
        this.internalHolder.channel = newVal.channel
        this.internalHolder.keyboardID = newVal.keyboardID
        this.internalHolder.keycode = newVal.keycode
      },
      deep: true // also search the properites of the holder
    },
    internalHolder: {
      // if internalHolder is changed, make event to parent that the two-way data binding
      handler (newVal, oldVal) {
        this.$emit('input', {
          universe: newVal.universe,
          channel: newVal.channel,
          keyboardID: newVal.keyboardID,
          keycode: newVal.keycode
        })
      },
      deep: true
    }
  },
  methods: {
    setMapping () {
      this.$apollo.mutate({
        mutation: SET_MAPPING,
        variables: {
          'universe': this.internalHolder.universe,
          'channel': this.internalHolder.channel,
          'keyboardID': this.internalHolder.keyboardID,
          'keycode': this.internalHolder.keycode
        }
      })
    }
  }
}
</script>
