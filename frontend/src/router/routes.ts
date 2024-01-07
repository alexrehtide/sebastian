import { RouteRecordRaw } from "vue-router";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: () => import("views/site/SiteLayout.vue"),
    children: [],
  },
  {
    path: "/admin",
    component: () => import("views/admin/AdminLayout.vue"),
    children: [],
  },
  {
    path: "/map",
    component: () => import("views/map/MapLayout.vue"),
    children: [],
  },
  {
    path: "/auth",
    component: () => import("views/auth/AuthLayout.vue"),
    children: [
      {
        path: "sign_in",
        component: () => import("views/auth/sign-in/SignInPage.vue"),
      },
      {
        path: "code/:platform",
        component: () => import("views/auth/sign-in/SignInPage.vue"),
        props: (to) => ({
          exchange: { platform: to.params["platform"], code: to.query["code"], state: to.query["state"] },
        }),
      },
      {
        path: "sign_up",
        component: () => import("views/auth/sign-up/SignUpPage.vue"),
        props: (to) => ({
          verificationCode: to.query["verification_code"],
        }),
      },
      {
        path: "forgot_password",
        component: () => import("views/auth/forgot-password/ForgotPasswordPage.vue"),
        props: (to) => ({
          resettingCode: to.query["resetting_code"],
        }),
      },
      {
        path: "blocked",
        component: () => import("views/auth/blocked/BlockedPage.vue"),
      },
    ],
  },
  {
    path: "/account",
    component: () => import("views/account/AccountLayout.vue"),
    children: [],
  },
];

export default routes;
