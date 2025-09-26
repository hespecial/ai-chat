<template>
  <Layout>
    <section class="relative">
      <div
        aria-hidden
        class="pointer-events-none absolute -top-24 -left-24 h-64 w-64 rounded-full bg-brand-500/20 blur-3xl"
      />
      <div
        aria-hidden
        class="pointer-events-none absolute -bottom-24 -left-24 h-64 w-64 rounded-full bg-brand-400/20 blur-3xl"
      />

      <div class="container mx-auto px-4 py-12 md:py-16">
        <div class="mx-auto max-w-3xl text-center">
          <h1
            class="text-3xl md:text-5xl font-extrabold tracking-tight bg-gradient-to-b from-slate-900 to-slate-700 dark:from-white dark:to-slate-300 bg-clip-text text-transparent"
          >
            AI Chat
          </h1>
          <p class="mt-4 text-slate-600 dark:text-slate-300">
            搜索并选择感兴趣的角色，开启一段由 AI 驱动的角色扮演之旅!
          </p>
          <div class="mt-8">
            <SearchBar v-model="query" @search="doSearch" />
          </div>

          <div class="mt-6 flex flex-wrap items-center justify-center gap-2 text-xs">
            <span class="text-slate-500">热门：</span>
            <button
              v-for="t in popularTags"
              :key="t"
              @click="applyTag(t)"
              class="rounded-full border border-slate-200 bg-white/60 px-3 py-1 text-slate-600 hover:border-brand-500 hover:text-brand-700 dark:border-slate-700 dark:bg-slate-900/60 dark:text-slate-300"
            >
              #{{ t }}
            </button>
          </div>
        </div>

        <div class="mt-10 grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
          <CharacterCard v-for="c in filtered" :key="c.id" :character="c" @start="goChat(c)" />
        </div>

        <div v-if="!filtered.length" class="mt-16 text-center text-slate-500">
          未找到相关角色，换个关键词试试～
        </div>
      </div>
    </section>
  </Layout>
</template>

<script setup lang="ts">
import Layout from '@/components/Layout/Layout.vue'
import SearchBar from '@/components/SearchBar.vue'
import CharacterCard from '@/components/CharacterCard.vue'
import { onMounted, ref, computed } from 'vue'
import type { GetCharactersResponse, Character } from '@/api/types/character'
import api from '@/api/api'

const query = ref('')
const all = ref<GetCharactersResponse>({ list: [] })
const popularTags = ['谋略', '科学', '历史', '文学', '女性']

const filtered = computed(() => {
  if (!query.value) return all.value.list
  const q = query.value.toLowerCase()
  return all.value.list.filter((c) =>
    [c.name, c.subtitle, c.description, c.tags].join(' ').toLowerCase().includes(q),
  )
})

function doSearch() {
  // already reactive via computed
}

function applyTag(t: string) {
  query.value = t
}

function goChat(c: Character) {
  window.location.href = `/chat/${encodeURIComponent(c.id)}`
}

onMounted(async () => {
  all.value = await api.getCharacters()
  // console.log(all.value.list)
  // all.value = await getCharacters()
  // console.log(all.value)
})
</script>
