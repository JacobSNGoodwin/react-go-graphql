import React from "react";
import { useQuery } from "@apollo/react-hooks";

import Spinner from "./ui/Spinner";
import { GET_USERS } from "../gql/queries";
import User from "./User";

// Maybe store this in common types later
interface IUser {
  id: string;
  name?: string;
  email: string;
  imageUri?: string;
  roles: string[];
}

interface IUserData {
  users: IUser[];
}

interface IUserVars {
  limit?: number;
  offset?: number;
}

const Users: React.FC = props => {
  const { loading, error, data } = useQuery<IUserData, IUserVars>(GET_USERS, {
    variables: {
      limit: 10
    }
  });

  if (loading)
    return (
      <div className="container has-text-centered">
        <Spinner radius={50} />;
      </div>
    );

  if (error) {
    return (
      <div className="container has-text-centered">
        <h1 className="title is-4">Error</h1>
        {error.graphQLErrors.map((error, i) => {
          return <p key={i}>{error.message}</p>;
        })}
      </div>
    );
  }

  const userList = data && data.users.map(user => <User user={user} />);

  return (
    <div className="section">
      <h1 className="title is-1 has-text-centered">Users</h1>
      {userList}
    </div>
  );
};

export default Users;
