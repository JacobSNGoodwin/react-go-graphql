import React from "react";
import { ApolloError } from "apollo-boost";
import { Redirect } from "@reach/router";
import { AuthContext } from "./contexts/AuthContext";

interface ErrorListProps {
  error: ApolloError;
}

const ErrorsList: React.FC<ErrorListProps> = props => {
  const authContext = React.useContext(AuthContext);
  const errorMessages: JSX.Element[] = [];
  let redirectProps: ErrorProps | undefined = undefined;
  props.error.graphQLErrors.forEach((error, i) => {
    if (error.extensions && error.extensions.type === "FORBIDDEN") {
      redirectProps = {
        messages: ["You are not authenticated"],
        includeLogin: true
      };
    }

    errorMessages.push(
      <p key={i} className="has-text-danger">
        {error.message}
      </p>
    );
  });

  if (redirectProps) {
    authContext.logout();
    return <Redirect to="error" state={redirectProps} noThrow />;
  }

  return (
    <div className="container has-text-centered">
      <h1 className="title is-4">Error</h1>
      {errorMessages}
    </div>
  );
};

export default ErrorsList;
