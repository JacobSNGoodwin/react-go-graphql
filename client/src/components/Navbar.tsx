import React from "react";
import { Link } from "@reach/router";

const navbar: React.FC = props => (
  <nav className="navbar is-dark">
    <div className="navbar-menu">
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
