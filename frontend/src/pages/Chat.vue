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
            <!--            <span class="hidden sm:inline">语音识别：</span>-->
            <!--            <span :class="recognizing ? 'text-brand-600' : ''">{{-->
            <!--              recognizing ? '进行中' : '未开启'-->
            <!--            }}</span>-->
            <button
              @click="enableVoice"
              :class="[
                'inline-flex h-10 items-center rounded-xl px-4 text-sm font-medium text-white',
                hasVoice ? 'bg-green-400 hover:bg-green-500' : 'bg-gray-400 hover:bg-gray-500',
              ]"
            >
              语音{{ hasVoice ? '已开启' : '已关闭' }}
            </button>
            <button
              class="inline-flex h-10 items-center rounded-xl bg-gray-400 px-4 text-sm font-medium text-white hover:bg-red-300"
              @click="truncateChat"
            >
              清空
            </button>
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
            <!-- 停止播放按钮，只在最后一条消息且消息数量大于1时显示 -->
            <button
              v-if="
                stopButtonVisibility &&
                i === messages.length - 1 &&
                messages.length > 1 &&
                m.role !== 'user' &&
                hasVoice
              "
              class="ml-2 mt-auto h-6 w-6 rounded-full bg-red-500 text-white text-xs flex items-center justify-center hover:bg-red-600"
              @click="stopAudioPlay"
            >
              <!-- ⏸ -->
              ⏹
            </button>
          </div>
          <div v-if="loading" class="text-xs text-slate-500">对方正在思考中…</div>
        </div>

        <form @submit.prevent="send" class="mt-4 flex items-end gap-3">
          <textarea
            v-model="input"
            rows="1"
            placeholder="输入内容，或使用麦克风…"
            @keydown="onKeydown"
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
const id = Number(route.params.id)
const character = ref<types.Character | null>(null)
const messages = ref<types.ChatMessage[]>([])
const input = ref('')
const loading = ref(false)
const initials = ref('')
const hasVoice = ref(true)
const stopButtonVisibility = ref(false)
// 音频对象引用
let currentAudio: HTMLAudioElement | null = null

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

function enableVoice() {
  hasVoice.value = !hasVoice.value
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

function onKeydown(e: KeyboardEvent) {
  if (e.key === 'Enter' && (e.ctrlKey || e.metaKey)) {
    e.preventDefault()
    input.value += '\n'
  } else if (e.key === 'Enter') {
    e.preventDefault()
    send()
  }
}

async function send() {
  const text = input.value.trim()
  if (!text) return
  input.value = ''
  messages.value.push({
    id: 0,
    role: 'user',
    content: text,
    created: Math.floor(Date.now() / 1000),
  })

  loading.value = true
  const reply = await api.createChat({ characterId: id, content: text })
  if (hasVoice.value) {
    try {
      const audioData = await api.getVoiceWave(reply.id)

      // 停止当前正在播放的音频
      if (currentAudio) {
        stopAudioPlay()
      }

      // 创建Blob对象并播放
      const blob = new Blob([audioData], { type: 'audio/mpeg' })
      const audioUrl = URL.createObjectURL(blob)

      currentAudio = new Audio(audioUrl)
      await currentAudio.play()

      // 音频播放结束后清理资源
      currentAudio.onended = () => {
        stopButtonVisibility.value = false
        cleanupAudio()
      }
    } catch (error) {
      console.error('播放音频失败:', error)
    }
  }

  messages.value.push({
    id: reply.id,
    role: 'assistant',
    content: reply.content,
    created: Math.floor(Date.now() / 1000),
  })

  loading.value = false
  stopButtonVisibility.value = true
  // if (hasVoice.value) {
  //   speak(reply.content)
  // }
}

function back() {
  history.length > 1 ? history.back() : (window.location.href = '/')
}

function stopAudioPlay() {
  stopButtonVisibility.value = false
  // 停止音频播放并清理资源
  if (currentAudio) {
    currentAudio.pause()
    cleanupAudio()
  }
}

// 清理音频资源
function cleanupAudio() {
  if (currentAudio) {
    if (currentAudio.src) {
      URL.revokeObjectURL(currentAudio.src)
    }
    currentAudio.onended = null
    currentAudio = null
  }
}

async function truncateChat() {
  messages.value.splice(1)
  await api.truncateChat(id)
}

watch(
  () => character.value?.id,
  () => initRecognition(),
)

onMounted(async () => {
  character.value = await api.getCharacterById(id)
  messages.value.push({
    id: 0,
    role: 'assistant',
    content: character.value.greeting
      ? character.value.greeting
      : `你好，我是${character.value.name}，很高兴与你交流。`,
    created: 0,
  })
  const resp = await api.getChatHistory(id)
  initials.value = initialsOf(character.value?.name)
  messages.value.push(...resp.histories)
  initRecognition()
})

function format(t: string) {
  return t.replace(/\n/g, '<br />')
}
</script>
