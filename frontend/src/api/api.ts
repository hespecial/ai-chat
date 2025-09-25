import http from '@/utils/http'
import * as types from './types'

export default {
  getCharacterById: (id: number): Promise<types.GetCharacterByIdResponse> => http.get(`/v1/character/${id}`),
  getCharacters: (): Promise<types.GetCharactersResponse> => http.get(`/v1/characters`),
  createChat: (params: types.CreateChatReq): Promise<types.CreateChatResponse> => http.post(`/v1/chat`, params),
  getChatHistory: (id: number): Promise<types.GetChatHistoryResponse> => http.get(`/v1/chat/${id}`),
  truncateChat: (id: number): Promise<object> => http.delete(`/v1/chat/${id}`),
}
