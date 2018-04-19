import React, { Component } from "react";

import Card from "./Card";

class Answers extends Component {
  render() {
    return (
      <div className="m-3">
        <div className="border rounded border-secondary">
          <Card text="goodbye" />
        </div>
        <div className="border rounded border-secondary">
          <Card text="hello" />
        </div>
        <div className="border rounded border-secondary">
          <Card text="home" />
        </div>
        <div className="border rounded border-secondary">
          <Card text="computer" />
        </div>
      </div>
    );
  }
}

export default Answers;
