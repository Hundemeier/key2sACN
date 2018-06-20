import React from 'react';
import { Mutation } from 'react-apollo';
import { setMapping } from './mutations';
import { cardWidthREM } from './constants';
import { keyMap } from './mappingTable';

export class MapConfig extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      "keycode": 0,
      "keyboardID": 0,
      "channel": 0,
      "universe": 1
    }
    this.handleInputChange = this.handleInputChange.bind(this);
  }

  handleInputChange(event) {
    const target = event.target;
    const value = target.value;
    const name = target.name;

    this.setState({
      [name]: value
    });
  }

  render() {
    return (
      <div class="card text-white border-secondary" style={{ width: cardWidthREM + "rem" }}>
        <div class="card-header">Mapping Edit/Add</div>
        <div class="container">
          <form id="form">
            <div className="form-row">
                <div class="form-group">
                  <label htmlFor="keycode">KeyCode <small class="text-primary">{keyMap[this.state.keycode]}</small></label>
                  <input class="form-control" type="number" name="keycode" min="0" max="65535" id="keycode" value={this.state.keycode} onChange={this.handleInputChange} />
                </div>
                  <div class="form-group">
                    <label htmlFor="keyboardID">Keyboard ID</label>
                    <input class="form-control" type="number" name="keyboardID" min="0" max="65535" id="keyboardID" value={this.state.keyboardID} onChange={this.handleInputChange} />
                  </div>
              </div>
              <div className="form-row">
                <div class="form-group">
                  <label htmlFor="universe">Universe</label>
                  <input class="form-control" type="number" name="universe" min="0" max="63999" id="universe" value={this.state.universe} onChange={this.handleInputChange} />
                </div>
                <div class="form-group">
                  <label htmlFor="channel">Channel</label>
                  <input class="form-control" type="number" name="channel" min="0" max="511" id="channel" value={this.state.channel} onChange={this.handleInputChange} />
                </div>
            </div>
            <Mutation mutation={setMapping}>
              {(mutateMapping, { data }) => (
                <button class="btn btn-secondary m-2" onClick={e => {
                  e.preventDefault();
                  mutateMapping({
                    variables: {
                      "universe": this.state.universe,
                      "channel": this.state.channel,
                      "keycode": this.state.keycode,
                      "keyboardID": this.state.keyboardID
                    }
                  })
                }}>Add / Edit</button>
              )}
            </Mutation>
          </form>
        </div>
      </div>
    );
  }
}