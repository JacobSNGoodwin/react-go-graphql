import React from "react";

import { navigate, RouteComponentProps } from "@reach/router";
import { AuthContext } from "../components/contexts/AuthContext";
import { hasRequiredRole } from "../util/util";

interface IPrivateRouteProps {
  as: React.FC;
  allowedRoles: string[];
}

const PrivateRoute: React.FC<IPrivateRouteProps &
  RouteComponentProps> = props => {
  const { user } = React.useContext(AuthContext);

  if (!user) {
    navigate("/login");
  }

  let { as: Comp } = props;

  if (user && hasRequiredRole(user.roles, props.allowedRoles)) {
    navigate("/login");
  }

  return <Comp {...props} />;
};

export default PrivateRoute;
