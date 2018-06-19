import React from 'react';
import { Mutation } from 'react-apollo';
import { cardWidthREM } from './constants';
import { mutateSacn } from './mutations';


class SacnConfig extends React.Component {
  constructor(props) {
    super(props);
    this.state = {destinations: "", universe: 1};
    this.handleInputChange = this.handleInputChange.bind(this);
  }

  handleInputChange(event) {
    const target = event.target;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const name = target.name;

    if (name === "destinations") {
      var btn = document.getElementById("sACNconfigBtn");
      //check input validity
      if (value.match(/[a-z]|\s/i)) {
        //letters found
        target.setCustomValidity('Given string is not a valid IP enumeration!');
        btn.disabled = true;
      } else {
        target.setCustomValidity('');
        btn.disabled = false;
      }
    }

    this.setState({
      [name]: value
    });
  }

  render() {
    var multicast;
    return (
      <div class="card text-white border-secondary" style={{ width: cardWidthREM + "rem" }}>
        <div class="card-header">sACN Add/Edit</div>
        <div class="container">
          <form id="form">
            <div class="form-group">
              <label htmlFor="universe">Universe</label>
              <input class="form-control" type="number" name="universe" min="0" max="63999" id="universe" value={this.state.universe} onChange={this.handleInputChange} />
            </div>
            <div class="form-check">
              <input class="form-check-input" type="checkbox" id="multicast" ref={
                node => { multicast = node; }} />
              <label htmlFor="multicast" class="form-check-label">Multicast</label>
            </div>
            <div class="form-group">
              <label htmlFor="destinations">Destinations <small class="text-muted">Seperate by , eg: "192.168.2.1, 192.168.2.2"</small></label>
              <input class="form-control" name="destinations" type="textarea" id="destinations" value={this.state.destinations} onChange={this.handleInputChange} />
            </div>
            <Mutation mutation={mutateSacn}>
              {(mutateSacn, { data }) => (
                <button class="btn btn-secondary m-2" id="sACNconfigBtn" onClick={e => {
                  e.preventDefault();
                  mutateSacn({
                    variables: {
                      "universe": this.state.universe,
                      "multicast": multicast.checked,
                      "destinations": this.state.destinations.split(",")
                    }
                  });
                }}>Add / Edit</button>
              )}
            </Mutation>
          </form>
        </div>
      </div>
    );
  }
}

export default SacnConfig;