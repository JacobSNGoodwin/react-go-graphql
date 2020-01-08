import React from "react";

import { RouteComponentProps } from "@reach/router";
import { AuthContext } from "../components/contexts/AuthContext";
import Error from "./Error";

interface IPrivateRouteProps {
  as: React.FC;
  admin?: boolean;
  editor?: boolean;
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

  let allowed: boolean = false;

  if (props.admin && user.roles.admin) {
    allowed = true;
  }

  if (props.editor && user.roles.editor) {
    allowed = true;
  }

  if (user && !allowed) {
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
