<template>
  <div
    class="min-h-screen bg-gradient-to-br from-slate-50 via-white to-slate-100 text-slate-800 dark:from-slate-900 dark:via-slate-950 dark:to-black dark:text-slate-100"
  >
    <header
      class="sticky top-0 z-30 backdrop-blur supports-[backdrop-filter]:bg-white/60 dark:supports-[backdrop-filter]:bg-slate-900/60 border-b border-slate-200/70 dark:border-slate-800/70"
    >
      <div class="container mx-auto px-4">
        <div class="flex h-16 items-center justify-between">
          <RouterLink to="/" class="flex items-center gap-3 group">
            <span
              class="inline-flex h-9 w-9 items-center justify-center rounded-xl bg-gradient-to-br from-brand-500 to-brand-600 text-white shadow-md group-hover:shadow-lg transition-shadow"
            >
              <svg
                width="18"
                height="18"
                viewBox="0 0 24 24"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M4 12c0-4.418 3.582-8 8-8 1.657 0 3 1.343 3 3s-1.343 3-3 3-3 1.343-3 3 1.343 3 3 3 3 1.343 3 3-1.343 3-3 3c-4.418 0-8-3.582-8-8Z"
                  fill="currentColor"
                />
              </svg>
            </span>
            <div>
              <div class="text-lg font-semibold leading-5 tracking-tight">AI Chat</div>
              <div class="text-[11px] text-slate-500 -mt-0.5">Role-play • Voice Chat</div>
            </div>
          </RouterLink>

          <nav class="hidden md:flex items-center gap-6 text-sm">
            <RouterLink
              to="/"
              class="text-slate-600 hover:text-slate-900 dark:text-slate-300 dark:hover:text-white transition-colors"
              >首页</RouterLink
            >
          </nav>
        </div>
      </div>
    </header>

    <main class="min-h-[calc(100vh-6rem)]">
      <slot />
    </main>

    <footer class="border-t border-slate-200/70 dark:border-slate-800/70">
      <div
        class="container mx-auto px-4 py-8 text-xs text-slate-500 flex flex-col sm:flex-row items-center justify-between gap-3"
      >
        <p>© {{ year }} 角色语聊 AI · 前端演示版</p>
        <p class="opacity-80">后端接口预留：/api/characters · /api/chat</p>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
const year = new Date().getFullYear()
</script>
