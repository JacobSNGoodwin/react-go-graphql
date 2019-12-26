import React from "react";
import { useQuery } from "@apollo/react-hooks";

import Spinner from "./ui/Spinner";
import { GET_USERS } from "../gql/queries";
import placeholder from "../images/placeholder.png";

// Maybe store this in common types later
interface IUser {
  id: string;
  name?: string;
  email: string;
  imageUri?: string;
  roles: string[];
}

interface IUserData {
  users: IUser[];
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
        <Spinner radius={50} />;
      </div>
    );

  if (error) {
    return (
      <div className="container">
        <h1>Error</h1>
        <p>{error.message}</p>
      </div>
    );
  }

  const userList =
    data &&
    data.users.map(user => (
      <div className="container is-fullhd" key={user.id}>
        <div className="card">
          <div className="card-content has-text-centered">
            <div
              style={{
                display: "flex",
                justifyContent: "center",
                padding: "1.5em 1em"
              }}
            >
              <figure className="image is-96x96">
                {!user.imageUri || user.imageUri === "" ? (
                  <img
                    className="is-rounded"
                    src={placeholder}
                    alt="No profile"
                  />
                ) : (
                  <img
                    className="is-rounded"
                    src={user.imageUri}
                    alt="User profile"
                  />
                )}
              </figure>
            </div>

            <p className="title is-4">{user.name}</p>
            <p className="subtitle is-6">{user.email}</p>
            <br />
            <div className="is-size-5 has-text-weight-bold">Roles</div>
            <div>{user.roles.join(", ")}</div>

            <div
              style={{
                display: "flex",
                justifyContent: "center",
                paddingTop: "1.25em"
              }}
            >
              <button style={{ border: "none" }} className="button is-primary">
                Edit
              </button>
            </div>
          </div>
        </div>
      </div>
    ));

  return (
    <div className="section">
      <h1 className="title is-1 has-text-centered">Users</h1>
      {userList}
    </div>
  );
};

export default Users;
