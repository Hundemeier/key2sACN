<template>
  <fieldset>
    <legend>sACN <small v-if="dirty" class="text-warning">Not saved</small></legend>
    <table class="table table-hover">
      <thead>
        <tr>
          <th>Universe</th>
          <th>Multicast</th>
          <th>Destinations</th>
          <th>Stop</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in sACN" :key="item.universe" @click="click(item)" :class="{
          'table-active': selectedItem === null ? false : item.universe === selectedItem.universe}">
          <td>{{item.universe}}</td>
          <td>{{item.multicast ? "Yes" : "No"}}</td>
          <td>
            <span v-for="dest in item.destinations" :key="dest">{{dest}}<br/></span>
          </td>
          <td>
            <button class="btn btn-outline-secondary btn-sm full-width" @click="stop(item.universe)">Stop</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-if="error" class="alert alert-danger">
      Error! Could not load data!<br/>
      <small>The displayed data could have changed!</small>
    </div>
    <EditSacn
    :universe="selectedItem === null ? 1 : selectedItem.universe"
    :multicast="selectedItem === null ? false : selectedItem.multicast"
    :destinations="selectedItem === null ? '' : selectedItem.destinations.join()" />
  </fieldset>
</template>

<script>
import { GET_ALL_SACN } from '@/constants/queries'
import { STOP_SACN } from '@/constants/mutations'
import EditSacn from '@/components/EditSacn'

export default {
  name: 'sACN',
  components: { EditSacn },
  data () {
    return {
      sACN: [],
      error: false, // for holding error state
      selectedItem: null
    }
  },
  props: {
    dirty: {
      type: Boolean,
      default: false
    }
  },
  apollo: {
    sACN: {
      query: GET_ALL_SACN,
      pollInterval: 1000,
      manual: true,
      result ({ data, loading }) {
        if (!loading) {
          if (data.sACN === undefined) {
            this.error = true
          } else {
            var arr = data.sACN.slice()
            this.sACN = arr.sort((a, b) => a.universe - b.universe)
            this.error = false
          }
        }
      }
    }
  },
  methods: {
    stop (universe) {
      this.$apolloClient.mutate({
        mutation: STOP_SACN,
        variables: {
          'universe': universe
        }
      }).then(() => {
        this.$apollo.queries.sACN.refetch()
      })
    },
    click (item) {
      this.selectedItem = item
    }
  }
}
</script>

<style>
.full-width {
  width: 100%;
}
</style>
