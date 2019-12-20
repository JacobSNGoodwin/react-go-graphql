import React from "react";

import Spinner from "./ui/Spinner";
import styles from "./Login.module.scss";

import {
  GoogleLogin,
  GoogleLoginResponse,
  GoogleLoginResponseOffline
} from "react-google-login";

import { RouteComponentProps } from "@reach/router";

import { AuthContext } from "./contexts/AuthContext";

import FacebookLogin, { ReactFacebookLoginInfo } from "react-facebook-login";
import FacebookIcon from "./icons/Facebook";
import GoogleIcon from "./icons/Google";

const Login: React.FC<RouteComponentProps> = props => {
  const authContext = React.useContext(AuthContext);

  const responseGoogle = (
    res: GoogleLoginResponse | GoogleLoginResponseOffline
  ) => {
    if ((res as GoogleLoginResponse).getAuthResponse) {
      const token = (res as GoogleLoginResponse).getAuthResponse().id_token;
      authContext.loginWithGoogle(token);
    }
  };

  const responseFacebook = (res: ReactFacebookLoginInfo) => {
    console.log(res);
    authContext.logout();
  };

  const googleClientid: string = process.env.REACT_APP_GOOGLE_CLIENT_ID
    ? process.env.REACT_APP_GOOGLE_CLIENT_ID
    : "";
  const fbClientid: string = process.env.REACT_APP_FACEBOOK_CLIENT_ID
    ? process.env.REACT_APP_FACEBOOK_CLIENT_ID
    : "";

  return (
    <>
      <div className="section">
        <div className="buttons is-centered">
          <GoogleLogin
            clientId={googleClientid}
            onSuccess={responseGoogle}
            onFailure={responseGoogle}
            render={renderProps => (
              <button
                className={`${styles.button} button is-large`}
                onClick={renderProps.onClick}
                disabled={renderProps.disabled}
              >
                <GoogleIcon width="36px" height="36px" />
                Sign In
              </button>
            )}
          />

          <FacebookLogin
            appId={fbClientid}
            textButton="Sign In"
            fields="name,email,picture"
            callback={responseFacebook}
            cssClass={`${styles.button} button is-large`}
            icon={<FacebookIcon width="36px" height="36px" />}
          />
        </div>

        <div className="columns is-centered">
          {authContext.loading && <Spinner radius={40} />}
        </div>

        <div className="content">
          <p>
            User authenticated? - {authContext.authenticated ? "true" : "false"}
          </p>
          <h3>User Roles</h3>
          {authContext.roles.map(role => (
            <ul>{role}</ul>
          ))}
        </div>
      </div>
    </>
  );
};

export default Login;
