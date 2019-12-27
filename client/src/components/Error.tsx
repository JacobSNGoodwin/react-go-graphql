import React from "react";
import { Link } from "@reach/router";

import Login from "./Login";

interface ErrorProps {
  messages?: string[];
  includeLogin?: boolean;
  includeHomeButton?: boolean;
}

const Error: React.FC<ErrorProps> = props => {
  return (
    <div className="container is-centered has-text-centered">
      <h1 className="title is-4">Error</h1>
      {props.messages &&
        props.messages.map((message, i) => (
          <p key={i} className="has-text-danger">
            {message}
          </p>
        ))}
      <br />

      {props.includeLogin && (
        <div>
          <Login />
        </div>
      )}
      <br />

      {props.includeHomeButton && (
        <div>
          <Link to="/" className="button is-primary">
            Home
          </Link>
        </div>
      )}
    </div>
  );
};

export default Error;
