<template>
  <div
    v-if="visible"
    class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50"
    @click.self="$emit('close')"
  >
    <div class="w-full max-w-md rounded-lg bg-white p-6 shadow-lg dark:bg-slate-800">
      <div class="flex items-center justify-between">
        <h3 class="text-lg font-medium text-slate-900 dark:text-white">角色技能</h3>
        <button @click="$emit('close')" class="text-slate-400 hover:text-slate-500">
          <span class="i-lucide:x h-6 w-6"></span>
        </button>
      </div>
      <ul class="mt-4 divide-y divide-slate-200 dark:divide-slate-700">
        <li
          v-for="skill in skills"
          :key="skill.id"
          class="group relative cursor-pointer py-3 transition-colors hover:bg-slate-100 dark:hover:bg-slate-700"
          @click="$emit('select-skill', skill)"
        >
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium text-slate-900 dark:text-white">{{ skill.name }}</span>
            <span class="i-lucide:help-circle h-5 w-5 text-slate-400"></span>
          </div>
          <div
            class="absolute bottom-full left-1/2 z-20 mb-2 -translate-x-1/2 transform whitespace-nowrap rounded-md bg-slate-900 px-3 py-2 text-sm text-white opacity-0 group-hover:opacity-100"
          >
            {{ skill.description }}
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Skill } from '@/api/types/character'

defineProps<{
  skills: Skill[]
  visible: boolean
}>()

const emit = defineEmits<{
  (e: 'close'): void
  (e: 'select-skill', skill: Skill): void
}>()
</script>