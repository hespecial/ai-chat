export interface Character {
  id: number
  name: string
  subtitle?: string
  description: string
  tags?: string
  language?: string // e.g., 'zh-CN' | 'en-US'
  greeting?: string
}

export interface GetCharactersResponse {
  list: Character[]
}

export interface GetCharacterByIdResponse extends Character {}
