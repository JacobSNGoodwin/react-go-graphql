import React from "react";
import Spinner from "./ui/Spinner";

interface DeleteUserProps {
  show: boolean;
  user: IUser;
  deletingUser: boolean;
  close: () => void;
  deleteSelectedUser: () => void;
}

const DeleteUser: React.FC<DeleteUserProps> = props => {
  return (
    <div className={"modal" + (props.show ? " is-active" : "")}>
      <div className="modal-background"></div>
      <div className="modal-card">
        <header className="modal-card-head">
          <p className="modal-card-title">Confirm Delete</p>
          <button
            onClick={props.close}
            className="delete"
            aria-label="close"
          ></button>
        </header>
        <section className="modal-card-body">
          <p>Are you sure you want to delete user the following user?</p>
          <p>{props.user.name}</p>
          <p>{props.user.email}</p>
        </section>
        <footer className="modal-card-foot">
          <button onClick={props.deleteSelectedUser} className="button is-info">
            Delete
          </button>
          <button onClick={props.close} className="button is-danger">
            Cancel
          </button>
          {props.deletingUser ? <Spinner radius={20} /> : null}
        </footer>
      </div>
    </div>
  );
};

export default DeleteUser;
