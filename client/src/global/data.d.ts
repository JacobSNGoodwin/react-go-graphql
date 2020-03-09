// General
interface IQueryVars {
  limit?: number;
  offset?: number;
}

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

interface IUserVars extends IQueryVars {}

// Role

interface IRole extends IQueryVars {
  admin: boolean;
  editor: boolean;
}

// Product

interface IProduct {
  id: string;
  name?: string;
  price: number;
  imageUri?: string;
  location?: string;
  categories?: ICategoryGQL[];
}

interface IProductData {
  products: IProduct[];
}

interface IProductVars extends IQueryVars {}

// Categories

interface ICategoy {
  id: string;
  title?: string;
  description?: string;
  products?: IProductGQL[];
}

interface ICategoryData {
  categories: ICategory[];
}

interface ICategoryVars extends IQueryVars {}
