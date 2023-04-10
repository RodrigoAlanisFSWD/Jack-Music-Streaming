<script lang="ts" setup>
const { authStore, signUp } = useAuthService()

const router = useRouter()

const handleSubmit = async (userData: Object) => {
    try {
        await signUp(userData)

        router.push("/auth/selectRole")
    } catch (error) {
        return
    }
}
</script>

<template>
    <auth-sign-up-form @submit="handleSubmit"></auth-sign-up-form>
</template>