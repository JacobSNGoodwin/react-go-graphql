import React from "react";
import { Link } from "@reach/router";
import { AuthContext } from "./contexts/AuthContext";

const Navbar: React.FC = props => {
  const authContext = React.useContext(AuthContext);
  return (
    <nav className="navbar is-fixed-top">
      <div className="navbar-brand">
        <Link to="/" className="navbar-item">
          <img
            src={`${process.env.PUBLIC_URL}/android-chrome-192x192.png`}
            alt="Home Link Logo"
          />
        </Link>
        <button className="navbar-burger burger">
          <span></span>
          <span></span>
          <span></span>
        </button>
      </div>
      <div className="navbar-menu">
        <div className="navbar-start">
          <Link to="/users" className="navbar-item">
            Users
          </Link>
        </div>
        <div className="navbar-end">
          <div className="navbar-item">
            <Link to="/login" className="button is-link">
              Sign In
            </Link>
          </div>
          <div className="navbar-item">
            <button onClick={authContext.logout} className="button is-link">
              Sign Out
            </button>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
