<template>
  <div id="app" class="bg-dark">
    <nav class="navbar navbar-expand-sm navbar-dark bg-dark sticky-top">
      <ul class="navbar-nav mr-auto">
        <li class="nav-item">
          <a class="nav-link" href="#devices">Devices</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#sacn">sACN</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#mapping">Mapping</a>
        </li>
      </ul>
      <div class="row align-items-center">
        <div class="mx-2">
          <config-btns :Dirty="Dirty"></config-btns>
        </div>
        <div class="mx-2">
          <clock></clock>
        </div>
      </div>
    </nav>
    <div class="row" style="width: 100%">
      <div id="devices" class="col-lg-6 col-xl-4">
        <Devices :dirty="Dirty.listeningDirty" />
      </div>
      <div id="sacn" class="col-lg-6 col-xl-4">
        <sACN :dirty="Dirty.sACNdirty" />
      </div>
      <div id="mapping" class="col-lg-6 col-xl-4">
        <Mapping :dirty="Dirty.keyMapDirty" />
      </div>
    </div>
    <vue-snotify></vue-snotify>
  </div>
</template>

<script>
import Devices from '@/components/Devices'
import sACN from '@/components/sACN'
import Mapping from '@/components/Mapping'
// import AutoPollBtn from '@/components/AutoPollBtn'
import Clock from '@/components/Clock'
import ConfigBtns from '@/components/ConfigBtns'
import { GET_CONFIG_DIRTY } from '@/constants/queries'

export default {
  name: 'App',
  components: {
    Devices,
    sACN,
    Mapping,
    Clock,
    ConfigBtns
  },
  data () {
    return {
      Dirty: { // inital values:
        keyMapDirty: false,
        listeningDirty: false,
        sACNdirty: false
      }
    }
  },
  apollo: {
    Dirty: {
      query: GET_CONFIG_DIRTY,
      pollInterval: 1000
    }
  }
}
</script>

<style>
.nav-tabs {
  padding-top: 1em;
  padding-left: 1em;
  margin-bottom: 1em;
}
</style>
