import type Either from "@/api/core/either";
import { api_get } from "@/api/core/fetch";
import type { UserFacade } from "./types";
import type { Page } from "@/api/core/page";
import { to_user_facade } from "./utils";

export default async function list_users() : Promise<Either<string, Page<UserFacade>>> {
    return api_get<Page<UserFacade>>("/api/iam/user/", (data: any) => { 
        return {
            items: data.items.map((entry: any) => to_user_facade(entry)),
            page: data.page,
            items_per_page: data.items_per_page,
            total_items: data.total_items
        }     
    });
}