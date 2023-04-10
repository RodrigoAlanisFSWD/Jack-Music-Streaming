import { AuthStatus } from "~~/store/authStore"

export default defineNuxtRouteMiddleware(async (to, from) => {
    const access_token = useCookie('jack_access_token')
    const refresh_token = useCookie('jack_refresh_token')

    const { authStore, getProfile } = useAuthService()

    if (access_token.value && refresh_token.value) {
        if (authStore.status === AuthStatus.AUTHENTICATED) {
            return
        }

        await getProfile()

        return
    }

    authStore.unaunthenticated

    return navigateTo("/auth/signIn")
})