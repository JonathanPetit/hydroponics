import React, { Component } from 'react';
import './App.css';
import Temperature from './Temperature';
import Humidity from './Humidity';

class App extends Component {
  render() {
    return (
      <React.Fragment>
        <Temperature/>
        <Humidity/>
      </React.Fragment>
        
    );
  }
}

export default App;