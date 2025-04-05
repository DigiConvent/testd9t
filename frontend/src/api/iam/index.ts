import type { ApiCollection } from "../core/endpoint"
import { jwt } from "./jwt"
import { login } from "./login"
import { permission } from "./permission"
import { permission_group } from "./permission_group"
import { user } from "./user"
import { user_status } from "./user_status"

export const iam = {
   jwt: jwt,
   login: login,
   permission_group: permission_group,
   permission: permission,
   user_status: user_status,
   user: user,
} satisfies ApiCollection
