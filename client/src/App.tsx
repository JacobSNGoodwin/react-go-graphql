import React from "react";
import { Router } from "@reach/router";
import { ApolloProvider } from "@apollo/react-hooks";
import ApolloClient from "apollo-boost";

import { AuthProvider } from "./components/contexts/AuthContext";
import Login from "./components/Login";
import Navbar from "./components/Navbar";
import "./App.scss";

const client = new ApolloClient({
  uri: process.env.REACT_APP_URI_GQL,
  credentials: "include",
  headers: {
    "X-Requested-With": "XMLHttpRequest"
  }
});

const App: React.FC = () => {
  return (
    <ApolloProvider client={client}>
      <AuthProvider>
        <Navbar />
        <div className="container has-text-centered">
          <Router>
            <Login path="login" />
          </Router>
        </div>
      </AuthProvider>
    </ApolloProvider>
  );
};

export default App;
