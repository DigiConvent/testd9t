import type Either from "@/api/core/either";
import { api_get } from "@/api/core/fetch";
import type { PermissionGroupProfile } from "./types";
import { to_permission_group_profile } from "./utils";

export default async function get_permission_group_profile(pid: string) : Promise<Either<string, PermissionGroupProfile>> {
    return api_get<PermissionGroupProfile>("/api/iam/permission-group/profile/" + pid + "/", to_permission_group_profile);
}