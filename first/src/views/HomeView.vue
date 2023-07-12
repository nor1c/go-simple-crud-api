<template>
  <main>
    <TheWelcome />
  </main>
</template>

<script setup lang="ts">
import { useFetch } from '@vueuse/core'

const { data, error } = await useFetch('http://localhost:3000')
console.log('data:', data.value);
</script>