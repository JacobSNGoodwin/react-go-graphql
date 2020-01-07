import React from "react";
import { useForm } from "react-hook-form";

interface EditUserProps {
  show: boolean;
  close: () => void;
  initUser?: IUser;
}

const EditUser: React.FC<EditUserProps> = props => {
  const { register, handleSubmit } = useForm<IUser>();

  const editMode = props.initUser ? true : false;
  const title = editMode ? "Edit User" : "Create User";

  const onSubmit = (data: any) => {
    console.log(data);

    props.close();
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
                className="input"
                name="name"
                type="text"
                placeholder="Name"
                ref={register}
              />
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
