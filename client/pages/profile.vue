<template>
    <div>
        {{  user?.name  }}
    </div>
</template>

<script lang="ts" setup>
import { useUser } from '~~/hooks/user';

    definePageMeta({
        middleware: 'auth'
    })

    const user = useUser()
</script>