import React, { Component } from "react";

import Card from "./Card";

class Question extends Component {
  render() {
    return (
      <div className="border border-primary">
        <Card text="привет" />
      </div>
    );
  }
}

export default Question;
