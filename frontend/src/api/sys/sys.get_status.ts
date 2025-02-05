import type Either from "../core/either"
import { apiGetRequest } from "../core/fetch"

export type SystemStatus = {
	online_since: string
	version: {
        major: string,
        minor: string,
        patch: string,
    }
	database_version: {
        major: string,
        minor: string,
        patch: string,
    }
	free_space: number
	total_space: number
	data_space: number
	built_at: string
}

export async function getStatus() : Promise<Either<string, SystemStatus>> {
    return await apiGetRequest<SystemStatus>("/api/sys/status");
}