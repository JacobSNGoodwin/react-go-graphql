import React from "react";
import { Link } from "@reach/router";

const navbar: React.FC = props => (
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
            Log in
          </Link>
        </div>
      </div>
    </div>
  </nav>
);

export default navbar;
