import Vue from 'vue'
import Router from 'vue-router'
import { normalizeURL, decode } from 'ufo'
import { interopDefault } from './utils'
import scrollBehavior from './router.scrollBehavior.js'

const _1d06589e = () => interopDefault(import('../pages/about-us.vue' /* webpackChunkName: "pages/about-us" */))
const _a62444e2 = () => interopDefault(import('../pages/appointment.vue' /* webpackChunkName: "pages/appointment" */))
const _b6c3acda = () => interopDefault(import('../pages/contacts.vue' /* webpackChunkName: "pages/contacts" */))
const _fcc957be = () => interopDefault(import('../pages/services/index.vue' /* webpackChunkName: "pages/services/index" */))
const _04efe362 = () => interopDefault(import('../pages/signin.vue' /* webpackChunkName: "pages/signin" */))
const _198639d8 = () => interopDefault(import('../pages/signup.vue' /* webpackChunkName: "pages/signup" */))
const _5cc8497c = () => interopDefault(import('../pages/index.vue' /* webpackChunkName: "pages/index" */))

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
    component: _1d06589e,
    name: "about-us"
  }, {
    path: "/appointment",
    component: _a62444e2,
    name: "appointment"
  }, {
    path: "/contacts",
    component: _b6c3acda,
    name: "contacts"
  }, {
    path: "/services",
    component: _fcc957be,
    name: "services"
  }, {
    path: "/signin",
    component: _04efe362,
    name: "signin"
  }, {
    path: "/signup",
    component: _198639d8,
    name: "signup"
  }, {
    path: "/",
    component: _5cc8497c,
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
