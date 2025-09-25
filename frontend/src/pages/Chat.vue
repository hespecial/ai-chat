<template>
  <Layout>
    <section class="container mx-auto px-4 py-6 md:py-10">
      <div class="mx-auto max-w-3xl">
        <button
          @click="back"
          class="mb-4 inline-flex items-center gap-2 text-sm text-slate-500 hover:text-slate-800 dark:hover:text-white"
        >
          <span class="i-lucide:chevron-left">←</span> 返回
        </button>

        <div
          class="flex items-center gap-4 rounded-2xl border border-slate-200 bg-white/70 p-4 backdrop-blur dark:border-slate-700 dark:bg-slate-900/60"
        >
          <div
            class="flex h-12 w-12 items-center justify-center rounded-xl bg-gradient-to-br from-brand-500 to-brand-600 text-white font-semibold"
          >
            {{ initials }}
          </div>
          <div>
            <div class="text-lg font-semibold">{{ character?.name || '未知角色' }}</div>
            <div class="text-xs text-slate-500">{{ character?.subtitle }}</div>
          </div>
          <div class="ml-auto flex items-center gap-3 text-xs text-slate-500">
            <span class="hidden sm:inline">语音识别：</span>
            <span :class="recognizing ? 'text-brand-600' : ''">{{
              recognizing ? '进行中' : '未开启'
            }}</span>
          </div>
        </div>

        <div
          class="mt-4 h-[55vh] overflow-y-auto rounded-2xl border border-slate-200 bg-white/70 p-4 backdrop-blur dark:border-slate-700 dark:bg-slate-900/60"
        >
          <div
            v-for="(m, i) in messages"
            :key="i"
            class="mb-4 flex"
            :class="m.role === 'user' ? 'justify-end' : 'justify-start'"
          >
            <div
              :class="
                m.role === 'user'
                  ? 'bg-brand-600 text-white'
                  : 'bg-slate-100 dark:bg-slate-800 text-slate-800 dark:text-slate-100'
              "
              class="max-w-[80%] rounded-2xl px-4 py-2 text-sm shadow"
            >
              <p v-html="format(m.content)"></p>
            </div>
          </div>
          <div v-if="loading" class="text-xs text-slate-500">对方正在思考中…</div>
        </div>

        <form @submit.prevent="send" class="mt-4 flex items-end gap-3">
          <textarea
            v-model="input"
            rows="1"
            placeholder="输入内容，或使用麦克风…"
            class="flex-1 resize-none rounded-xl border border-slate-300 bg-white/80 px-3 py-2 text-sm shadow-sm outline-none focus:border-brand-500 focus:ring-2 focus:ring-brand-200 dark:border-slate-700 dark:bg-slate-900/60"
          ></textarea>
          <button
            type="button"
            :disabled="!canVoice"
            @click="toggleVoice"
            class="h-10 w-10 rounded-xl border border-slate-300 bg-white/80 text-slate-700 hover:text-brand-700 hover:border-brand-500 disabled:opacity-50 dark:border-slate-700 dark:bg-slate-900/60"
          >
            🎤
          </button>
          <button
            type="submit"
            class="inline-flex h-10 items-center rounded-xl bg-brand-600 px-4 text-sm font-medium text-white hover:bg-brand-700 disabled:opacity-60"
          >
            发送
          </button>
        </form>
      </div>
    </section>
  </Layout>
</template>

<script setup lang="ts">
import Layout from '@/components/Layout/Layout.vue'
import { onMounted, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import * as types from '@/api/types'
import api from '@/api/api'

const route = useRoute()
const id = String(route.params.id || '')
const character = ref<types.Character | null>(null)
const messages = ref<types.ChatMessage[]>([])
const input = ref('')
const loading = ref(false)
const initials = ref('')

// speech
declare global {
  interface Window {
    webkitSpeechRecognition: any
  }
}
let recognition: any | null = null
const recognizing = ref(false)
const canVoice = 'webkitSpeechRecognition' in window || 'SpeechRecognition' in (window as any)

function initRecognition() {
  const Ctor = (window as any).SpeechRecognition || window.webkitSpeechRecognition
  if (!Ctor) return
  recognition = new Ctor()
  recognition.lang = character.value?.language || 'zh-CN'
  recognition.continuous = false
  recognition.interimResults = false
  recognition.onstart = () => (recognizing.value = true)
  recognition.onend = () => (recognizing.value = false)
  recognition.onerror = () => (recognizing.value = false)
  recognition.onresult = (e: any) => {
    const text = Array.from(e.results)
      .map((r: any) => r[0].transcript)
      .join(' ')
    input.value = text
  }
}

function toggleVoice() {
  if (!recognition) return
  if (recognizing.value) recognition.stop()
  else recognition.start()
}

function speak(text: string) {
  try {
    const utter = new SpeechSynthesisUtterance(text)
    utter.lang = character.value?.language || 'zh-CN'
    utter.rate = 1
    speechSynthesis.speak(utter)
  } catch {}
}

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

function persist() {
  if (!character.value) return
  localStorage.setItem(`chat:${character.value.id}`, JSON.stringify(messages.value))
}

async function loadHistory() {
  if (!character.value) return
  const raw = localStorage.getItem(`chat:${character.value.id}`)
  messages.value = raw
    ? JSON.parse(raw)
    : [
        {
          role: 'assistant',
          content:
            character.value.greeting || `你好，我是${character.value.name}，很高兴与你交流。`,
        },
      ]
}

async function send() {
  const text = input.value.trim()
  if (!text) return
  input.value = ''
  messages.value.push({ role: 'user', content: text })
  persist()

  loading.value = true
  // const reply = await sendChat({ characterId: id, messages: messages.value })
  const reply = await api.createChat({ characterId: Number(id), content: text })
  loading.value = false
  messages.value.push({role: 'assistant', content: reply.content})
  persist()
  // speak(reply.content)
}

function back() {
  history.length > 1 ? history.back() : (window.location.href = '/')
}

watch(
  () => character.value?.id,
  () => initRecognition(),
)

onMounted(async () => {
  character.value = await api.getCharacterById(Number(id))
  initials.value = initialsOf(character.value?.name)
  await loadHistory()
  initRecognition()
})

function format(t: string) {
  return t.replace(/\n/g, '<br />')
}
</script>
