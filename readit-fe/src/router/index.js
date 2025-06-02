import { createRouter, createWebHistory } from 'vue-router';
import SigninView from '../views/SigninView.vue';
import SignupView from '../views/SignupView.vue';
import PostView from '../views/PostView.vue';

const routes = [
  { path: "/signup", component: SignupView },
  { path: "/signin", component: SigninView },
  { path: "/posts", component: PostView, meta: { requiresAuth: true } },
  { path: "/", redirect: "/posts" },
  { path: "/:pathMatch(.*)*", redirect: "/posts" }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

function getTokenFromCookie() {
  const match = document.cookie.match(new RegExp('(^| )token=([^;]+)'))
  return match ? match[2] : null;
}

router.beforeEach((to, from, next) => {
  const isAuthenticated = !!getTokenFromCookie();
  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ path: '/signin' });
  } else {
    next();
  }
})

export default router;