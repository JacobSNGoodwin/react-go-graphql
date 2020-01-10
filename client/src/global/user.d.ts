interface IUserGQL {
  id: string;
  name?: string;
  email: string;
  imageUri?: string;
  roles: string[];
}

interface IUser {
  id: string;
  name?: string;
  email: string;
  imageUri?: string;
  roles: IRole;
}

interface IUserData {
  users: IUserGQL[];
}

interface IUserVars {
  limit?: number;
  offset?: number;
}

interface IRole {
  admin: boolean;
  editor: boolean;
}
