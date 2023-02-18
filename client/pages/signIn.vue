<script lang="ts" setup>
import { useAuthService } from '~~/hooks/authService'

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
    <div class="grid grid-cols-[1fr_650px] w-screen h-[calc(100vh-70px)]">
        <div class="background w-full bg-white">

        </div>
        <auth-sign-in-form @submit="handleSubmit"></auth-sign-in-form>
    </div>
</template>

<style>
    .background {
        background-image: url("../assets/img/guitar.png");
        background-size: cover;
    }
</style>