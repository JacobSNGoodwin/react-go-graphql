import React from "react";
import { useQuery, useMutation } from "@apollo/react-hooks";

import User from "./User";
import { GET_USERS } from "../gql/queries";
import { DELETE_USER } from "../gql/mutations";
import Spinner from "./ui/Spinner";
import { transformUserFromGQL } from "../util/util";

const Users: React.FC = props => {
  const { loading: loadingUsers, error: errorUsers, data } = useQuery<
    IUserData,
    IUserVars
  >(GET_USERS, {
    variables: {
      limit: 10
    }
  });

  const [deleteUser, { loading: deletingUser }] = useMutation<
    { deleteUser: string },
    { id: string }
  >(DELETE_USER, {
    onCompleted: () => {
      console.log("user deleted");
    },
    onError: error => {
      const errors = error.graphQLErrors.map(error => {
        const type = error.extensions ? error.extensions.type : undefined;
        return {
          message: error.message,
          type: type
        };
      });
      console.log(errors);
    },
    refetchQueries: [
      {
        query: GET_USERS,
        variables: {
          limit: 10
        }
      }
    ]
  });

  if (loadingUsers)
    return (
      <div className="container">
        <div className="columns is-centered">
          <Spinner radius={50} />
        </div>
      </div>
    );

  if (errorUsers) {
    return (
      <div className="container has-text-centered">
        <h1 className="title is-4">Error</h1>
        {errorUsers.graphQLErrors.map((error, i) => {
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
          <User
            user={user}
            deletingUser={deletingUser}
            deleteUser={() => {
              deleteUser({ variables: { id: user.id } });
            }}
          />
        </div>
      );
    });

  return (
    <>
      <h1 className="title is-1 has-text-centered">Users</h1>
      <div className="container">
        <div className="columns is-vcentered is-multiline">{userList}</div>
      </div>
    </>
  );
};

export default Users;
