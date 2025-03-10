import { api } from "@/api";
import get_web_app from "./telegram";
import type Either from "@/api/core/either";
import { ref } from "vue";
import type { PermissionFacade } from "@/api/iam/permission/types";

export default class JwtAuthenticator {
    private static instance: JwtAuthenticator | undefined;
    private token: string | undefined;
    public is_authenticated = ref<boolean>(false);
    private permissions = ref<string[]>([]);
    
    constructor() {
        const t = localStorage.getItem('token');
        if (t) {
            this.token = t;
            this.is_authenticated.value = true;
            this.load_permissions();
        } else {
            this.is_authenticated.value = false;
        }
    }

    public get_token() : { id: string, exp: number, tgid: number } | null {
        try {
            if (this.token == undefined) {
                return null;
            }
            const payload_base64 = this.token.split(".")[1];
            const payload_json = atob(payload_base64);
            return JSON.parse(payload_json);
        } catch {
            return null;
        }
    }

    public has_permission(permission: Exclude<string, 'super'>) : boolean {
        if (this.permissions.value.includes("super")) {
            return true;
        }
        return this.permissions.value.includes(permission);
    }

    public async load_permissions() {
        const result = await api.iam.user.list_permissions();
        result.fold(
            (error: string) => {
                console.error(error);
            },
            (permissions: PermissionFacade[]) => {
                this.permissions.value = permissions.map((p) => p.name);
            }
        )
    }

    static get_instance() : JwtAuthenticator {
        if (this.instance == undefined) {
            this.instance = new JwtAuthenticator();
        }
        return this.instance;
    }

    async login_using_telegram() : Promise<boolean> {
        return this.login(api.iam.login.telegram(get_web_app().initData));
    }

    async login_using_credentials(emailaddress: string, password: string) : Promise<boolean> {
        return this.login(api.iam.login.credentials(emailaddress, password));
    }

    async connect_telegram_user() : Promise<boolean>{
        const result = await api.iam.user.connect_telegram(get_web_app().initData);
        return result.isRight();
    }

    async login(response: Promise<Either<string,string>>) : Promise<boolean> {
        const result = await response;
        if (result.isRight()) {
            const token = result.getRight();
            this.token = token;
            localStorage.setItem('token', token!);
            this.is_authenticated.value = true;
            await this.load_permissions();
            return true;
        } else {
            return false;
        }
    }

    async login_using_token() {
        this.login(api.iam.jwt.refresh(this.token!));
    }

    logout() {
        this.token = undefined;
        localStorage.removeItem('token');
        this.is_authenticated.value = false;
    }
}