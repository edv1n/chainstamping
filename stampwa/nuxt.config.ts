// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: [
    '@nuxt/eslint',
    '@nuxt/ui'
  ],

  devtools: {
    enabled: true
  },

  css: ['~/assets/css/main.css'],

  routeRules: {
    '/': { prerender: true }
  },

  compatibilityDate: '2025-01-15',

  eslint: {
    config: {
      stylistic: {
        commaDangle: 'never',
        braceStyle: '1tbs'
      }
    }
  },

  imports: {
    dirs: [
      // Scan top-level composables
      '~/composables',
      // ... or scan composables nested one level deep with a specific name and file extension
      '~/composables/*/index.{ts,js,mjs,mts}',
    ],
  },

  ssr: false,
})
