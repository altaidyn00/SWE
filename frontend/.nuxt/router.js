import Vue from 'vue'
import Router from 'vue-router'
import { normalizeURL, decode } from 'ufo'
import { interopDefault } from './utils'
import scrollBehavior from './router.scrollBehavior.js'

const _438590f6 = () => interopDefault(import('../pages/about-us.vue' /* webpackChunkName: "pages/about-us" */))
const _c8a8b392 = () => interopDefault(import('../pages/appointment.vue' /* webpackChunkName: "pages/appointment" */))
const _69c53c2a = () => interopDefault(import('../pages/contacts.vue' /* webpackChunkName: "pages/contacts" */))
const _18f4ea79 = () => interopDefault(import('../pages/services/index.vue' /* webpackChunkName: "pages/services/index" */))
const _23dd5c8c = () => interopDefault(import('../pages/signin.vue' /* webpackChunkName: "pages/signin" */))
const _1c317c2c = () => interopDefault(import('../pages/index.vue' /* webpackChunkName: "pages/index" */))

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
    component: _438590f6,
    name: "about-us"
  }, {
    path: "/appointment",
    component: _c8a8b392,
    name: "appointment"
  }, {
    path: "/contacts",
    component: _69c53c2a,
    name: "contacts"
  }, {
    path: "/services",
    component: _18f4ea79,
    name: "services"
  }, {
    path: "/signin",
    component: _23dd5c8c,
    name: "signin"
  }, {
    path: "/",
    component: _1c317c2c,
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
