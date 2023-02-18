import { User } from "~~/models/auth"
import { useAuthStore } from "~~/store/authStore"
import api from "./api"

export const useAuthService = () => {

    const authStore = useAuthStore()

    const access_token = useCookie('jack_access_token')

    const signUp = async (userData: Object) => {
        try {
            const { data } = await api.post<User>("/auth/signUp", userData, {
                withCredentials: true
            })

            authStore.authenticateUser(data)
        } catch (error) {
            authStore.authError()
        }
    }

    const signIn = async (userData: Object) => {
        try {
            const { data } = await api.post<User>("/auth/signIn", userData, {
                withCredentials: true
            })

            authStore.authenticateUser(data)
        } catch (error) {
            authStore.authError()
        }
    }

    const getProfile = async () => {
        try {
            const { data } = await api.get<User>("/auth/profile", {
                headers: {
                    "Authorization": `Bearer ${access_token.value}`
                }
            })

            authStore.authenticateUser(data)
        } catch (error) {
            console.log(error)
            authStore.authError()
        }
    }

    return {
        signUp,
        signIn,
        getProfile,
        authStore
    }
}