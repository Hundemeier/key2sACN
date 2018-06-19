import React from 'react';
import { cardWidthREM } from './constants';
import { Mutation } from 'react-apollo';
import { writeConfig, deleteConfig } from './mutations';

function ConfigCard() {
  return (
    <div class="card text-white border-secondary" style={{ width: cardWidthREM + "rem" }}>
      <div class="card-header">Configuration <div class="text-muted">Server-side</div></div>
      <div class="row m-2">
        <Mutation mutation={writeConfig} >
          {(writeConf, { data }) => (
            <div class="col text-center">
              <button class="btn btn-secondary" onClick={e => {
                e.preventDefault();
                writeConf();
              }}>Save Configuration</button>
            </div>
          )}
        </Mutation>
        <Mutation mutation={deleteConfig} >
          {(delConf, { data }) => (
            <div class="col text-center">
              <button class="btn btn-secondary" onClick={e => {
                e.preventDefault();
                delConf();
              }}>Delete Configuration</button>
            </div>
          )}
        </Mutation>
      </div>
    </div>
  );
}

export default ConfigCard;