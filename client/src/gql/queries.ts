import {gql} from 'apollo-boost';

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

const GET_USERS = gql`
  query getUsers($limit: Int, $offset: Int){
    users(limit: $limit, offset: $offset) {
      id
      name
      email
      roles
      imageUri
    }
  }
`;

export {ME, GET_USERS};
