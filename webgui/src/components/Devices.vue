<template>
  <fieldset>
    <legend>Devices</legend>
    <table class="table table-hover" id="devicesTable">
      <thead>
        <tr>
          <th scope="col">ID</th>
          <th scope="col">Name</th>
          <th scope="col">Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="device in Devices" :key="device.id" :class="{'table-success': device.listening}">
          <td>{{device.id}}</td>
          <td>{{device.name}}</td>
          <td v-if="device.listening">
            <button class="btn btn-outline-secondary btn-sm" style="width: 100%" @click="setListen(device.id, false)">Stop</button>
          </td>
          <td v-else>
            <button class="btn btn-outline-secondary btn-sm" style="width: 100%" @click="setListen(device.id, true)">Listen</button>
          </td>
        </tr>
      </tbody>
    </table>
  </fieldset>
</template>

<script>
import { GET_ALL_DEVICES } from '@/constants/queries'
import { SET_LISTEN } from '@/constants/mutations'

export default {
  name: 'Devices',
  data () {
    return {
      Devices: []
    }
  },
  apollo: {
    Devices: {
      query: GET_ALL_DEVICES,
      pollInterval: 1000
    }
  },
  methods: {
    setListen (id, listen) {
      this.$apolloClient.mutate({
        mutation: SET_LISTEN,
        variables: {
          'listen': listen,
          'deviceID': id
        }
      }).then(() => {
        this.$apollo.queries.Devices.refetch()
      })
    }
  }
}
</script>
