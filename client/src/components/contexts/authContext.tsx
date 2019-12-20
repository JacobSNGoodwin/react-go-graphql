import React, { useState } from "react";

import { useMutation } from "@apollo/react-hooks";
import { gql } from "apollo-boost";

interface IAuthContext {
  authenticated: boolean;
  loading: boolean;
  roles: string[];
  loginWithGoogle(token: string): void;
  logout(): void;
}

// maybe add error if login/logout aren't defined in Auth provided
const defaultAuth: IAuthContext = {
  authenticated: false,
  loading: false,
  roles: [],
  loginWithGoogle: () => {}, // produce error if not overwritten in Provider?
  logout: () => {}
};

const LOGIN_GOOGLE = gql`
  mutation LoginGoogle($idToken: String!) {
    googleLoginWithToken(idToken: $idToken) {
      id
      name
      email
      imageUri
      roles
    }
  }
`;

const AuthContext = React.createContext<IAuthContext>(defaultAuth);

const AuthProvider: React.FC = props => {
  // useState for these properties
  const [authenticated, setAuthenticated] = useState<boolean>(false);
  const [roles, setRoles] = useState<string[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [
    loginGoogleMutation,
    {
      data: googleData,
      loading: googleLoading,
      error: googleError,
      called: googleCalled
    }
  ] = useMutation(LOGIN_GOOGLE);

  // Fetch user from jwt cookie that is js accessible

  // Add login functions (for setting state here)
  const loginWithGoogle = (token: string) => {
    // will use callback here since we maintain loading state for the auth provider in this component
    loginGoogleMutation({
      variables: {
        idToken: token
      }
    });

    if (googleLoading) {
      setLoading(true);
    }

    if (googleCalled && !googleLoading) {
      setLoading(false);
      if (googleData) {
        console.log(googleData);
      }
      if (googleError) {
        console.log(googleError);
      }
    }
  };

  const logout = () => {
    setLoading(true);
    setTimeout(() => {
      setAuthenticated(false);
      setRoles([]);
      setLoading(false);
    }, 1000);
  };

  return (
    <AuthContext.Provider
      value={{
        authenticated,
        loading,
        roles,
        loginWithGoogle,
        logout
      }}
      {...props}
    >
      {props.children}
    </AuthContext.Provider>
  );
};

export { AuthProvider, AuthContext };
