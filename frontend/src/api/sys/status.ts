import type Either from "../core/either"
import { api_get } from "../core/fetch"
import { to_system_status, type SystemStatus } from "./types";



export async function get_status() : Promise<Either<string, SystemStatus>> {
    return await api_get<SystemStatus>("/api/sys/status/", to_system_status);
}