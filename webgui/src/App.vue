<template>
  <div id="app" class="bg-dark">
    <ul class="nav nav-tabs bg-secondary">
      <li class="nav-item">
        <router-link class="nav-link" :class="{ active: this.$router.currentRoute.path === '/sacn' }" v-bind:to="'/sacn'">sACN</router-link>
      </li>
      <li class="nav-item">
        <router-link class="nav-link" :class="{ active: this.$router.currentRoute.path === '/devices' }" v-bind:to="'/devices'">Devices</router-link>
      </li>
      <li class="nav-item">
        <router-link class="nav-link" :class="{ active: this.$router.currentRoute.path === '/mapping' }" v-bind:to="'/mapping'">Mapping</router-link>
      </li>
    </ul>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
.nav-tabs {
  padding-top: 1em;
  padding-left: 1em;
  margin-bottom: 1em;
}
</style>
