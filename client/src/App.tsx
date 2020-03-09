import React from "react";
import { Router } from "@reach/router";
import { ApolloProvider } from "@apollo/react-hooks";
import ApolloClient from "apollo-boost";

import { AuthProvider } from "./components/contexts/AuthContext";
import Navbar from "./components/Navbar";
import Home from "./components/pages/Home";
import Login from "./components/pages/Login";
import UsersAdmin from "./components/pages/UsersAdmin";
import ProductsAdmin from "./components/pages/ProductsAdmin";
import NotFound from "./components/pages/NotFound";
import Error from "./components/pages/Error";
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
        <>
          <Navbar />
          <div className="section">
            <Router>
              <Home path="/" />
              <Login path="login" />
              <PrivateRoute as={UsersAdmin} admin path="users-admin" />
              <PrivateRoute
                as={ProductsAdmin}
                admin
                editor
                path="products-admin"
              />
              {/* <PrivateRoute
                as={CategoriesAdmin}
                admin
                editor
                path="categories-admin"
              /> */}
              <Error path="error" />
              <NotFound default />
            </Router>
          </div>
        </>
      </AuthProvider>
    </ApolloProvider>
  );
};

export default App;
