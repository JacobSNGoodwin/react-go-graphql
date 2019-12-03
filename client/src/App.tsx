import React from "react";
import "./App.css";
import {
  GoogleLogin,
  GoogleLoginResponse,
  GoogleLoginResponseOffline
} from "react-google-login";

const App: React.FC = () => {
  const responseGoogle = (
    res: GoogleLoginResponse | GoogleLoginResponseOffline
  ) => {
    if ((res as GoogleLoginResponse).getAuthResponse) {
      console.log((res as GoogleLoginResponse).getAuthResponse().id_token);
    }
  };

  return (
    <div className="App">
      <GoogleLogin
        clientId="23806398157-4i36lv7vhdh3k3j5gnspggm75qg83o5v.apps.googleusercontent.com"
        buttonText="Login"
        onSuccess={responseGoogle}
        onFailure={responseGoogle}
        cookiePolicy={"single_host_origin"}
      />
    </div>
  );
};

export default App;
