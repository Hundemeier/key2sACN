import React from 'react';
import { Query } from 'react-apollo';
import { getSacn } from './queries';
import { cardWidthREM } from './constants';


function SacnTable(props) {
  return (
    <table class="table table-hover">
      <thead>
        <tr>
          <th>Universe</th>
          <th>Multicast</th>
          <th>Destinations</th>
        </tr>
      </thead>
      <tbody>
        {props.sACN.map((sACN) => <SacnItem sACN={sACN} />)}
      </tbody>
    </table>
  );
}

function SacnItem(props) {
  return (
    <tr key={props.sACN.universe} >
      <td>{props.sACN.universe}</td>
      <td>{props.sACN.multicast ? "Yes" : "No"}</td>
      <td>
      <SacnDestination destinations={props.sACN.destinations} />
      </td>
    </tr>
  );
}

function SacnDestination(props) {
  return (
    <ul class="list-group">
      {props.destinations.map((dest) => <li class="list-item-destination" key={dest}>{dest}</li>)}
    </ul>
  );
}

function SacnCard() {
  return (
    <div class="card text-white border-secondary" style={{ width: cardWidthREM + "rem" }}>
      <div class="card-header">sACN Output</div>
      <Query query={getSacn} pollInterval={1000} >
        {({ loading, error, data }) => {
          if (loading) return <span class="badge badge-info">Loading...</span>;
          if (error) return <span class="badge badge-danger">Error: could not load</span>;
          //sort the list of sACN universes
          var list = data.sACN.slice();
          list.sort((a, b) => {
            if (a.universe > b.universe) {
              return 1;
            } else if (a.universe === b.universe) {
              return 0;
            } else {
              return -1;
            }
          })
          return <SacnTable sACN={list} />;
        }}
      </Query>
    </div>
  );
}

export default SacnCard;