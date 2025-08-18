import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import ShorUrl from '../views/ShorUrl.vue';

export const router = createRouter({
  routes: [
    {
      path: '/',
      component: Home,
    },
    {
      path: '/short',
      component: ShorUrl,
    },
  ],
  history: createWebHistory(),
});