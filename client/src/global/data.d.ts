// User
interface IUserGQL {
  id: string;
  name?: string;
  email: string;
  imageUri?: string;
  roles: string[];
}

interface IUserCreateGQL {
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

// Role

interface IRole {
  admin: boolean;
  editor: boolean;
}

// Product

interface IProductGQL {
  id: string;
  name?: string;
  price: number;
  imageUri?: string;
  location?: string;
  categories?: ICategoryGQL[];
}

interface IProductData {
  products: IProductGQL[];
}

interface IProductVars {
  limit?: number;
  offset?: number;
}

// Categories

interface ICategoyGQL {
  id: string;
  title?: string;
  description?: string;
  products?: IProductGQL[];
}

interface ICategoryData {
  categories: ICategoryGQL[];
}
