import React from "react";

import { navigate, RouteComponentProps } from "@reach/router";
import { AuthContext } from "../components/contexts/AuthContext";
import { hasRequiredRole } from "../util/util";
import Login from "./Login";

interface IPrivateRouteProps {
  as: React.FC;
  allowedRoles: string[];
}

const PrivateRoute: React.FC<IPrivateRouteProps &
  RouteComponentProps> = props => {
  const authContext = React.useContext(AuthContext);
  if (!authContext.user) {
    return <Login />;
  }

  let { as: Comp } = props;

  if (!hasRequiredRole(authContext.user.roles, props.allowedRoles)) {
    navigate("/login");
  }

  return <Comp {...props} />;
};

export default PrivateRoute;
