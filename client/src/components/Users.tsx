import React from "react";
import { useQuery } from "@apollo/react-hooks";

import Spinner from "./ui/Spinner";
import { GET_USERS } from "../gql/queries";
import User from "./User";
import { transformUserFromGQL } from "../util/util";

interface IUserData {
  users: IUserGQL[];
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
      <div className="container">
        <div className="columns is-centered">
          <Spinner radius={50} />
        </div>
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

  const userList =
    data &&
    data.users.map(userGQL => {
      const user = transformUserFromGQL(userGQL);
      return (
        <div key={user.id} className="column is-half">
          <User user={user} />
        </div>
      );
    });

  return (
    <div className="section">
      <h1 className="title is-1 has-text-centered">Users</h1>
      <div className="container">
        <div className="columns is-vcentered is-multiline">{userList}</div>
      </div>
    </div>
  );
};

export default Users;
