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
