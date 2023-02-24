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

            console.log(data)

            authStore.authenticateUser(data)
        } catch (error) {
            authStore.authError()
        }
    }

    const updateRole = async (roleId: number) => {
        try {
            console.log(roleId)
            await api.put("/auth/update", {
                ...authStore.user,
                role_id: roleId,
                role: {
                    id: roleId
                }
            }, {
                headers: {
                    "Authorization": `Bearer ${access_token.value}`
                }
            })

            await getProfile()

            console.log("update")
        } catch (error) {
            authStore.authError()
        }
    }

    return {
        signUp,
        signIn,
        getProfile,
        updateRole,
        authStore
    }
}