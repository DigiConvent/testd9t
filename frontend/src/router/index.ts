import { createRouter, createWebHistory } from "vue-router"

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
            {
               path: "admin",
               name: "admin",
               component: () => import("../views/admin/admin.vue"),
               children: [
                  {
                     path: "iam",
                     name: "iam",
                     component: () => import("../views/admin/iam/iam.vue"),
                     children: [
                        {
                           path: "users",
                           name: "iam.user.list",
                           component: () => import("../views/admin/iam/user/list.vue"),
                           children: [
                              {
                                 path: "create",
                                 component: () => import("../components/iam/user/create.vue"),
                              },
                           ],
                        },
                        {
                           path: "users/:id",
                           name: "iam.user.profile",
                           component: () => import("../views/admin/iam/user/profile.vue"),
                           children: [
                              {
                                 path: "create",
                                 component: () => import("../components/iam/user/create.vue"),
                              },
                           ],
                        },
                        {
                           path: "permission-group/:id",
                           name: "iam.pg.profile",
                           component: () =>
                              import("../views/admin/iam/permission_group/profile.vue"),
                        },
                     ],
                  },
                  {
                     path: "sys",
                     name: "sys",
                     component: () => import("../views/admin/sys/sys.vue"),
                     children: [],
                  },
               ],
            },
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

export default router
