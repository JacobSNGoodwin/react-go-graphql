import Cookies from 'js-cookie';
import jwt from 'jsonwebtoken';

export const getUserFromCookie = function (): string | undefined {
  const userCookie = Cookies.get('userinfo');
  if (!userCookie) {
    return undefined;
  }
  const decoded = jwt.decode(userCookie + '.');

  if (!decoded || typeof decoded !== 'object') {
    return undefined;
  }

  return decoded['id'];  // should return string or undefined if key doesn't
  // exist
};

export const hasRequiredRole = function (
  userRoles: string[], allowedRoles: string[]): boolean {
  for (let i = 0; i < allowedRoles.length; i++) {
    if (userRoles.includes(allowedRoles[i])) {
      return true;
    }
  }

  return false;
};

export const transformUserToGQL = function (user: IUser): IUserGQL {
  const userGQL: IUserGQL = {
    id: user.id,
    name: user.name,
    email: user.email,
    imageUri: user.imageUri,
    roles: []
  }

  if (user.roles.admin) {
    userGQL.roles.push("Admin")
  }

  if (user.roles.editor) {
    userGQL.roles.push("Editor")
  }

  return userGQL;
};

export const transformUserFromGQL = function (gqlUser: IUserGQL): IUser {
  const user: IUser = {
    id: gqlUser.id,
    name: gqlUser.name,
    email: gqlUser.email,
    imageUri: gqlUser.imageUri,
    roles: {
      admin: false,
      editor: false,
    }
  }

  gqlUser.roles.forEach(role => {
    if (role === 'Admin') {
      user.roles.admin = true;
    }

    if (role === 'Editor') {
      user.roles.editor = true;
    }
  })

  return user;
};

export const emailPattern: RegExp =
  // eslint-disable-next-line no-control-regex, no-useless-escape
  /(?!.*\.{2})^([a-z\d!#$%&'*+\-\/=?^_`{|}~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+(\.[a-z\d!#$%&'*+\-\/=?^_`{|}~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+)*|"((([ \t]*\r\n)?[ \t]+)?([\x01-\x08\x0b\x0c\x0e-\x1f\x7f\x21\x23-\x5b\x5d-\x7e\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|\\[\x01-\x09\x0b\x0c\x0d-\x7f\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]))*(([ \t]*\r\n)?[ \t]+)?")@(([a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|[a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF][a-z\d\-._~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]*[a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])\.)+([a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|[a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF][a-z\d\-._~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]*[a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])\.?$/i;

// eslint-disable-next-line no-useless-escape
export const imageUrlPattern: RegExp = /^((http[s]?|ftp):\/)?\/?([^:\/\s]+)((\/\w+)*\/)([\w\-\.]+[^#?\s]+)(.*)?(#[\w\-]+)?$/;