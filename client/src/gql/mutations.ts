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

export { LOGIN_GOOGLE, LOGIN_FACEBOOK };
