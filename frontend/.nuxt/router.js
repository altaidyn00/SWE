import Vue from 'vue'
import Router from 'vue-router'
import { normalizeURL, decode } from 'ufo'
import { interopDefault } from './utils'
import scrollBehavior from './router.scrollBehavior.js'

const _6697a6d5 = () => interopDefault(import('../pages/about-us.vue' /* webpackChunkName: "pages/about-us" */))
const _5d6a6010 = () => interopDefault(import('../pages/appointment.vue' /* webpackChunkName: "pages/appointment" */))
const _23a1106c = () => interopDefault(import('../pages/contacts.vue' /* webpackChunkName: "pages/contacts" */))
const _d7ff9130 = () => interopDefault(import('../pages/request.vue' /* webpackChunkName: "pages/request" */))
const _209b528a = () => interopDefault(import('../pages/requests.vue' /* webpackChunkName: "pages/requests" */))
const _1df0dd18 = () => interopDefault(import('../pages/services/index.vue' /* webpackChunkName: "pages/services/index" */))
const _641d5359 = () => interopDefault(import('../pages/signin.vue' /* webpackChunkName: "pages/signin" */))
const _7fec6e6e = () => interopDefault(import('../pages/create/doctor/index.vue' /* webpackChunkName: "pages/create/doctor/index" */))
const _63d91d6c = () => interopDefault(import('../pages/create/patient/index.vue' /* webpackChunkName: "pages/create/patient/index" */))
const _e6ecfd6e = () => interopDefault(import('../pages/users/doctors/index.vue' /* webpackChunkName: "pages/users/doctors/index" */))
const _77a22aef = () => interopDefault(import('../pages/users/patients/index.vue' /* webpackChunkName: "pages/users/patients/index" */))
const _96826098 = () => interopDefault(import('../pages/users/doctors/_id/index.vue' /* webpackChunkName: "pages/users/doctors/_id/index" */))
const _ab03934c = () => interopDefault(import('../pages/users/patients/_id/index.vue' /* webpackChunkName: "pages/users/patients/_id/index" */))
const _564a8d18 = () => interopDefault(import('../pages/users/doctors/_id/edit.vue' /* webpackChunkName: "pages/users/doctors/_id/edit" */))
const _7d5e8c9c = () => interopDefault(import('../pages/users/patients/_id/edit.vue' /* webpackChunkName: "pages/users/patients/_id/edit" */))
const _7195106b = () => interopDefault(import('../pages/index.vue' /* webpackChunkName: "pages/index" */))

const emptyFn = () => {}

Vue.use(Router)

export const routerOptions = {
  mode: 'history',
  base: '/',
  linkActiveClass: 'nuxt-link-active',
  linkExactActiveClass: 'nuxt-link-exact-active',
  scrollBehavior,

  routes: [{
    path: "/about-us",
    component: _6697a6d5,
    name: "about-us"
  }, {
    path: "/appointment",
    component: _5d6a6010,
    name: "appointment"
  }, {
    path: "/contacts",
    component: _23a1106c,
    name: "contacts"
  }, {
    path: "/request",
    component: _d7ff9130,
    name: "request"
  }, {
    path: "/requests",
    component: _209b528a,
    name: "requests"
  }, {
    path: "/services",
    component: _1df0dd18,
    name: "services"
  }, {
    path: "/signin",
    component: _641d5359,
    name: "signin"
  }, {
    path: "/create/doctor",
    component: _7fec6e6e,
    name: "create-doctor"
  }, {
    path: "/create/patient",
    component: _63d91d6c,
    name: "create-patient"
  }, {
    path: "/users/doctors",
    component: _e6ecfd6e,
    name: "users-doctors"
  }, {
    path: "/users/patients",
    component: _77a22aef,
    name: "users-patients"
  }, {
    path: "/users/doctors/:id",
    component: _96826098,
    name: "users-doctors-id"
  }, {
    path: "/users/patients/:id",
    component: _ab03934c,
    name: "users-patients-id"
  }, {
    path: "/users/doctors/:id/edit",
    component: _564a8d18,
    name: "users-doctors-id-edit"
  }, {
    path: "/users/patients/:id/edit",
    component: _7d5e8c9c,
    name: "users-patients-id-edit"
  }, {
    path: "/",
    component: _7195106b,
    name: "index"
  }],

  fallback: false
}

export function createRouter (ssrContext, config) {
  const base = (config._app && config._app.basePath) || routerOptions.base
  const router = new Router({ ...routerOptions, base  })

  // TODO: remove in Nuxt 3
  const originalPush = router.push
  router.push = function push (location, onComplete = emptyFn, onAbort) {
    return originalPush.call(this, location, onComplete, onAbort)
  }

  const resolve = router.resolve.bind(router)
  router.resolve = (to, current, append) => {
    if (typeof to === 'string') {
      to = normalizeURL(to)
    }
    return resolve(to, current, append)
  }

  return router
}
