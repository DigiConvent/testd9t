import type Either from "../core/either";
import { apiPostRequest } from "../core/fetch";

export default async function refreshJwt(token: string) : Promise<Either<string, string>> {
    return apiPostRequest<string>("/api/iam/jwt/refresh", {token}, (data) => {
        return data.jwt
    })
}