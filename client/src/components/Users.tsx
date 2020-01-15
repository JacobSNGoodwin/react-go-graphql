import React from "react";
import { useQuery, useMutation } from "@apollo/react-hooks";

import User from "./User";
import EditUser from "./EditUser";
import { GET_USERS } from "../gql/queries";
import { CREATE_USER, EDIT_USER, DELETE_USER } from "../gql/mutations";
import Spinner from "./ui/Spinner";
import { transformUserFromGQL } from "../util/util";

const Users: React.FC = props => {
  const [createActive, setCreateActive] = React.useState<boolean>(false);
  /*
   * Create User
   */
  const [
    createUser,
    { loading: creatingUser, error: createError }
  ] = useMutation<{ createdUser: IUserGQL }, { user: IUserGQL }>(CREATE_USER, {
    refetchQueries: [
      {
        query: GET_USERS,
        variables: {
          limit: 10
        }
      }
    ]
  });

  /*
   * Read (get) Users
   */
  const { loading: loadingUsers, error: errorUsers, data } = useQuery<
    IUserData,
    IUserVars
  >(GET_USERS, {
    variables: {
      limit: 10
    }
  });

  /*
   * Update Users
   */

  const [editUser, { loading: editingUser, error: editError }] = useMutation<
    { editedUser: IUserGQL },
    { user: IUserGQL }
  >(EDIT_USER);

  /*
   * Delete User
   */
  const [
    deleteUser,
    { loading: deletingUser, error: deleteError }
  ] = useMutation<{ deleteUser: string }, { id: string }>(DELETE_USER, {
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
            editingUser={editingUser}
            editUser={editUser}
            deletingUser={deletingUser}
            deleteUser={deleteUser}
          />
        </div>
      );
    });

  return (
    <>
      <h1 className="title is-1 has-text-centered">Users</h1>
      <div className="container">
        <div className="columns is-centered">
          <button
            className="button is-warning"
            onClick={() => {
              setCreateActive(true);
            }}
          >
            Create
          </button>
        </div>
        <div className="columns is-vcentered is-multiline">{userList}</div>
      </div>

      <EditUser
        show={createActive}
        editSelectedUser={gqlUser => {
          createUser({
            variables: {
              user: gqlUser
            }
          }).then(() => {
            setCreateActive(false);
          });
        }}
        editingUser={creatingUser}
        close={() => setCreateActive(false)}
      />
    </>
  );
};

export default Users;
