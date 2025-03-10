import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      children: [
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
                path: "user",
                name: "iam.user",
                component: () => import("../views/admin/iam/user/list.vue"),
                children: [
                  {
                    path: "create",
                    component: () => import("../components/iam/user/create.vue")
                  }
                ]
              }
            ]},
            {
              path: "sys",
              name: "sys",
              component: () => import("../views/admin/sys/sys.vue"),
              children: [],
            }
          ]
        },        
      ]
    },
  ]
})

export default router
