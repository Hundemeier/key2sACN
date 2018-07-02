import gql from 'graphql-tag'

export const GET_ALL_DEVICES = gql`query {
  Devices{
    id
    name
    listening
  }
}`
