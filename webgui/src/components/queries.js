import gql from "graphql-tag";

export const getDevices = gql`
{
  Devices {
    id
    listening
    name
  }
}`;

export const getSacn = gql`
{
  sACN {
    universe
    multicast
    destinations
  }
}`;