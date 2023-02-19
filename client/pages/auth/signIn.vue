<script lang="ts" setup>
const { authStore, signIn } = useAuthService()

const router = useRouter()

const handleSubmit = async (data: Object) => {
    try {
        await signIn(data)

        router.push("/profile")
    } catch (error) {
        return
    }
}
</script>

<template>
    <auth-sign-in-form @submit="handleSubmit"></auth-sign-in-form>
</template>