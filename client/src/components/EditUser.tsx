import React from "react";
import { useForm } from "react-hook-form";

import Spinner from "./ui/Spinner";

import {
  emailPattern,
  imageUrlPattern,
  transformUserToGQL
} from "../util/util";

interface EditUserProps {
  show: boolean;
  initUser?: IUser;
  editingUser: boolean;
  close: () => void;
  editSelectedUser: (gqlUser: IUserGQL) => void;
}

const EditUser: React.FC<EditUserProps> = props => {
  const { register, handleSubmit, errors: formErrors } = useForm<IUser>({
    defaultValues: props.initUser
  });

  const editMode = props.initUser ? true : false;
  const title = editMode ? "Edit User" : "Create User";

  const onSubmit = (user: IUser) => {
    if (editMode && props.initUser) {
      user.id = props.initUser.id;
      const userGQL = transformUserToGQL(user);
      props.editSelectedUser(userGQL);
    } else {
      const userGQL = transformUserToGQL(user);
      props.editSelectedUser(userGQL);
    }
  };

  return (
    <div
      className={"has-text-centered modal" + (props.show ? " is-active" : "")}
    >
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
          <button onClick={handleSubmit(onSubmit)} className="button is-danger">
            Save changes
          </button>
          <button onClick={props.close} className="button is-info">
            Cancel
          </button>
          {props.editingUser ? <Spinner radius={20} /> : null}
        </footer>
      </div>
    </div>
  );
};

export default EditUser;
