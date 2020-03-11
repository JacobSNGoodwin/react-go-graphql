import React from "react";
import { Link, RouteComponentProps } from "@reach/router";

import Login from "./Login";

const Error: React.FC<RouteComponentProps> = props => {
  const state: ErrorProps = props.location
    ? (props.location.state as ErrorProps)
    : {
        messages: ["Unknown error"],
        includeLogin: false
      };
  return (
    <div className="container is-centered has-text-centered">
      <h1 className="title is-4">Error</h1>
      {state.messages &&
        state.messages.map((message, i) => (
          <p key={i} className="has-text-danger">
            {message}
          </p>
        ))}
      <br />

      {state.includeLogin && (
        <div>
          <Login />
        </div>
      )}
      <br />

      <div>
        <Link to="/" className="button is-primary">
          Home
        </Link>
      </div>
    </div>
  );
};

export default Error;
