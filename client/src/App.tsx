import React from "react";
import { Router } from "@reach/router";
import Login from "./components/Login";
import Navbar from "./components/Navbar";
import { AuthProvider } from "./components/contexts/AuthContext";
import "./App.scss";

const App: React.FC = () => {
  return (
    <AuthProvider>
      <Navbar />
      <div className="container has-text-centered">
        <Router>
          <Login path="login" />
        </Router>
      </div>
    </AuthProvider>
  );
};

export default App;
