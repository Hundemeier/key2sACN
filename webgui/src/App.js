import React from 'react';
//import logo from './logo.svg';
import './App.css';
import './cyborg.min.css';
import Clock from './components/clock';
import DeviceCard from './components/deviceTable';
import SacnCard from './components/sACNtable';
import { ApolloProvider } from "react-apollo";
import ApolloClient from "apollo-boost";
import ConfigCard from './components/configTable';
import SacnConfig from './components/sACNconfigurator';
import MapCard, { KeyCodes } from './components/mappingTable';
import { MapConfig } from './components/mapConfig';
import WebsocketCard from './components/websocket';

const client = new ApolloClient({
  uri: "/graphql"
});

class App extends React.Component {
  render() {
    return (
      <ApolloProvider client={client}>
        <Clock />
        <div class="row">
          <SacnCard />
          <SacnConfig />
          <DeviceCard />
          <MapCard />
          <MapConfig />
          <WebsocketCard />
          <ConfigCard />
          <KeyCodes />
        </div>
      </ApolloProvider>
    );
  }
}

export default App;

/*class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React</h1>
        </header>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
      </div>
    );
  }
}

export default App;*/
