import React from "react";
import Cookies from "js-cookie";

import { RouteComponentProps, Redirect } from "@reach/router";
import { AuthContext } from "./contexts/AuthContext";

interface IPrivateRouteProps extends RouteComponentProps {
  as: React.FC;
  admin?: boolean;
  editor?: boolean;
}

const PrivateRoute: React.FC<IPrivateRouteProps> = props => {
  const { user, logout } = React.useContext(AuthContext);
  const userCookie = Cookies.get("userinfo");

  if (!user || !userCookie) {
    const errorProps: ErrorProps = {
      messages: ["You are not logged in"],
      includeLogin: true
    };

    logout();

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
