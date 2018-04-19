import React, { Component } from "react";
import Turn from "./card/Turn";

class App extends Component {
  render() {
    return (
      <div className="container">
        <div className="App">
          <Turn />
        </div>
      </div>
    );
  }
}

export default App;
