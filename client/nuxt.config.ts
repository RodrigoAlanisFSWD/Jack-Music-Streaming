// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: ['@nuxtjs/tailwindcss'],
  app: {
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