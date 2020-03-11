import React from "react";

import { RouteComponentProps, navigate } from "@reach/router";
import { AuthContext } from "./contexts/AuthContext";

interface IPrivateRouteProps extends RouteComponentProps {
  as: React.FC;
  admin?: boolean;
  editor?: boolean;
}

const PrivateRoute: React.FC<IPrivateRouteProps> = props => {
  const { user } = React.useContext(AuthContext);

  const { as: Comp } = props;

  let allowed: boolean = false;

  if (props.admin && user && user.roles.admin) {
    allowed = true;
  }

  if (props.editor && user && user.roles.editor) {
    allowed = true;
  }

  if (user && !allowed) {
    const errorProps: ErrorProps = {
      messages: ["You are not allowed to access this resrouce"],
      includeLogin: true
    };
    navigate("error", { state: errorProps });
  }

  return <Comp {...props} />;
};

export default PrivateRoute;
