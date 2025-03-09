import type { UserFacade } from "./types";

export function to_user_facade(data: any) : UserFacade {
	return {
		id: data.id,
		name: data.name,
		implied: data.implied,
		status_id: data.status_id,
		status_name: data.status_name
	}
}