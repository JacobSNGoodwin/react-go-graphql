import React from "react";
import { Router } from "@reach/router";
import Login from "./components/Login";
import Navbar from "./components/Navbar";
import "./App.scss";

const App: React.FC = () => {
  return (
    <>
      <Navbar />
      <h1 className="is-size-3">Hello from App!</h1>

      <div className="container has-text-centered">
        <Router>
          <Login path="login" />
        </Router>
      </div>
    </>
  );
};

export default App;
