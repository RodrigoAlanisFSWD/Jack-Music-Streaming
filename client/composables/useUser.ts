import { useAuthStore } from "~~/store/authStore"

export const useUser = () => {
    const authStore = useAuthStore()

    return authStore.user
}