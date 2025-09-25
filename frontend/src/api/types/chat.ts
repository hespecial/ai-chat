export interface ChatMessage {
  role: 'user' | 'assistant'
  content: string
  created: number
}

export interface CreateChatReq {
  characterId: number
  content: string
}

export interface CreateChatResponse {
  content: string
}

export interface GetChatHistoryResponse {
  histories: ChatMessage[]
}
