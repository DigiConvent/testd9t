import type Either from "../core/either";
import { apiPostRequest } from "../core/fetch";
import refreshJwt from "./jwt.refresh";
import loginWithCredentials from "./login.credentials";
import loginWithTelegram from "./login.telegram";

export interface IamRoutes {
    loginWithTelegram: (dataString: string) => Promise<Either<string, string>>;
    loginWithCredentials: (emailaddress: string, password: string) => Promise<Either<string, string>>;
    refreshJwt: (token: string) => Promise<Either<string, string>>;
}

export async function getIamRoutes(): Promise<IamRoutes> {
    return {
        loginWithTelegram: loginWithTelegram,
        loginWithCredentials: loginWithCredentials,
        refreshJwt: refreshJwt
    };
}