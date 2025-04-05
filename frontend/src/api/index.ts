import { sys } from "./sys"
import { iam } from "./iam"
import type { ApiCollection } from "./core/endpoint"

export const api = {
   iam: iam,
   sys: sys,
} satisfies ApiCollection
