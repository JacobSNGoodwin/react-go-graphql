import React from "react";

import placeholder from "../images/placeholder.png";
import EditUser from "./EditUser";

interface UserProps {
  user: IUser;
}

const User: React.FC<UserProps> = props => {
  const [editActive, setEditActive] = React.useState<boolean>(false);

  const userRoles: string[] = [];

  if (props.user.roles.admin) {
    userRoles.push("Admin");
  }

  if (props.user.roles.editor) {
    userRoles.push("Editor");
  }

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

        <div
          style={{
            display: "flex",
            justifyContent: "center",
            paddingTop: "1.25em"
          }}
        >
          <button
            onClick={() => setEditActive(true)}
            style={{ border: "none" }}
            className="button is-primary"
          >
            Edit
          </button>
        </div>

        <EditUser
          show={editActive}
          close={() => setEditActive(false)}
          initUser={props.user}
        />
      </div>
    </div>
  );
};

export default User;
