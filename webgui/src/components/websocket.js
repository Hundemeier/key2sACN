import React from 'react';
import { cardWidthREM } from './constants';

export default class WebsocketCard extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      eventList: []
    }

    //var socket = new WebSocket("ws://"+window.location.host+"/websocket");
    var socket = new WebSocket("ws://192.168.1.20:8080/websocket");
    console.log("created WS");
    socket.onmessage = e => {
      var event = JSON.parse(e.data)
      console.log(event);
      var newList = this.state.eventList.slice(-4);
      var newEvent = {
        Time: event.Time,
        Data: ""
      };

      if (event.Type === 1) {
        if (event.Data.Value === 0) {
          newEvent.Data = event.Data.KeyboardID + ":" + event.Data.KeyCode + " KEY UP";
        } else if (event.Data.Value === 1) {
          newEvent.Data = event.Data.KeyboardID + ":" + event.Data.KeyCode + " KEY DOWN";
        } else {
          return;
        }
      } else if (event.Type === 3) {
        newEvent.Data = "Delete Config: " + (event.Data ? "successful" : event.Error);
      } else if (event.Type === 2) {
        newEvent.Data = "Wrote Config: " + (event.Data ? "successful" : event.Error);
      } else {
        return;
      }

      newList.push(newEvent);
      this.setState({
        eventList: newList
      })
    }
  }

  render() {
    return (
      <div class="card text-white border-secondary" style={{ width: cardWidthREM + "rem" }}>
        <div class="card-header">Events</div>
        <ul>
          {this.state.eventList.map(event => <EventItem time={event.Time} event={event.Data} />)}
        </ul>
      </div>
    );
  }
}

function EventItem(props) {
  var date = new Date(props.time);
  return (
    <li class="list-item-no-dot text-console">
      {date.toLocaleTimeString()} :  {props.event}
    </li>
  );
}