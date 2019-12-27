import React from "react";

import { RouteComponentProps } from "@reach/router";
import { AuthContext } from "../components/contexts/AuthContext";
import Error from "./Error";

import { hasRequiredRole } from "../util/util";

interface IPrivateRouteProps {
  as: React.FC;
  allowedRoles: string[];
}

const PrivateRoute: React.FC<IPrivateRouteProps &
  RouteComponentProps> = props => {
  const { user } = React.useContext(AuthContext);

  if (!user) {
    return (
      <Error
        messages={[
          "User needs to be logged in with proper permissions to access this resource"
        ]}
        includeLogin
      />
    );
  }

  let { as: Comp } = props;

  if (user && !hasRequiredRole(user.roles, props.allowedRoles)) {
    return (
      <Error
        messages={["User not permitted to view this resource"]}
        includeLogin
        includeHomeButton
      />
    );
  }

  return <Comp {...props} />;
};

export default PrivateRoute;
