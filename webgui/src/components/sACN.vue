<template>
  <fieldset>
    <legend>sACN</legend>
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
        <tr v-for="item in sACN" :key="item.universe">
          <td>{{item.universe}}</td>
          <td>{{item.multicast ? "Yes" : "No"}}</td>
          <td>
            <ul class="no-points" v-for="dest in item.destinations" :key="dest">
              <li>{{dest}}</li>
            </ul>
          </td>
          <td>
            <button class="btn btn-outline-secondary btn-sm full-width" @click="stop(item.universe)">Stop</button>
          </td>
        </tr>
      </tbody>
    </table>
    <EditSacn />
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
      sACN: []
    }
  },
  apollo: {
    sACN: {
      query: GET_ALL_SACN,
      pollInterval: 1000
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
    }
  }
}
</script>

<style>
.no-points {
  list-style-type: none;
}

.full-width {
  width: 100%;
}
</style>
