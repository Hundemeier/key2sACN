// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
// import router from './router'

import './cyborg.min.css'
import './global.css'
import 'vue-snotify/styles/material.css'

import { ApolloClient } from 'apollo-client'
import { HttpLink } from 'apollo-link-http'
import { InMemoryCache } from 'apollo-cache-inmemory'
import VueApollo from 'vue-apollo'
import Snotify from 'vue-snotify'

const httpLink = new HttpLink({
  // uri: 'http://' + window.location.host + '/graphql'
  // uri: 'http://192.168.1.20:8080/graphql'
  uri: '/graphql'
})

const apolloClient = new ApolloClient({
  link: httpLink,
  cache: new InMemoryCache(),
  connectToDevTools: true
})
// store the client global for special use cases
Vue.prototype.$apolloClient = apolloClient

const apolloProvider = new VueApollo({
  defaultClient: apolloClient,
  defaultOptions: {
    $loadingKey: 'loading'
  }
})

Vue.use(Snotify)
Vue.use(VueApollo)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  // router,
  provide: apolloProvider.provide(),
  components: { App },
  template: '<App/>'
})
