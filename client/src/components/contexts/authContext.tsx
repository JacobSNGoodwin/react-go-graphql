import React, { useState, useEffect } from "react";
import { useMutation, useLazyQuery } from "@apollo/react-hooks";
import Cookies from "js-cookie";

import { LOGIN_GOOGLE, LOGIN_FACEBOOK } from "../../gql/mutations";
import { ME } from "../../gql/queries";

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

  // lazy query to fetch me only on initial load
  // in the future we could do this more frequently
  const [getMe] = useLazyQuery<{ me: IUser }>(ME, {
    errorPolicy: "none",
    onCompleted: ({ me }) => {
      setUser(me);
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

  // get useriD from cookie on initial load
  // attempt to load user from me
  useEffect(() => {
    setLoading(true);
    if (Cookies.get("userinfo")) {
      // only attempt to get user if a cookie is present
      getMe();
    } else {
      setLoading(false);
    }
  }, [getMe]);

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
    Cookies.remove("userinfo");
    setUser(undefined);
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
