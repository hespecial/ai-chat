import http from '@/utils/http'
import * as types from './types'

export default {
  getCharacterById: (id: number): Promise<types.GetCharacterByIdResponse> => http.get(`/v1/character/${id}`),
  getCharacters: (): Promise<types.GetCharactersResponse> => http.get(`/v1/characters`),
  createChat: (params: types.CreateChatReq): Promise<types.CreateChatResponse> => http.post(`/v1/chat`, params),
  getChatHistory: (id: number): Promise<types.GetChatHistoryResponse> => http.get(`/v1/chat/${id}`),
  truncateChat: (id: number): Promise<object> => http.delete(`/v1/chat/${id}`),
  getVoiceWave: (id: number): Promise<ArrayBuffer> => http.getBinary(`/v1/chat/voice?chatHistoryId=${id}`),
  getCharacterSkills: (id: number): Promise<types.GetCharacterSkillsResponse> => http.get(`/v1/character/${id}/skill`),
  useSkill: (sufPath: string,params: types.SkillReq): Promise<types.SkillResponse> => http.post(`/v1/skill/${sufPath}`,params)
}
