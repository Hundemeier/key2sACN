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