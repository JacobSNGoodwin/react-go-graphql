import React from "react";
import { ExecutionResult, MutationFunctionOptions } from "@apollo/react-common";

import placeholder from "../images/placeholder.png";
import DeleteUser from "./DeleteUser";
import EditUser from "./EditUser";

import styles from "./User.module.scss";
import { ApolloError } from "apollo-boost";

interface UserProps {
  user: IUser;
  editingUser: boolean;
  editUser: (
    options?:
      | MutationFunctionOptions<
          {
            editedUser: IUserGQL;
          },
          {
            user: IUserGQL;
          }
        >
      | undefined
  ) => Promise<
    ExecutionResult<{
      editedUser: IUserGQL;
    }>
  >;
  editError: ApolloError | undefined;
  deletingUser: boolean;
  deleteUser: (
    options?:
      | MutationFunctionOptions<
          {
            deleteUser: string;
          },
          {
            id: string;
          }
        >
      | undefined
  ) => Promise<
    ExecutionResult<{
      deleteUser: string;
    }>
  >;
  deleteError: ApolloError | undefined;
}

const User: React.FC<UserProps> = props => {
  const [editActive, setEditActive] = React.useState<boolean>(false);
  const [deleteActive, setDeleteActive] = React.useState<boolean>(false);
  const userRoles: string[] = [];

  if (props.user.roles.admin) {
    userRoles.push("Admin");
  }

  if (props.user.roles.editor) {
    userRoles.push("Editor");
  }

  const editMutation = (gqlUser: IUserGQL) => {
    props
      .editUser({
        variables: {
          user: gqlUser
        }
      })
      .then(result => {
        if (result.data) {
          setEditActive(false);
        }
      });
  };

  const deleteMutation = () => {
    props
      .deleteUser({
        variables: {
          id: props.user.id
        }
      })
      .then(result => {
        if (result.data) {
          setDeleteActive(false);
        }
      });
  };

  return (
    <div className="card" key={props.user.id}>
      <div className="card-content has-text-centered">
        <div
          style={{
            display: "flex",
            justifyContent: "center",
            padding: "1.5em 1em"
          }}
        >
          <figure className="image is-96x96">
            {!props.user.imageUri || props.user.imageUri === "" ? (
              <img className="is-rounded" src={placeholder} alt="No profile" />
            ) : (
              <img
                className="is-rounded"
                src={props.user.imageUri}
                alt="User profile"
              />
            )}
          </figure>
        </div>

        <p className="title is-4">{props.user.name}</p>
        <p className="subtitle is-6">{props.user.email}</p>
        <br />
        <div className="is-size-5 has-text-weight-bold">Roles</div>
        <div>{userRoles.length > 0 ? userRoles.join(", ") : "None"}</div>

        <div className={styles.userButtons}>
          <button
            onClick={() => setEditActive(true)}
            className="button is-primary"
          >
            Edit
          </button>
          <button
            onClick={() => setDeleteActive(true)}
            className="button is-link"
          >
            Delete
          </button>
        </div>

        <EditUser
          show={editActive}
          editSelectedUser={editMutation}
          editingUser={props.editingUser}
          close={() => setEditActive(false)}
          initUser={props.user}
          error={props.editError}
        />

        <DeleteUser
          show={deleteActive}
          deleteSelectedUser={deleteMutation}
          deletingUser={props.deletingUser}
          close={() => setDeleteActive(false)}
          user={props.user}
        />
      </div>
    </div>
  );
};

export default User;
