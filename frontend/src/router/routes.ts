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
];

export default routes;
