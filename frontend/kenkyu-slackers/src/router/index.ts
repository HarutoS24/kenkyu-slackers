import AppLayout from '@/layouts/AppLayout.vue'
import AppRoot from '@/pages/AppRoot/AppRoot.vue'
import ResultPage from '@/pages/ResultPage/ResultPage.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      component: AppLayout,
      children: [
        { path: "", component: AppRoot },
        { path: "result", component: ResultPage },
      ]
    }
  ],
})

export default router
