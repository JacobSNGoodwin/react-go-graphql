import React from "react";

interface IAuthContext {
  authenticated: boolean;
  loading: boolean;
  roles: Array<string>;
  login(token: string): void;
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
  let authenticated = false;
  let roles: string[] = [];
  let loading: boolean = false;

  // Fetch user from jwt cookie that is js accessible

  // Add login functions (for setting state here)
  const login = (token: string) => {
    if (token) {
      loading = true;
      // timeout just to show a spinner or something
      setTimeout(() => {}, 1000);
      authenticated = true;
      loading = false;
      roles.push("admin", "editor");
    }
  };

  // Add login functions (for setting state here)
  const logout = () => {
    loading = true;
    setTimeout(() => {}, 1000);
    authenticated = true;
    loading = false;
    roles = [];
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
    />
  );
};

// Example returns a function that returns the context
const useAuth = React.useContext(AuthContext);

export { AuthProvider, useAuth };
