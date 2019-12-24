import { gql } from "apollo-boost";

const ME = gql`
  query me {
    me {
      id
      name
      email
      imageUri
      roles
    }
  }
`;

export { ME };
