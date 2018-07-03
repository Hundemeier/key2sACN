<template>
  <div id="app" class="bg-dark">
    <nav class="navbar navbar-expand-sm navbar-dark bg-dark sticky-top">
      <a class="navbar-brand" href="#">key2sACN</a>
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
        <li class="nav-item">
          <a class="nav-link" href="#config">Configuration</a>
        </li>
      </ul>
    </nav>
    <div class="row" style="width: 100%">
      <div id="devices" class="col-sm-6 col-xl-4">
        <Devices />
      </div>
      <div id="sacn" class="col-sm-6 col-xl-4">
        <sACN />
      </div>
      <div style="height: 1000px"></div>
    </div>
  </div>
</template>

<script>
import Devices from '@/components/Devices'
import sACN from '@/components/sACN'

export default {
  name: 'App',
  components: {
    Devices,
    sACN
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
