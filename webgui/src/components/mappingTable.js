import React from 'react';
import { Query, Mutation } from 'react-apollo';
import { deleteMapping } from './mutations';
import { getMapping } from './queries';
import { cardWidthREM } from './constants';

function MapItem(props) {
  return (
    <tr>
      <td>{props.map.keyboardID}</td>
      <td>{props.map.keycode}</td>
      <td>{props.map.universe}</td>
      <td>{props.map.channel}</td>
      <td>
        <Mutation mutation={deleteMapping}>
          {(delMap, { data }) => (
            <button class="btn btn-outline-secondary btn-sm" onClick={e => {
              e.preventDefault();
              delMap({
                variables: {
                  "keycode": props.map.keycode,
                  "keyboardID": props.map.keyboardID
                }
              });
            }}>Delete</button>
          )}
        </Mutation>
      </td>
    </tr>
  );
}

function MapTable(props) {
  return (
    <table class="table table-hover">
      <thead>
        <tr>
          <th>Keyboard ID</th>
          <th>Keycode</th>
          <th>Universe</th>
          <th>Channel</th>
          <th>Delete</th>
        </tr>
      </thead>
      <tbody>
        {props.mapping.map((map) => <MapItem map={map} />)}
      </tbody>
    </table>
  );
}

function MapCard() {
  return (
    <div class="card text-white border-secondary" style={{ width: cardWidthREM + "rem" }}>
      <div class="card-header">Mapping</div>
      <Query query={getMapping} pollInterval={1000} >
        {({ loading, error, data }) => {
          if (loading) return <span class="badge badge-info">Loading...</span>;
          if (error) return <span class="badge badge-danger">Error: could not load</span>;

          var list = data.Mapping.slice();
          list.sort((a, b) => {
            if (a.keyboardID > b.keyboardID) {
              return 1;
            } else if (a.keyboardID < b.keyboardID) {
              return -1;
            } else {
              //a.keyboardID === b.keyboardID; sort with keycode if keyboardID is the same
              if (a.keycode > b.keycode) {
                return 1
              } else if (a.keycode < b.keycode) {
                return -1;
              } else {
                return 0;
              }
            }
          })
          return <MapTable mapping={list} />;
        }}
      </Query>
    </div>
  );
}

export default MapCard;


export class KeyCodes extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      show: false
    };
    this.toggleShow = this.toggleShow.bind(this);
  }

  toggleShow() {
    this.setState({
      show: !this.state.show
    })
  }

  render() {
    return (
      <div class="card text-white border-secondary" style={{ width: cardWidthREM + "rem" }}>
        <div class="card-header">KeyCodes</div>
        <div class="row">
          <div class="col text-center">
            <button class="btn btn-secondary" onClick={e => {
              e.preventDefault();
              this.toggleShow();
            }}>Toggle KeyCodes</button>
          </div>
        </div>
        {this.state.show ? 
        <table class="table table-hover">
          <thead>
            <tr>
              <th>KeyCode</th>
              <th>Key</th>
            </tr>
          </thead>
          <tbody>
            {Object.keys(keyMap).map((key, ind) => (
              <tr>
                <td>{key}</td>
                <td>{keyMap[key]}</td>
                </tr>
            ))}
          </tbody>
        </table>
        : ""}
      </div>
    );
  }
}

const keyMap = {
  "1": "ESC",
  "2": "1",
  "3": "2",
  "4": "3",
  "5": "4",
  "6": "5",
  "7": "6",
  "8": "7",
  "9": "8",
  "10": "9",
  "11": "0",
  "12": "-",
  "13": "=",
  "14": "BS",
  "15": "TAB",
  "16": "Q",
  "17": "W",
  "18": "E",
  "19": "R",
  "20": "T",
  "21": "Y",
  "22": "U",
  "23": "I",
  "24": "O",
  "25": "P",
  "26": "[",
  "27": "]",
  "28": "ENTER",
  "29": "L_CTRL",
  "30": "A",
  "31": "S",
  "32": "D",
  "33": "F",
  "34": "G",
  "35": "H",
  "36": "J",
  "37": "K",
  "38": "L",
  "39": ";",
  "40": "'",
  "41": "`",
  "42": "L_SHIFT",
  "43": "\\",
  "44": "Z",
  "45": "X",
  "46": "C",
  "47": "V",
  "48": "B",
  "49": "N",
  "50": "M",
  "51": ",",
  "52": ".",
  "53": "/",
  "54": "R_SHIFT",
  "55": "*",
  "56": "L_ALT",
  "57": "SPACE",
  "58": "CAPS_LOCK",
  "59": "F1",
  "60": "F2",
  "61": "F3",
  "62": "F4",
  "63": "F5",
  "64": "F6",
  "65": "F7",
  "66": "F8",
  "67": "F9",
  "68": "F10",
  "69": "NUM_LOCK",
  "70": "SCROLL_LOCK",
  "71": "HOME",
  "72": "UP_8",
  "73": "PGUP_9",
  "74": "-",
  "75": "LEFT_4",
  "76": "5",
  "77": "RT_ARROW_6",
  "78": "+",
  "79": "END_1",
  "80": "DOWN",
  "81": "PGDN_3",
  "82": "INS",
  "83": "DEL",
  "84": "",
  "85": "",
  "86": "",
  "87": "F11",
  "88": "F12",
  "89": "",
  "90": "",
  "91": "",
  "92": "",
  "93": "",
  "94": "",
  "95": "",
  "96": "R_ENTER",
  "97": "R_CTRL",
  "98": "/",
  "99": "PRT_SCR",
  "100": "R_ALT",
  "101": "",
  "102": "Home",
  "103": "Up",
  "104": "PgUp",
  "105": "Left",
  "106": "Right",
  "107": "End",
  "108": "Down",
  "109": "PgDn",
  "110": "Insert",
  "111": "Del",
  "112": "",
  "113": "",
  "114": "",
  "115": "",
  "116": "",
  "117": "",
  "118": "",
  "119": "Pause"
};