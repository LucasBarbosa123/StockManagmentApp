import { createApp } from "vue";
import { createWebHistory, createRouter } from "vue-router";

// styles

import "@fortawesome/fontawesome-free/css/all.min.css";
import "@/assets/styles/tailwind.css";

// mouting point for the whole app

import App from "@/App.vue";

// layouts

import Admin from "@/layouts/Admin.vue";
// views for Admin layout

import Index from "@/views/Index.vue";
import Users from "@/views/Users.vue";

// routes

const routes = [
  {
    path: "/",
    redirect: "/Home",
    component: Admin,
    children: [
      {
        path: "/Home",
        component: Index,
        meta: {
          title: 'Inicio'
        },  
      },
      {
        path: "/Users",
        component: Users,
        meta: {
          title: 'Utilizadores'
        },  
      },
    ],
  },
  { path: "/:pathMatch(.*)*", redirect: "/" },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

createApp(App).use(router).mount("#app");
