import type Either from "../../core/either";
import { api_post } from "../../core/fetch";

export default async function telegram(data_string: string) : Promise<Either<string, string>> {
    return api_post<string>("/api/iam/login/telegram", {"payload": data_string}, (data) => {
        return data.jwt
    })
}