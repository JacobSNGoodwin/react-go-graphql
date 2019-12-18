import React from "react";
import { Router } from "@reach/router";
import Login from "./components/Login";
import Navbar from "./components/Navbar";
import "./App.scss";

const App: React.FC = () => {
  return (
    <>
      <Navbar />
      <div className="container has-text-centered">
        <Router>
          <Login path="login" />
        </Router>
      </div>
    </>
  );
};

export default App;
