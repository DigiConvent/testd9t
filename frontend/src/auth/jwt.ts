import { api } from "@/api";
import getWebApp from "./telegram";
import type Either from "@/api/core/either";

export default class JwtAuthenticator {
    private static instance: JwtAuthenticator | undefined;
    private token: string | undefined;
    constructor() {
        const t = localStorage.getItem('token');
        if (t) {
            this.token = t;
        }
    }

    static getInstance() : JwtAuthenticator {
        if (this.instance == undefined) {
            this.instance = new JwtAuthenticator();
        }
        return this.instance;
    }

    getToken(data: any) : string {
        return data.token;
    }

    async loginUsingTelegram() {
        this.login(api.iam.login.telegram(getWebApp().initData));
    }

    async loginUsingCredentials(emailaddress: string, password: string) {
        this.login(api.iam.login.credentials(emailaddress, password));
    }

    async login(response: Promise<Either<string,string>>) {
        (await response).fold((err: string) => {
            console.error(err);
        }, (token: string) => {
            this.token = token;
        });
    }

    async loginUsingToken() {
        this.login(api.iam.jwt.refresh(this.token!));
    }
}