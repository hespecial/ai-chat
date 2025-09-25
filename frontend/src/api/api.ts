import http from '@/utils/http'
import * as types from './types'
import type { CreateChatReq } from './types'

export default {
  getCharacterById: (id: number): Promise<types.GetCharacterByIdResponse> => http.get(`/v1/character/${id}`),
  getCharacters: (): Promise<types.GetCharactersResponse> => http.get(`/v1/characters`),
  createChat: (params: CreateChatReq): Promise<types.CreateChatResponse> => http.post(`/v1/chat`, params),
}
