import React, { useState } from "react";

interface IAuthContext {
  authenticated: boolean;
  loading: boolean;
  roles: string[];
  login(): void;
  logout(): void;
}

// maybe add error if login/logout aren't defined in Auth provided
const defaultAuth: IAuthContext = {
  authenticated: false,
  loading: false,
  roles: [],
  login: () => {},
  logout: () => {}
};

const AuthContext = React.createContext<IAuthContext>(defaultAuth);

const AuthProvider: React.FC = props => {
  // useState for these properties
  const [authenticated, setAuthenticated] = useState<boolean>(false);
  const [roles, setRoles] = useState<string[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  // Fetch user from jwt cookie that is js accessible

  // Add login functions (for setting state here)
  const login = () => {
    setLoading(true);
    setTimeout(() => {
      setAuthenticated(true);
      setRoles(["admin", "editor"]);
      setLoading(false);
    }, 1000);
  };

  const logout = () => {
    setLoading(true);
    setTimeout(() => {
      setAuthenticated(false);
      setRoles([]);
      setLoading(false);
    }, 1000);
  };

  return (
    <AuthContext.Provider
      value={{
        authenticated,
        loading,
        roles,
        login,
        logout
      }}
      {...props}
    >
      {props.children}
    </AuthContext.Provider>
  );
};

export { AuthProvider, AuthContext };
