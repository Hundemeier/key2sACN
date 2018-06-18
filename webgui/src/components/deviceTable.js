import React from 'react';
import { Query, Mutation } from 'react-apollo';
import { getDevices } from './queries';
import { setListen } from './mutations';
import { cardWidthREM } from './constants';

function DeviceTable(props) {
  return (
    <table class="table table-hover">
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        {props.devices.map((device) => <DeviceItem device={device} />)}
      </tbody>
    </table>
  );
}

function DeviceItem(props) {
  //Helper function for less code: sets the listening of a deivce to true or false
  function sendListen(event, setFunc, listen) {
    event.preventDefault();
    setFunc({
      variables: {
        "id": props.device.id,
        "listen": listen
      }
    });
  }
  return (
    <tr class={props.device.listening ? "table-success" : ""} key={props.device.id}>
      <td>{props.device.id}</td>
      <td>{props.device.name}</td>
      <td>
        <Mutation mutation={setListen} >
          {(setListen, { data }) => (
            <div>
              {props.device.listening ?
                <button class="btn btn-outline-secondary btn-sm" style={{width: 100 +"%"}} onClick={e => sendListen(e, setListen, false)}>Stop</button> :
                <button class="btn btn-outline-secondary btn-sm" style={{width: 100 +"%"}} onClick={e => sendListen(e, setListen, true)}>Listen</button>
              }
            </div>
          )}
        </Mutation>
      </td>
    </tr>
  );
}


function DeviceCard() {
  return (
    <div class="card text-white border-secondary" style={{ width: cardWidthREM + "rem" }}>
      <div class="card-header">Devices<div class="text-muted">Stoping may take a second</div></div>
      <Query query={getDevices} pollInterval={1000} >
        {({ loading, error, data }) => {
          if (loading) return <span class="badge badge-info">Loading...</span>;
          if (error) return <span class="badge badge-danger">Error: could not load</span>;
          return <DeviceTable devices={data.Devices} />;
        }}
      </Query>
    </div>
  );
}

export default DeviceCard;