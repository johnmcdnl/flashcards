import React, { Component } from "react";

import Question from "./Question";
import Answers from "./Answers";

class Turn extends Component {
  render() {
    return (
      <div className="container">
        <Question />
        <Answers />
      </div>
    );
  }
}

export default Turn;
