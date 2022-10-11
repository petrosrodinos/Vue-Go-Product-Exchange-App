import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import store from "../state/store";
const Home = require("../views/Home.vue").default;
const Products = require("../views/user/product/Products.vue").default;
const Trades = require("../views/user/trades/Trades.vue").default;
const Settings = require("../views/user/Settings.vue").default;
const SendFeedBack = require("../views/utils/SendFeedBack.vue").default;
const About = require("../views/utils/About.vue").default;
const Login = require("../views/auth/Login.vue").default;
const Register = require("../views/auth/Register.vue").default;
const ResetPassword = require("../views/auth/ResetPassword.vue").default;
const AddProduct = require("../views/user/product/AddProduct.vue").default;
const Profile = require("../views/user/Profile.vue").default;
const EditProduct = require("../views/user/product/EditProduct.vue").default;

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Home,
    meta: { requiresAuth: false },
  },
  {
    path: "/profile/products/:id",
    name: "Product",
    component: EditProduct,
    props: true,
    meta: { requiresAuth: true },
  },
  {
    path: "/profile/trades",
    name: "Trades",
    component: Trades,
    meta: { requiresAuth: true },
  },
  {
    path: "/profile/products",
    name: "Products",
    component: Products,
    meta: { requiresAuth: true },
  },
  {
    path: "/profile/:id",
    name: "Profile",
    component: Profile,
    props: true,
    meta: { requiresAuth: false },
  },
  {
    path: "/profile/products/add",
    name: "AddProduct",
    component: AddProduct,
    meta: { requiresAuth: true },
  },
  {
    path: "/profile/settings",
    name: "Settings",
    component: Settings,
    props: true,
    meta: { requiresAuth: true },
  },
  {
    path: "/auth/login",
    name: "Login",
    component: Login,
    meta: { requiresAuth: false, authPage: true },
  },
  {
    path: "/auth/register",
    name: "Register",
    component: Register,
    meta: { requiresAuth: false, authPage: true },
  },
  {
    path: "/auth/reset-password",
    name: "ResetPassword",
    component: ResetPassword,
    meta: { requiresAuth: false },
  },
  {
    path: "/feedback",
    name: "FeedBack",
    component: SendFeedBack,
    meta: { requiresAuth: false },
  },
  {
    path: "/about",
    name: "About",
    component: About,
    meta: { requiresAuth: false },
  },
  {
    path: "/:catchAll(.*)",
    name: "not-found",
    component: () => require("../views/utils/NotFound.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach((to, from, next) => {
  const isLoggedIn = !!store.getters.getToken;
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    if (!isLoggedIn) {
      next({
        path: "/auth/login",
      });
    } else {
      next();
    }
  } else if (to.matched.some((record) => record.meta.authPage)) {
    if (isLoggedIn) {
      next({
        path: "/",
      });
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
