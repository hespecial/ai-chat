<template>
  <div
    class="group rounded-2xl border border-slate-200 bg-white/70 p-5 shadow-sm backdrop-blur transition hover:-translate-y-0.5 hover:shadow-md dark:border-slate-700 dark:bg-slate-900/60"
  >
    <div class="flex items-start gap-4">
      <div
        class="flex h-12 w-12 shrink-0 items-center justify-center rounded-xl bg-gradient-to-br from-brand-500 to-brand-600 text-white font-semibold"
      >
        {{ initials }}
      </div>
      <div class="min-w-0 flex-1">
        <h3 class="truncate text-lg font-semibold text-slate-900 dark:text-white">
          {{ character.name }}
        </h3>
        <p class="truncate text-xs text-slate-500">{{ character.subtitle }}</p>
        <p class="mt-2 line-clamp-3 text-sm text-slate-600 dark:text-slate-300">
          {{ character.description }}
        </p>
        <div class="mt-2 flex flex-wrap gap-2">
          <span
            v-for="t in tags"
            :key="t"
            class="rounded-full border border-slate-200 bg-white/60 px-2 py-0.5 text-[11px] text-slate-500 dark:border-slate-700 dark:bg-slate-900/60"
            >#{{ t }}</span
          >
        </div>
      </div>
    </div>
    <div class="mt-4 flex items-center justify-between">
      <span class="text-xs text-slate-500">支持语音与文字</span>
      <button
        @click="$emit('start')"
        class="inline-flex items-center gap-2 rounded-xl bg-brand-600 px-3 py-1.5 text-sm font-medium text-white hover:bg-brand-700"
      >
        开始对话 <span aria-hidden>→</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Character } from '@/api/types/character'
import { computed } from 'vue'

const props = defineProps<{ character: Character }>()
const tags = computed(() => {
  // 反序列化 character.tags
  if (props.character.tags) {
    try {
      const tagArr: string[] = JSON.parse(props.character.tags)
      return tagArr
    } catch (error) {
      console.error('解析 tags 失败:', error)
    }
  }
})

function initialsOf(name?: string) {
  if (!name) return 'AI'
  const parts = name.replace(/\s+/g, ' ').trim().split(' ')
  if (/[\u4e00-\u9fa5]/.test(name)) return name.slice(0, 1)
  return parts
    .slice(0, 2)
    .map((p) => p[0])
    .join('')
    .toUpperCase()
}
const initials = initialsOf(props.character?.name)
</script>
