import type { RouteRecordRaw } from "vue-router"

export default function generate_routes(
   prefix: string[],
   structure: any | string[],
): RouteRecordRaw[] {
   if (typeof structure.length == "number") {
      return structure.map((key: string) => {
         const path = key
         if (key.startsWith(":") && key.includes("/")) key = key.split("/")[1]

         return {
            path: path.replace("_", "-"),
            name: prefix.concat(key).join("."),
            component: () => import("../views/" + prefix.concat(key).join("/") + ".vue"),
         }
      })
   }
   return Object.keys(structure).map((key: string) => {
      const kids = generate_routes(prefix.concat(key), structure[key])
      return {
         path: key.replace("_", "-"),
         name: prefix.concat(key).join("."),
         children: kids,
         // component: () => import("../views/" + prefix.concat(key).join("/") + ".vue"),
      }
   })
}
