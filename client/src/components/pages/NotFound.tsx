import React from "react";

import { Link, RouteComponentProps } from "@reach/router";

const NotFound: React.FC<RouteComponentProps> = props => {
  return (
    <div className="container is-centered has-text-centered">
      <h1 className="title is-4">Page Not Found</h1>

      <div>
        <Link to="/" className="button is-primary">
          Home
        </Link>
      </div>
    </div>
  );
};

export default NotFound;
