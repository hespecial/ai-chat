export interface Character {
  id: number
  name: string
  subtitle?: string
  description: string
  tags?: string[]
  language?: string // e.g., 'zh-CN' | 'en-US'
  greeting?: string
}

export interface ChatMessage {
  role: 'user' | 'assistant'
  content: string
}

export interface ChatPayload {
  characterId: string
  messages: ChatMessage[]
}

const baseURL = (import.meta as any).env?.VITE_API_URL
  ? `${(import.meta as any).env.VITE_API_URL}/api`
  : '/api'

async function tryFetch(input: RequestInfo | URL, init?: RequestInit) {
  try {
    const res = await fetch(input, init)
    if (!res.ok) throw new Error('bad status')
    return await res.json()
  } catch {
    return null
  }
}

export async function getCharacters(id?: number): Promise<Character[]> {
  const q = id ? `?id=${id}` : ''
  const data = await tryFetch(`${baseURL}/characters${q}`)
  if (data) return data as Character[]
  const { mockGetCharacters } = await import('./mockServer')
  return mockGetCharacters(id)
}

export async function sendChat(payload: ChatPayload): Promise<ChatMessage> {
  const data = await tryFetch(`${baseURL}/chat`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  })
  if (data) return data as ChatMessage
  const { mockSendChat } = await import('./mockServer')
  return mockSendChat(payload)
}
