import type Either from "@/api/core/either";
import { api_get } from "@/api/core/fetch";
import type { UserStatusRead } from "..";


export default async function list_user_status() : Promise<Either<string, UserStatusRead[]>> {
    return api_get<UserStatusRead[]>("/api/iam/user-status/", (data: any) => { 
        const result: UserStatusRead[] = [];
        for (const entry of data.items) {
            result.push({
                id: entry.id,
                name: entry.name,
                abbr: entry.abbr,
                description: entry.description,
                archived: entry.archived
            });
        };
        return result
    });
}