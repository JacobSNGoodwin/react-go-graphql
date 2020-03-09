import React from "react";

import { RouteComponentProps } from "@reach/router";

const Home: React.FC<RouteComponentProps> = () => {
  return (
    <div className="container is-centered has-text-centered">
      <h1 className="title is-4">Store Administrator</h1>
      <p>Helping small stores manage their inventory</p>
      <p>With nearly useless home pages</p>
    </div>
  );
};

export default Home;
