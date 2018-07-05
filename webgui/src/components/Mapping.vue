<template>
  <fieldset>
    <legend>Mapping</legend>
    <table class="table table-hover">
      <thead>
        <tr>
          <th>Device</th>
          <th>KeyCode</th>
          <th>Universe.Channel</th>
          <th>Change</th>
          <th>Delete</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="map in Mapping" :key="map.keycode+(map.keyboardID*65536)">
          <td>{{map.keyboardID}}</td>
          <td>{{map.keycode}} <small class="text-primary">{{keyMap[map.keycode]}}</small></td>
          <td>{{map.universe}}.{{map.channel}}</td>
          <td>
            <button class="btn btn-outline-secondary btn-sm full-width" @click="setEdit(map)">Change</button>
          </td>
          <td>
            <button class="btn btn-outline-secondary btn-sm full-width" @click="deleteMapping(map.keycode, map.keyboardID)">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-if="error" class="alert alert-danger">
      Error! Could not load data!<br/>
      <small>The displayed data could have changed!</small>
    </div>
    <EditMapping
    v-model="selectedItem" />
  </fieldset>
</template>

<script>
import { GET_ALL_MAPPING } from '@/constants/queries'
import { DEL_MAPPING } from '@/constants/mutations'
import { keyMap } from '@/constants/constants'
import EditMapping from '@/components/EditMapping'

export default {
  name: 'Mapping',
  components: {
    EditMapping
  },
  data () {
    return {
      Mapping: [],
      keyMap: keyMap,
      selectedItem: {
        universe: 1,
        channel: 0,
        keyboardID: 0,
        keycode: 0
      },
      error: false
    }
  },
  methods: {
    deleteMapping (keycode, keyboardID) {
      this.$apolloClient.mutate({
        mutation: DEL_MAPPING,
        variables: {
          'keycode': keycode,
          'keyboardID': keyboardID
        }
      }).then(() => {
        this.$apollo.queries.Mapping.refetch()
      })
    },
    setEdit (item) {
      this.selectedItem = item
      // this.$set(this.selectedItem, 'universe', item.universe)
    }
  },
  apollo: {
    Mapping: {
      query: GET_ALL_MAPPING,
      pollInterval: 1000,
      manual: true,
      result ({ data, loading }) {
        if (!loading) {
          if (data.Mapping === undefined) {
            this.error = true
          } else {
            var arr = data.Mapping.slice()
            this.Mapping = arr.sort((a, b) => (a.keycode + (a.keyboardID * 65536)) - (b.keycode + (b.keyboardID * 65536)))
            this.error = false
          }
        }
      }
    }
  }
}
</script>
