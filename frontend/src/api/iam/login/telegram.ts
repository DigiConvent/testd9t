import type Either from "../../core/either";
import { api_post } from "../../core/fetch";

export default async function telegram(dataString: string) : Promise<Either<string, string>> {
    return api_post<string>("/api/iam/login/telegram", {"payload": dataString}, (data) => {
        return data.jwt
    })
}