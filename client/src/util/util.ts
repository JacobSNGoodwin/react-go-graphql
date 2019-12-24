import Cookies from "js-cookie";
import jwt from "jsonwebtoken";

export const getUserFromCookie = function(): string | undefined {
  const userCookie = Cookies.get("userinfo");
  if (!userCookie) {
    return undefined;
  }
  const decoded = jwt.decode(userCookie + ".");

  if (!decoded || typeof decoded !== "object") {
    return undefined;
  }

  return decoded["id"]; // should return string or undefined if key doesn't exist
};

export const hasRequiredRole = function(
  userRoles: string[],
  allowedRoles: string[]
): boolean {
  for (let i = 0; i < allowedRoles.length; i++) {
    if (userRoles.includes(allowedRoles[i])) {
      return true;
    }
  }

  return false;
};
