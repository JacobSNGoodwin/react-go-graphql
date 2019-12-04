import React from "react";
import "./App.scss";
import {
  GoogleLogin,
  GoogleLoginResponse,
  GoogleLoginResponseOffline
} from "react-google-login";
import FacebookLogin, { ReactFacebookLoginInfo } from "react-facebook-login";

const App: React.FC = () => {
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
        buttonText="Login"
        onSuccess={responseGoogle}
        onFailure={responseGoogle}
        cookiePolicy={"single_host_origin"}
      />
      <FacebookLogin
        appId={fbClientid}
        textButton="Login"
        fields="name,email,picture"
        callback={responseFacebook}
        icon="fa-facebook"
      />
    </div>
  );
};

export default App;
