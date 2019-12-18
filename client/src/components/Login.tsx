import React from "react";

import {
  GoogleLogin,
  GoogleLoginResponse,
  GoogleLoginResponseOffline
} from "react-google-login";

import FacebookLogin, { ReactFacebookLoginInfo } from "react-facebook-login";
import FacebookIcon from "./icons/Facebook";
import GoogleIcon from "./icons/Google";

const login: React.FC = () => {
  const responseGoogle = (
    res: GoogleLoginResponse | GoogleLoginResponseOffline
  ) => {
    if ((res as GoogleLoginResponse).getAuthResponse) {
      console.log((res as GoogleLoginResponse).getAuthResponse().id_token);
    }
  };

  const responseFacebook = (res: ReactFacebookLoginInfo) => {
    console.log(res);
  };

  const googleClientid: string = process.env.REACT_APP_GOOGLE_CLIENT_ID
    ? process.env.REACT_APP_GOOGLE_CLIENT_ID
    : "";
  const fbClientid: string = process.env.REACT_APP_FACEBOOK_CLIENT_ID
    ? process.env.REACT_APP_FACEBOOK_CLIENT_ID
    : "";

  return (
    <div className="App">
      <GoogleLogin
        clientId={googleClientid}
        onSuccess={responseGoogle}
        onFailure={responseGoogle}
        render={renderProps => (
          <button
            className={`button is-large`}
            onClick={renderProps.onClick}
            disabled={renderProps.disabled}
          >
            <span className="icon is-large">
              <GoogleIcon width="30px" height="30px" />
            </span>
            <span>Sign In</span>
          </button>
        )}
      />
      <FacebookLogin
        appId={fbClientid}
        textButton="Sign In"
        typeButton="medium"
        fields="name,email,picture"
        callback={responseFacebook}
        cssClass={`button is-large`}
        icon={<FacebookIcon width="36px" height="36px" />}
      />
    </div>
  );
};

export default login;
