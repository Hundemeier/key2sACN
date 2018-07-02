<template>
  <div id="app">
    <ul class="nav nav-tabs">
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
