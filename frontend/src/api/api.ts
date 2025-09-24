import http from '@/utils/http'
import {
  type GetCharacterByIdResponse,
  type GetCharactersResponse,
} from './types/character'

export default {
  getCharacterById: (id: number): Promise<GetCharacterByIdResponse> => http.get(`/v1/character/${id}`),
  getCharacters: (): Promise<GetCharactersResponse> => http.get(`/v1/characters`),
}
