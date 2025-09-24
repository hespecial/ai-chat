package middleware

import "net/http"

func Cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 设置跨域头
		w.Header().Set("Access-Control-Allow-Origin", "*")                                // 允许所有来源（生产环境建议指定具体域名）
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // 允许的 HTTP 方法
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")     // 允许的自定义头
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")                 // 允许前端访问的响应头
		w.Header().Set("Access-Control-Allow-Credentials", "true")                        // 是否允许携带 Cookie

		// 如果是 OPTIONS 请求，直接返回 204 状态码
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// 继续处理请求
		next(w, r)
	}
}
