import React from "react";
import { useQuery, useMutation } from "@apollo/react-hooks";

import User from "./User";
import EditUser from "./EditUser";
import { GET_USERS } from "../gql/queries";
import { CREATE_USER, EDIT_USER, DELETE_USER } from "../gql/mutations";
import Spinner from "./ui/Spinner";
import { transformUserFromGQL } from "../util/util";
import ErrorsList from "./ErrorsList";

const Users: React.FC = props => {
  const [createActive, setCreateActive] = React.useState<boolean>(false);
  /*
   * Create User
   */
  const [
    createUserMutation,
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
  const { loading: loadingUsers, error: usersError, data } = useQuery<
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

  const [
    editUserMutation,
    { loading: editingUser, error: editError }
  ] = useMutation<{ editedUser: IUserGQL }, { user: IUserGQL }>(EDIT_USER, {
    errorPolicy: "ignore"
  });

  /*
   * Delete User
   */
  const [
    deleteUserMutation,
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

  // functions to wrap mutations to pass to edit/delete components
  const editUser = (user: IUserGQL) => {
    return editUserMutation({
      variables: {
        user
      }
    });
  };

  const deleteUser = (id: string) => {
    return deleteUserMutation({
      variables: {
        id
      }
    });
  };

  /*
   * Rendering
   */
  if (loadingUsers)
    return (
      <div className="container">
        <div className="columns is-centered">
          <Spinner radius={50} />
        </div>
      </div>
    );

  // manage all errors in state and with useEffect
  if (usersError) {
    return <ErrorsList error={usersError} />;
  }

  if (editError) {
    return <ErrorsList error={editError} />;
  }

  if (createError) {
    return <ErrorsList error={createError} />;
  }

  if (deleteError) {
    return <ErrorsList error={deleteError} />;
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
          createUserMutation({
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
