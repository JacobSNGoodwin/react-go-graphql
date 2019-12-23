import React, { useState } from "react";

import { useMutation } from "@apollo/react-hooks";
import Cookies from "js-cookie";
import { LOGIN_GOOGLE, LOGIN_FACEBOOK } from "../../gql/mutations";

interface IAuthContext {
  user: IUser | undefined;
  loading: boolean;
  errors: IError[];
  loginWithGoogle(token: string): void;
  loginWithFacebook(token: string): void;
  logout(): void;
}

interface IError {
  message: string;
  type: string | undefined;
}

interface IUser {
  id: string;
  name?: string;
  email: string;
  imageUri?: string;
  roles: string[];
}

// maybe add error if login/logout aren't defined in Auth provided
const defaultAuth: IAuthContext = {
  user: undefined,
  loading: false,
  errors: [],
  loginWithGoogle: () => {}, // produce error if not overwritten in Provider?
  loginWithFacebook: () => {},
  logout: () => {}
};

const AuthContext = React.createContext<IAuthContext>(defaultAuth);

const AuthProvider: React.FC = props => {
  // get userID from userinfo cookie which holds the id of authorized user
  const userCookie = Cookies.get("userinfo");

  console.log("The user cookie: ", userCookie);

  // useState for these properties
  const [user, setUser] = useState<IUser | undefined>(undefined);
  const [errors, setErrors] = useState<IError[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [loginGoogleMutation] = useMutation<
    { googleLoginWithToken: IUser },
    { idToken: string }
  >(LOGIN_GOOGLE, {
    errorPolicy: "ignore",
    onCompleted: ({ googleLoginWithToken }) => {
      setUser(googleLoginWithToken);
      setLoading(false);
    },
    onError: error => {
      const errors = error.graphQLErrors.map(error => {
        const type = error.extensions ? error.extensions.type : undefined;
        return {
          message: error.message,
          type: type
        };
      });
      setErrors(errors);
      setLoading(false);
    }
  });

  const [loginFacebookMutation] = useMutation<
    { fbLoginWithToken: IUser },
    { accessToken: string }
  >(LOGIN_FACEBOOK, {
    errorPolicy: "ignore",
    onCompleted: ({ fbLoginWithToken }) => {
      setUser(fbLoginWithToken);
      setLoading(false);
    },
    onError: error => {
      const errors = error.graphQLErrors.map(error => {
        const type = error.extensions ? error.extensions.type : undefined;
        return {
          message: error.message,
          type: type
        };
      });
      setErrors(errors);
      setLoading(false);
    }
  });

  // Add login functions (for setting state here)
  const loginWithGoogle = (token: string) => {
    setLoading(true);

    loginGoogleMutation({
      variables: {
        idToken: token
      }
    });
  };

  const loginWithFacebook = (token: string) => {
    setLoading(true);

    loginFacebookMutation({
      variables: {
        accessToken: token
      }
    });
  };

  const logout = () => {
    console.log("logging out");
  };

  return (
    <AuthContext.Provider
      value={{
        user,
        loading,
        errors,
        loginWithGoogle,
        loginWithFacebook,
        logout
      }}
      {...props}
    >
      {props.children}
    </AuthContext.Provider>
  );
};

export { AuthProvider, AuthContext };
