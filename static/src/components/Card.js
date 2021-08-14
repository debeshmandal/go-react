import React from "react";

function Card(props) {
    return (
      <div className="col-sm-4" key={props.index}>
        <div className="card mb-4">
          <div className="card-header">{props.name}</div>
          <div className="card-body">{props.description}</div>
          <div className="card-footer">
          </div>
        </div>
      </div>
    )
};

export default Card;