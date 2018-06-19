import gql from "graphql-tag";

export const setListen = gql`
mutation setListening($id: Int!, $listen: Boolean!){
  setListening(deviceID: $id, listen: $listen) {
    name id listening
  }
}`;

export const writeConfig = gql`
mutation {
  writeConfig
}`;

export const deleteConfig = gql`
mutation {
  deleteConfig
}`;

export const stopSacn = gql`
mutation sACN($universe:Int) {
  stopSACN(universe:$universe)
}`;

export const mutateSacn = gql`
mutation sACN($universe:Int!, $destinations:[String], $multicast:Boolean) {
  sACN(universe: $universe, destinations: $destinations, multicast: $multicast) {
    universe
    multicast
    destinations
  }
}`;

export const setMapping = gql`
mutation KeyMap($universe:Int!, $channel:Int!, $keycode:Int!, $keyboardID:Int!) {
  KeyMap(universe: $universe, channel:$channel, keycode:$keycode, keyboardID:$keyboardID) {
    universe
    channel
    keycode
    keyboardID
  }
}`;

export const deleteMapping = gql`
mutation deleteKeyMap($keycode:Int!, $keyboardID:Int!) {
  deleteKeyMap(keycode: $keycode, keyboardID:$keyboardID)
}`;