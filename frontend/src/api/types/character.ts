export interface Character {
  id: number
  name: string
  subtitle?: string
  description: string
  tags?: string
  language?: string // e.g., 'zh-CN' | 'en-US'
  greeting?: string
  skills?: Skill[]
}

export interface GetCharactersResponse {
  list: Character[]
}

export interface GetCharacterByIdResponse extends Character {}

export interface Skill{
  id: number
  name: string
  description: string
  sufPath: string
}

export interface GetCharacterSkillsResponse {
  skills: Skill[]
}
