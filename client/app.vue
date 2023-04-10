<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>

<script lang="ts" setup>
const { authStore, getProfile } = useAuthService()

const access_token = useCookie('jack_access_token')
const refresh_token = useCookie('jack_refresh_token')

onMounted(async () => {

    if (access_token.value && refresh_token.value) {
        if (authStore.status === AuthStatus.AUTHENTICATED) {
            return
        }

        await getProfile()

        return
    }

    authStore.unaunthenticated()

    return
})
</script>

<style>
.page-enter-active,
.page-leave-active {
  transition: all 0.4s;
}
.page-enter-from,
.page-leave-to {
  opacity: 0;
}
</style>