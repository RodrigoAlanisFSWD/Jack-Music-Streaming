<script lang="ts" setup>
definePageMeta({
    middleware: ['auth']
})

const { updateRole } = useAuthService()

const router = useRouter()

const selectRole = async (id: number) => {
    try {
        await updateRole(id)

        router.push("/profile")
    } catch (error) {
        return
    }
}   
</script>

<template>
    <div class="w-full flex flex-col items-center">
        <div class="flex flex-col text-white p-8">
            <h2 class="text-4xl font-bold mb-3">
                Select A Role
            </h2>
        </div>
        <div class="flex">
            <div class="flex justify-around items-center h-[530px]">
                <div class="text-white flex flex-col w-[325px] p-5 border-r border-white">
                    <h2 class="text-2xl mb-2">
                        Listener
                    </h2>
                    <p class="text-sm mb-3">
                        Your only want to listen to music and follow your favorite artists
                    </p>
                    <app-button @click="selectRole(1)" type="primary-outlined">I Want To Listen</app-button>
                    <h3 class="text-sm text-primary mt-5">
                        What’s included
                    </h3>
                    <div class="flex items-start my-2">
                        <i class="uil uil-check-circle text-green-500 text-2xl mr-3"></i>
                        <p class="mt-[3px]">
                            Listen All The Music You Want
                        </p>
                    </div>
                    <div class="flex items-start my-2">
                        <i class="uil uil-check-circle text-green-500 text-2xl mr-3"></i>
                        <p class="mt-[3px]">
                            Follow All The Artists You Love!
                        </p>
                    </div>
                    <div class="flex items-start my-2">
                        <i class="uil uil-check-circle text-green-500 text-2xl mr-3"></i>
                        <p class="mt-[3px]">
                            Create Playlists Or Save Public Playlists
                        </p>
                    </div>
                    <div class="flex items-start my-2">
                        <i class="uil uil-check-circle text-green-500 text-2xl mr-3"></i>
                        <p class="mt-[3px]">
                            Download The Free Songs
                        </p>
                    </div>
                    <div class="flex items-start my-2">
                        <i class="uil uil-times-circle text-white text-2xl mr-3"></i>
                        <p class="mt-[3px]">
                            Access To Premium Songs
                        </p>
                    </div>
                    <div class="flex items-start my-2">
                        <i class="uil uil-times-circle text-white text-2xl mr-3"></i>
                        <p class="mt-[3px]">
                            Upload Music
                        </p>
                    </div>
                </div>
            </div>
            <div class="flex justify-around items-start h-[530px]">
                <div class="text-white flex flex-col w-[325px] p-5">
                    <h2 class="text-2xl mb-2">
                        Artist
                    </h2>
                    <p class="text-sm mb-3">
                        Your want to share your music, alone or with your band
                    </p>
                    <app-button @click="selectRole(2)" type="primary-outlined">I'm An Artist</app-button>
                    <h3 class="text-sm text-primary mt-5">
                        What’s included
                    </h3>
                    <div class="flex items-start my-2">
                        <i class="uil uil-check-circle text-green-500 text-2xl mr-3"></i>
                        <p class="mt-[3px]">
                            All The Listener Permissions
                        </p>
                    </div>
                    <div class="flex items-start my-2">
                        <i class="uil uil-check-circle text-green-500 text-2xl mr-3"></i>
                        <p class="mt-[3px]">
                            Create A Band With Other Artists
                        </p>
                    </div>
                    <div class="flex items-start my-2">
                        <i class="uil uil-check-circle text-green-500 text-2xl mr-3"></i>
                        <p class="mt-[3px]">
                            Upload Music All The Music You Want!
                        </p>
                    </div>
                    <div class="flex items-start my-2">
                        <i class="uil uil-check-circle text-green-500 text-2xl mr-3"></i>
                        <p class="mt-[3px]">
                            Create Ads (With Cost)
                        </p>
                    </div>
                    <div class="flex items-start my-2">
                        <i class="uil uil-check-circle text-green-500 text-2xl mr-3"></i>
                        <p class="mt-[3px]">
                            Create Albums
                        </p>
                    </div>
                    <div class="flex items-start my-2">
                        <i class="uil uil-exclamation-octagon text-red-500 text-2xl mr-3"></i>
                        <p class="mt-[3px]">
                            You Need To Pass A Verification
                        </p>
                    </div>
                </div>
            </div>
        </div>


    </div>
</template>