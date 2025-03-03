import getWebApp from "./telegram";

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

    async loginUsingTelegram() {
        const body: BodyInit = new FormData();
        body.set('payload', getWebApp().initData);
        this.login(await fetch("/api/auth/telegram", {
            method: "POST",
            body: body,
        }));
    }

    async loginUsingCredentials(emailaddress: string, password: string) {
        const body: BodyInit = new FormData();
        body.set('emailaddress', emailaddress);
        body.set('password', password);
        this.login(await fetch("/api/auth/credentials", {
            method: "POST",
            body: body,
        }));
    }

    async login(response: Response) {
        if (response.ok) {
            const data = await response.json();
            this.token = data.token;
            localStorage.setItem('token', data.token);
        }
        console.log("Logged in with token " + this.token);
    }

    async loginUsingToken() {
        const body: BodyInit = new FormData();
        body.set('token', this.token!);
        this.login(await fetch("/api/auth/token", {
            method: "POST",
            body: body,
        }));
    }
}