import React from "react";
import Login from "./components/Login";

import "./App.scss";

const App: React.FC = () => {
  return (
    <div className="container has-text-centered">
      <h1>Hello from App!</h1>
      <Login />
    </div>
  );
};

export default App;
