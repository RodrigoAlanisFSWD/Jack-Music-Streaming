// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: ['@nuxtjs/tailwindcss', '@pinia/nuxt'],
  app: {
    pageTransition: { name: 'page', mode: 'out-in' },
    head: {
      title: "Jack Music Streaming",
      link: [
        {
          rel: 'stylesheet',
          href: 'https://unicons.iconscout.com/release/v4.0.0/css/line.css'
        }
      ]
    }
  }
})