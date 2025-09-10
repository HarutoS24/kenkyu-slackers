import AppLayout from '@/layouts/AppLayout.vue'
import AppRoot from '@/pages/AppRoot/AppRoot.vue'
import DevOnlyTestPage from '@/pages/DevOnlyTestPage/DevOnlyTestPage.vue'
import ResultPage from '@/pages/ResultPage/ResultPage.vue'
import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: AppLayout,
    children: [
      { path: "", component: AppRoot },
      { path: "result", component: ResultPage },
    ]
  },
];

if (import.meta.env.DEV) {
  routes.push(
    {
      path: "/dev-only/",
      children: [
        { path: "test", component: DevOnlyTestPage }
      ]
    }
  );
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router
