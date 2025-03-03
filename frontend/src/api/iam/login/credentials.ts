import type Either from "@/api/core/either"
import { apiPostRequest } from "@/api/core/fetch"

export default async function credentials(emailaddress: string, password: string) : Promise<Either<string, string>> {
    return apiPostRequest<string>("/api/iam/login/credentials", { emailaddress, password }, (data) => {
        return data.jwt
    })
}   