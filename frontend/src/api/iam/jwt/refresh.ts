import type Either from "../../core/either";
import { api_post } from "../../core/fetch";

export default async function refresh(token: string) : Promise<Either<string, string>> {
    return api_post<string>("/api/iam/jwt/refresh", {"payload": token}, (data) => {
        return data.jwt
    })
}