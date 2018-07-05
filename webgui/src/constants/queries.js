import gql from 'graphql-tag'

export const GET_ALL_DEVICES = gql`{
  Devices{
    id
    name
    listening
  }
}`

export const GET_ALL_SACN = gql`{
  sACN {
    destinations
    multicast
    universe
  }
}`

export const GET_ALL_MAPPING = gql`{
  Mapping{
    channel
    universe
    keycode
    keyboardID
  }
}`
