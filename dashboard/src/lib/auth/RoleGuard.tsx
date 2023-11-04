import { Navigate } from "react-router-dom";
import { ReactNode } from "react";

// List of all roles in the Dashboard
const rolesLevel = {
	SUPER: 0,
	ADMIN: 1,
	DEVELOPER: 2,
	MODERATOR: 3,
	GUEST: 4,
} as const;

// Making a type for roles
type UserRole = keyof typeof rolesLevel;
// type UserType = typeof rolesLevel; // { readonly SUPER: number, readonly ADMIN: number, ... }
// type UserType = (typeof rolesLevel)["GUEST"]; // 4
// type UserType = (typeof rolesLevel)["GUEST" | "MODERATOR"]; // 4 | 3
// type UserType = (typeof rolesLevel)[UserRole]; // 4 | 3 | 2 | 1 | 0

/**
 * Check User Role before redirecting
 */
export const ProtectedRoles = ({
	role,
	children,
}: {
	role: UserRole;
	children: ReactNode;
}) => {
	// if (UserRoles[user.role] > UserRoles[role] ) return <Navigate to="/" replace />;
	if (rolesLevel["SUPER"] > rolesLevel[role])
		return <Navigate to="/" replace />;
	return children;
};

// Exporting all roles as array
export const userRoles = [...Object.keys(rolesLevel)] as const;
