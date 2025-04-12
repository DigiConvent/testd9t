import { createRouter, createWebHistory } from "vue-router"
import generate_routes from "./generate"

const router = createRouter({
   history: createWebHistory(),
   routes: [
      {
         path: "/",
         name: "",
         redirect: "/app",
      },
      {
         path: "/logout",
         name: "logout",
         component: () => import("../views/auth/logout.vue"),
      },
      {
         path: "/app",
         name: "home",
         children: [
            {
               path: "connect-telegram-user",
               name: "connect-telegram-user",
               component: () => import("../views/auth/connect_telegram_user.vue"),
            },
            generate_routes([], {
               admin: {
                  iam: {
                     user: [
                        "create",
                        "list",
                        ":id/read",
                        ":id/update",
                        ":id/delete",
                        ":id/profile",
                     ],
                     permission: [":name/profile"],
                     permission_group: [
                        "create",
                        "list",
                        ":id/read",
                        ":id/update",
                        ":id/delete",
                        ":id/profile",
                     ],
                     user_role: [
                        "create",
                        "list",
                        ":id/read",
                        ":id/update",
                        ":id/delete",
                        ":id/profile",
                     ],
                     user_status: [
                        "create",
                        "list",
                        ":id/read",
                        ":id/update",
                        ":id/delete",
                        ":id/profile",
                     ],
                  },
                  sys: {},
               },
            })[0],
            {
               path: "user",
               name: "user",
               children: [
                  {
                     path: "settings",
                     name: "user.settings",
                     component: () => import("../views/user/settings.vue"),
                  },
                  {
                     path: "profile",
                     name: "user.profile",
                     component: () => import("../views/user/profile.vue"),
                  },
               ],
            },
         ],
      },
   ],
})

router.beforeEach((to, from, next) => {
   window.Telegram.WebApp.HapticFeedback.impactOccurred("light")
   next()
})

export default router
