import { gql } from "apollo-boost";

const LOGIN_GOOGLE = gql`
  mutation LoginGoogle($idToken: String!) {
    googleLoginWithToken(idToken: $idToken) {
      id
      name
      email
      imageUri
      roles
    }
  }
`;

const LOGIN_FACEBOOK = gql`
  mutation fbLogin($accessToken: String!) {
    fbLoginWithToken(accessToken: $accessToken) {
      id
      name
      email
      imageUri
      roles
    }
  }
`;

const EDIT_USER = gql`
mutation editUser($user: EditUserInput!) {
 	editUser(user: $user) {
    id
    name
    email
    imageUri
    roles
  }
}
`;

const DELETE_USER = gql`
  mutation deleteUser($id: String!) {
  deleteUser(id: $id)
}
`;


export { LOGIN_GOOGLE, LOGIN_FACEBOOK, EDIT_USER, DELETE_USER };
