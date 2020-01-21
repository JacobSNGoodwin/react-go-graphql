import React, { useState, useEffect } from "react";
import { useMutation, useQuery } from "@apollo/react-hooks";
import { navigate, redirectTo } from "@reach/router";
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
  const [errors, setErrors] = useState<IError[]>(defaultAuth.errors);

  const [loginGoogleMutation, { loading: googleLoading }] = useMutation<
    { googleLoginWithToken: IUserGQL },
    { idToken: string }
  >(LOGIN_GOOGLE, {
    errorPolicy: "ignore",
    onCompleted: ({ googleLoginWithToken }) => {
      setUser(transformUserFromGQL(googleLoginWithToken));
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
    }
  });

  const [loginFacebookMutation, { loading: facebookLoading }] = useMutation<
    { fbLoginWithToken: IUserGQL },
    { accessToken: string }
  >(LOGIN_FACEBOOK, {
    errorPolicy: "ignore",
    onCompleted: ({ fbLoginWithToken }) => {
      setUser(transformUserFromGQL(fbLoginWithToken));
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
    }
  });

  // lazy query to fetch me only on initial load
  // in the future we could do this more frequently
  const { loading: meLoading, refetch: refetchMe } = useQuery<{ me: IUserGQL }>(
    ME,
    {
      errorPolicy: "none",
      onCompleted: ({ me }) => {
        setUser(transformUserFromGQL(me));
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
      }
    }
  );

  useEffect(() => {
    async function getMe() {
      const resp = await refetchMe();
      if (resp.data) {
        setUser(transformUserFromGQL(resp.data.me));
      }
    }

    const userCookie = Cookies.get("userinfo");

    if (userCookie && !user) {
      getMe();
    }
  });

  // Add login functions (for setting state here)
  const loginWithGoogle = (token: string) => {
    loginGoogleMutation({
      variables: {
        idToken: token
      }
    });
  };

  const loginWithFacebook = (token: string) => {
    loginFacebookMutation({
      variables: {
        accessToken: token
      }
    });
  };

  const logout = () => {
    Cookies.remove("userinfo");
    setUser(undefined);
    redirectTo("/login");
  };

  if (googleLoading || facebookLoading || meLoading) {
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
        loading: googleLoading || facebookLoading || meLoading,
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
