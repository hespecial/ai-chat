export interface ChatMessage {
  role: 'user' | 'assistant'
  content: string
}

export interface CreateChatReq {
  characterId: number
  content: string
}

export interface CreateChatResponse {
  content: string;
}
