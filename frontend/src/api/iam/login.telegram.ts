import type Either from "../core/either";
import { apiPostRequest } from "../core/fetch";

export default async function loginWithTelegram(dataString: string) : Promise<Either<string, string>> {
    return apiPostRequest<string>("/api/iam/login/telegram", {"payload": dataString}, (data) => {
        return data.jwt
    })
}