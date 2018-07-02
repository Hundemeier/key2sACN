import gql from 'graphql-tag'

export const SET_LISTEN = gql`mutation ($deviceID: Int!, $listen: Boolean!) {
  setListening(deviceID: $deviceID, listen: $listen) {
    listening
  }
}`
