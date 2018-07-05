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

export const SET_MAPPING = gql`
mutation($universe:Int!,$channel:Int!,$keycode:Int!,$keyboardID:Int!) {
  KeyMap(universe:$universe, channel:$channel, keycode:$keycode, keyboardID:$keyboardID) {
    channel
    universe
    keycode
    keyboardID
  }
}`

export const DEL_MAPPING = gql`mutation($keycode:Int!,$keyboardID:Int!) {
  deleteKeyMap(keycode:$keycode, keyboardID:$keyboardID)
}`
