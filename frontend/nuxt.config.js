export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: "frontend",
    htmlAttrs: {
      lang: "en",
    },
    meta: [
      { charset: "utf-8" },
      {
        name: "viewport",
        content: "width=device-width, initial-scale=1",
      },
      { hid: "description", name: "description", content: "" },
      { name: "format-detection", content: "telephone=no" },
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }],
  },

  ssr: false,
  mode: "spa",

  target: "server",
  server: {
    host: "0.0.0.0",
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: ["@/assets/scss/app.scss"],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    { src: "~/plugins/vuelidate" },
    { src: "~/plugins/repository" },
    { src: "~/plugins/toast" },
  ],
  axios: {
    proxy: true,
  },
  proxy: {
    "/api": {
      target:
        process.env.NODE_ENV === "development" ? "http://localhost:8080" : "",
      pathRewrite: { "^/api": "/" },
    },
  },

  auth: {
    redirect: {
      logout: "/signin",
      login: "/signin",
      home: "/",
      callback: false,
    },
    strategies: {
      local: {
        token: {
          property: "access_token",
          global: true,
          maxAge: 86400,
        },
        autoFetchUser: false,
        endpoints: {
          login: {
            url: "api/login",
            method: "post",
            propertyName: "access_token",
          },
          user: false,
        },
        autoLogout: true,
      },
    },
  },

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/bootstrap
    "bootstrap-vue/nuxt",
    "@nuxtjs/auth",
    "@nuxtjs/axios",
  ],

  env: {
    development: {
      BASE_URL: "http://localhost:8080",
    },
    production: {
      BASE_URL: "https://apartchain.io",
    },
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {},
};
