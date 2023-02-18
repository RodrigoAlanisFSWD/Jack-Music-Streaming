<script lang="ts" setup>
import { useAuthService } from '~~/hooks/authService';
const { authStore, signUp } = useAuthService()

const router = useRouter()

const handleSubmit = async (userData: Object) => {
    try {
        await signUp(userData)
        
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
        <auth-sign-up-form @submit="handleSubmit"></auth-sign-up-form>
    </div>
</template>

<style>
.background {
    background-image: url("../assets/img/guitar.png");
    background-size: cover;
}</style>