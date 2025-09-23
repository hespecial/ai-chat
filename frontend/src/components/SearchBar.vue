<template>
  <div class="relative mx-auto max-w-2xl">
    <input
      :value="modelValue"
      @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
      @keyup.enter="$emit('search')"
      type="text"
      class="w-full rounded-2xl border border-slate-300 bg-white/80 px-5 py-3 text-sm shadow-sm outline-none focus:border-brand-500 focus:ring-2 focus:ring-brand-200 dark:border-slate-700 dark:bg-slate-900/60"
      placeholder="搜索角色名称、关键词…"
    />
    <button
      @click="$emit('search')"
      class="absolute right-2 top-1/2 -translate-y-1/2 rounded-xl bg-brand-600 px-3 py-1.5 text-xs font-medium text-white hover:bg-brand-700"
    >
      搜索
    </button>
  </div>
</template>

<script setup lang="ts">
defineProps<{ modelValue: string }>()
</script>
