import React, { useState, useEffect } from "react";
import { useMutation, useLazyQuery } from "@apollo/react-hooks";
import { navigate } from "@reach/router";
import Cookies from "js-cookie";

import { LOGIN_GOOGLE, LOGIN_FACEBOOK } from "../../gql/mutations";
import { ME } from "../../gql/queries";
import Spinner from "../ui/Spinner";
import { transformUserFromGQL } from "../../util/util";

interface IAuthContext {
  user: IUser | undefined;
  loading: boolean;
  errors: IError[];
  loginWithGoogle(token: string): void;
  loginWithFacebook(token: string): void;
  logout(): void;
}

// maybe add error if login/logout aren't defined in Auth provided
const defaultAuth: IAuthContext = {
  user: undefined,
  loading: true,
  errors: [],
  loginWithGoogle: () => {}, // produce error if not overwritten in Provider?
  loginWithFacebook: () => {},
  logout: () => {}
};

const AuthContext = React.createContext<IAuthContext>(defaultAuth);

const AuthProvider: React.FC = props => {
  // useState for these properties
  const [user, setUser] = useState<IUser | undefined>(defaultAuth.user);
  const [userCookie, setUserCookie] = useState<string | undefined>(
    Cookies.get("userinfo")
  );
  const [errors, setErrors] = useState<IError[]>(defaultAuth.errors);
  const [loading, setLoading] = useState<boolean>(defaultAuth.loading);
  const [loginGoogleMutation] = useMutation<
    { googleLoginWithToken: IUserGQL },
    { idToken: string }
  >(LOGIN_GOOGLE, {
    errorPolicy: "ignore",
    onCompleted: ({ googleLoginWithToken }) => {
      setUserCookie(Cookies.get("userinfo"));
      setUser(transformUserFromGQL(googleLoginWithToken));
      setLoading(false);
      navigate("/");
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
    { fbLoginWithToken: IUserGQL },
    { accessToken: string }
  >(LOGIN_FACEBOOK, {
    errorPolicy: "ignore",
    onCompleted: ({ fbLoginWithToken }) => {
      setUserCookie(Cookies.get("userinfo"));
      setUser(transformUserFromGQL(fbLoginWithToken));
      setLoading(false);
      navigate("/");
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
  const [getMe] = useLazyQuery<{ me: IUserGQL }>(ME, {
    errorPolicy: "none",
    onCompleted: ({ me }) => {
      setUserCookie(Cookies.get("userinfo"));
      setUser(transformUserFromGQL(me));
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
    if (userCookie && !user) {
      // only attempt to get user if a cookie is present
      setLoading(true);
      getMe();
    } else if (!userCookie) {
      setUser(undefined);
      setLoading(false);
    } else {
      setLoading(false);
    }
  }, [getMe, user, userCookie]);

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
    setLoading(true);
    navigate("/login");
    Cookies.remove("userinfo");
    setUserCookie(undefined);
    setUser(undefined);
  };

  if (loading) {
    return (
      <div className="section">
        <div className="container">
          <div className="columns is-centered">
            <Spinner />
          </div>
        </div>
      </div>
    );
  }

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
