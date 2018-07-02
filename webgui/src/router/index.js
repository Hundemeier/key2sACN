import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Devices from '@/components/Devices'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/devices',
      component: Devices
    },
    {
      path: '/sacn',
      component: HelloWorld
    }
  ]
})
