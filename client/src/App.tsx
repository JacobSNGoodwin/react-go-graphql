import React from "react";
import { Router } from "@reach/router";
import { ApolloProvider } from "@apollo/react-hooks";
import ApolloClient from "apollo-boost";

import { AuthProvider } from "./components/contexts/AuthContext";
import Navbar from "./components/Navbar";
import Home from "./components/Home";
import Login from "./components/Login";
import Users from "./components/Users";
import PrivateRoute from "./components/PrivateRoute";
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
        <div className="section">
          <Router>
            <Home path="/" />
            <Login path="login" />
            <PrivateRoute as={Users} allowedRoles={["Admin"]} path="users" />
          </Router>
        </div>
      </AuthProvider>
    </ApolloProvider>
  );
};

export default App;
