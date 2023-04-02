import { useAuthStore } from "~~/store/authStore"

export const useUser = () => {
    const authStore = useAuthStore()

    const { user } = storeToRefs(authStore)

    return user
}