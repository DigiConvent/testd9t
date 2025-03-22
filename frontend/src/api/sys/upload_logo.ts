import type Either from "../core/either"
import { api_multipart } from "../core/fetch"

export default function upload_logo(
   variant: "small" | "large",
   file: File,
): Promise<Either<string, boolean>> {
   return api_multipart<boolean>("/api/sys/logo/" + variant, {}, new Map().set("file", file), 200)
}
