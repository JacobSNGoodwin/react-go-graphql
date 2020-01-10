import React from "react";

import { RouteComponentProps, Redirect } from "@reach/router";
import { AuthContext } from "../components/contexts/AuthContext";

interface IPrivateRouteProps extends RouteComponentProps {
  as: React.FC;
  admin?: boolean;
  editor?: boolean;
}

const PrivateRoute: React.FC<IPrivateRouteProps> = props => {
  const { user } = React.useContext(AuthContext);

  if (!user) {
    const errorProps: ErrorProps = {
      messages: ["You are not logged in"],
      includeLogin: true
    };

    return <Redirect to="error" state={errorProps} noThrow />;
  }

  let { as: Comp } = props;

  let allowed: boolean = false;

  if (props.admin && user.roles.admin) {
    allowed = true;
  }

  if (props.editor && user.roles.editor) {
    allowed = true;
  }

  if (user && !allowed) {
    const errorProps: ErrorProps = {
      messages: ["You are not allowed to access this resrouce"],
      includeLogin: true
    };
    return <Redirect to="error" state={errorProps} noThrow />;
  }

  return <Comp {...props} />;
};

export default PrivateRoute;
