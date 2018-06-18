import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(<App />, document.getElementById('root'));
registerServiceWorker();



/*
client
  .query({
    query: gql`
    {
        Devices {
          id
          listening
          name
        }
      }
    `
  })
  .then(result => console.log(result));*/