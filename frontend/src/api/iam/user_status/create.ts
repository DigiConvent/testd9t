import type Either from "@/api/core/either";
import { api_post } from "@/api/core/fetch";

export type UserStatusCreate = {
    name: string,
    abbr: string,
    description: string
    archived: boolean
}

export default async function create_user_status(user_status_create: UserStatusCreate) : Promise<Either<string, string>> {
    return api_post<string>("/api/iam/user-status/", user_status_create, (data: any) => {
        return data.id;
    });
}
