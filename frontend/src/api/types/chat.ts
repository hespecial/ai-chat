export interface ChatMessage {
  id: number
  role: 'user' | 'assistant'
  content: string
  created: number
}

export interface CreateChatReq {
  characterId: number
  content: string
}

export interface CreateChatResponse {
  id: number
  content: string
}

export interface GetChatHistoryResponse {
  histories: ChatMessage[]
}
