import type { ApiCollection } from "@/api/core/endpoint"
import refresh from "./refresh"

export const jwt: ApiCollection = {
   refresh: refresh,
}
