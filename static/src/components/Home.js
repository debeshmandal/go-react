import React, { useState } from "react";
import { FiThumbsUp, FiThumbsDown } from 'react-icons/fi';

const Home = () => {
  const [voted, setVoted] = useState(['']);
  const [products, setProducts] = useState([
    {
      id: 1,
      Name: "World of Authcraft",
      Slug: "world-of-authcraft",
      Description:
        "Battle bugs and protect yourself from invaders while you explore a scary world with no security",
    },
    {
      id: 2,
      Name: "Ocean Explorer",
      Slug: "ocean-explorer",
      Description:
        "Explore the depths of the sea in this one of a kind underwater experience",
    },
  ]);

  const vote = (type) => {
    alert(type);
  }

  return (
    <div className="container">
      <div className="jumbotron text-center mt-5">
        <h1>We R VR</h1>
        <p>
          Hi, Pal! Below you'll find the latest games that need feedback. Please provide honest feedback so developers can make the best games.
        </p>
        <div className="row">
          {products.map(function (product) {
            return (
              <div className="col-sm-4">
                <div className="card">
                  <div className="card-header">
                    {product.Name}
                    <span className="float-left">{voted}</span>
                  </div>
                  <div className="card-body">{product.Description}</div>
                  <div className="card-footer">
                    <a onClick={() => vote("Upvoted")} className="btn btn-default float-left">
                        <FiThumbsUp />
                    </a>
                    <a onClick={() => vote("Downvoted")} className="btn btn-default float-right">
                      <FiThumbsDown />
                    </a>
                  </div>
                </div>
              </div>  
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default Home;