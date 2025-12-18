<script setup>
useHead({
  meta: [
    { name: 'viewport', content: 'width=device-width, initial-scale=1' }
  ],
  link: [
    { rel: 'icon', href: '/favicon.ico' }
  ],
  htmlAttrs: {
    lang: 'en'
  }
})

const title = 'Chainstamping • Git-based timestamping on Ethereum'
const description = 'Easily timestamp your Git commits on the Ethereum blockchain with Chainstamping.'

useSeoMeta({
  title,
  description,
  ogTitle: title,
  ogDescription: description,
  //ogImage: 'https://ui.nuxt.com/assets/templates/nuxt/starter-light.png',
  //twitterImage: 'https://ui.nuxt.com/assets/templates/nuxt/starter-light.png',
  //twitterCard: 'summary_large_image'
})
</script>

<template>
  <UApp>
    <UMain>
      <UContainer>
        <div class="flex items-center justify-center min-h-screen">
          <NuxtPage />
        </div>
      </UContainer>
    </UMain>
  </UApp>
</template>
