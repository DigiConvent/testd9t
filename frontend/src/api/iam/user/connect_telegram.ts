import type Either from "@/api/core/either";
import { api_post } from "@/api/core/fetch";

export async function connect_telegram(init_data: string) : Promise<Either<string, boolean>> {
   return api_post<boolean>("/api/iam/auth/telegram/connect", { payload: init_data }, undefined, 200);
}