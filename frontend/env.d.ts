// env.d.ts
/// <reference types="vite/client" />

interface ImportMetaEnv {
  /**
   * API基础地址
   * 开发环境: http://localhost:3000
   * 生产环境: https://api.yourdomain.com
   * 测试环境: https://staging-api.yourdomain.com
   */
  readonly VITE_API_URL: string

  /**
   * 应用名称
   */
  readonly VITE_APP_TITLE: string

  /**
   * 应用版本
   */
  readonly VITE_APP_VERSION: string

  /**
   * 是否启用调试模式
   * @default 'false'
   */
  readonly VITE_DEBUG?: string

  /**
   * 其他自定义环境变量...
   */
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
