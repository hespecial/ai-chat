import type { Character, ChatMessage, ChatPayload } from './client'

export const characters: Character[] = [
  {
    id: 'harry-potter',
    name: '哈利·波特',
    subtitle: '霍格沃茨的小巫师',
    description: '勇敢、忠诚、富有正义感。与你聊聊魔法世界、魁地奇与友谊。',
    tags: ['魔法', '冒险', '友情'],
    language: 'zh-CN',
    greeting: '嗨！我是哈利·波特。准备好一起探索魔法世界了吗？',
  },
  {
    id: 'socrates',
    name: '苏格拉底',
    subtitle: '古希腊哲学家',
    description: '以“产婆术”启发式提问著称。善于引导你思考真理、善良与灵魂。',
    tags: ['哲学', '思辨', '伦理'],
    language: 'zh-CN',
    greeting: '你好，我是苏格拉底。我们不如以问题开始吧？',
  },
  {
    id: 'sherlock-holmes',
    name: '夏洛克·福尔摩斯',
    subtitle: '天才侦探',
    description: '观察入微、逻辑缜密。擅长推理与分析，帮助你理清线索与真相。',
    tags: ['推理', '逻辑', '悬疑'],
    language: 'zh-CN',
    greeting: '福尔摩斯在此。请告诉我你观察到的全部细节。',
  },
  {
    id: 'albert-einstein',
    name: '阿尔伯特·爱因斯坦',
    subtitle: '物理学家',
    description: '想象力比知识更重要。一起聊聊相对论、科学与好奇心。',
    tags: ['科学', '相对论', '好奇心'],
    language: 'zh-CN',
    greeting: '你好，我是爱因斯坦。让我们一起思考宇宙的奥秘。',
  },
]

export function mockGetCharacters(query?: string): Promise<Character[]> {
  if (!query) return Promise.resolve(characters)
  const q = query.toLowerCase()
  return Promise.resolve(
    characters.filter((c) =>
      [c.name, c.subtitle, c.description, ...(c.tags || [])].join(' ').toLowerCase().includes(q),
    ),
  )
}

function styleReply(characterId: string, user: string): string {
  switch (characterId) {
    case 'harry-potter':
      return `咒语闪耀！你提到“${user}”。在霍格沃茨，我们会这样看待它：\n\n· 保持勇敢与善良\n· 相信伙伴的力量\n· 别忘了享受魁地奇！`
    case 'socrates':
      return `关于“${user}”，我们不妨先提出几个问题：\n1. 你为何这样认为？\n2. 这种看法是否始终成立？\n3. 若换个角度，结论会变化吗？\n\n请从你的直觉出发，再给出一个理由。`
    case 'sherlock-holmes':
      return `当我们谈到“${user}”时，重要的是证据与逻辑：\n- 先观察细节\n- 再提出假设\n- 然后验证与排除\n\n请告诉我更多线索，我来协助推理。`
    case 'albert-einstein':
      return `对“${user}”的思考很有趣。想象力将引导我们：\n· 提出大胆的设想\n· 简化问题的本质\n· 以实验或思想实验检验\n\n别害怕提出看似异想天开的观点。`
    default:
      return `关于“${user}”，我很乐意继续交流。`
  }
}

export async function mockSendChat(payload: ChatPayload): Promise<ChatMessage> {
  const lastUser = [...payload.messages].reverse().find((m) => m.role === 'user')
  const content = styleReply(payload.characterId, lastUser?.content || '')
  await new Promise((r) => setTimeout(r, 600))
  return { role: 'assistant', content }
}
