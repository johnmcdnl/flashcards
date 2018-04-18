import React, { Component } from "react";

class Card extends Component {
  render() {
    return (
      <div className="card m-3">
        <div className="card-body">
          <p className="card-text mx-auto">{this.props.text}</p>
        </div>
      </div>
    );
  }
}

export default Card;
