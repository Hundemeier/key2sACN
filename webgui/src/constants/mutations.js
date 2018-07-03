import gql from 'graphql-tag'

export const SET_LISTEN = gql`mutation ($deviceID: Int!, $listen: Boolean!) {
  setListening(deviceID: $deviceID, listen: $listen) {
    listening
  }
}`

export const SET_SACN = gql`
mutation ($universe:Int!, $destinations: [String], $multicast:Boolean) {
  sACN(multicast: $multicast, universe: $universe, destinations: $destinations) {
    destinations
    multicast
    universe
  }
}`

export const STOP_SACN = gql`mutation ($universe:Int!) {
  stopSACN(universe: $universe)
}`
