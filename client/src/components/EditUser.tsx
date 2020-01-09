import React from "react";
import { useForm } from "react-hook-form";

import {
  emailPattern,
  imageUrlPattern,
  transformUserToGQL
} from "../util/util";
import { EDIT_USER } from "../gql/mutations";
import { useMutation } from "@apollo/react-hooks";

interface EditUserProps {
  show: boolean;
  close: () => void;
  initUser?: IUser;
}

const EditUser: React.FC<EditUserProps> = props => {
  const { register, handleSubmit, errors: formErrors } = useForm<IUser>({
    defaultValues: props.initUser
  });

  const [editUser] = useMutation<{ updatedUser: IUserGQL }, { user: IUserGQL }>(
    EDIT_USER,
    {
      onCompleted: () => {
        props.close();
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
      }
    }
  );

  const editMode = props.initUser ? true : false;
  const title = editMode ? "Edit User" : "Create User";

  const onSubmit = (user: IUser) => {
    if (editMode && props.initUser) {
      user.id = props.initUser.id;
      const userGQL = transformUserToGQL(user);
      editUser({
        variables: {
          user: userGQL
        }
      });
    }
  };

  return (
    <div className={"modal" + (props.show ? " is-active" : "")}>
      <div className="modal-background"></div>
      <div className="modal-card">
        <header className="modal-card-head">
          <p className="modal-card-title">{title}</p>
          <button
            onClick={props.close}
            className="delete"
            aria-label="close"
          ></button>
        </header>
        <section className="modal-card-body">
          <div className="field">
            <label className="label">Name</label>
            <div className="control">
              <input
                className={`input ${
                  formErrors.name ? "is-danger" : "is-primary"
                }`}
                name="name"
                type="text"
                placeholder="Name"
                ref={register({ required: true })}
              />
              {formErrors.name ? (
                <p className="help is-danger">Name required</p>
              ) : null}
            </div>
          </div>
          <div className="field">
            <label className="label">Email</label>
            <div className="control">
              <input
                className={`input ${
                  formErrors.email ? "is-danger" : "is-primary"
                }`}
                name="email"
                type="email"
                placeholder="Email"
                ref={register({ required: true, pattern: emailPattern })}
              />
              {formErrors.email ? (
                <p className="help is-danger">
                  A valid email address is required
                </p>
              ) : null}
            </div>
          </div>
          <div className="field">
            <label className="label">Image URL</label>
            <div className="control">
              <input
                className={`input ${
                  formErrors.email ? "is-danger" : "is-primary"
                }`}
                name="imageUri"
                type="url"
                placeholder="Image URL"
                ref={register({ pattern: imageUrlPattern })}
              />
              {formErrors.imageUri ? (
                <p className="help is-danger">Not a valid URL</p>
              ) : null}
            </div>
          </div>
          <div className="field">
            <label className="label">Roles</label>
            <div className="control">
              <div className="columns is-centered">
                <div className="column has-text-centered">
                  <label className="checkbox">
                    <input type="checkbox" name="roles.admin" ref={register} />
                    Admin
                  </label>
                </div>
                <div className="column has-text-centered">
                  <label className="checkbox">
                    <input type="checkbox" name="roles.editor" ref={register} />
                    Editor
                  </label>
                </div>
              </div>
            </div>
          </div>
        </section>
        <footer className="modal-card-foot">
          <button onClick={handleSubmit(onSubmit)} className="button is-info">
            Save changes
          </button>
          <button onClick={props.close} className="button is-danger">
            Cancel
          </button>
        </footer>
      </div>
    </div>
  );
};

export default EditUser;
