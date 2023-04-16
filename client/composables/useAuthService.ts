import { User } from "~~/models/auth"
import { useAuthStore } from "~~/store/authStore"
import api from "./api"

export const useAuthService = () => {

    const authStore = useAuthStore()

    const access_token = useCookie('jack_access_token')
    const refresh_token = useCookie('jack_refresh_token')

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
            const { data } = await api.get<User>("/auth/profile")

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
            })

            await getProfile()

            console.log("update")
        } catch (error) {
            authStore.authError()
        }
    }

    const refresh = async () => {
        try {
            await api.post("/auth/refresh", {}, {
                withCredentials: true,
                headers: {
                    Authorization: 'Bearer ' + refresh_token
                }
            })
        } catch (error) {
            authStore.authError()
        }
    }

    const uploadUserLogo = async (logo: File) => {
        try {
            const formData = new FormData()

            formData.append("logo", logo)

            await api.put("/auth/logo", formData)

            const { data } = await api.get<User>("/auth/profile")

            authStore.authenticateUser(data)
        } catch (error) {
            authStore.authError()
          }
    }

    return {
        signUp,
        signIn,
        getProfile,
        updateRole,
        authStore,
        refresh,
        uploadUserLogo
    }
}