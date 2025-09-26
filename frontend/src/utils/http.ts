import axios, {
  type AxiosInstance,
  type AxiosRequestConfig,
  type AxiosResponse,
  type ResponseType,
} from 'axios';

// 定义接口返回的数据结构
interface ApiResponse<T = any> {
  code: number;
  message: string;
  data: T;
}

class HttpClient {
  private instance: AxiosInstance;

  constructor(baseURL: string) {
    this.instance = axios.create({
      baseURL,
      timeout: 10000, // 请求超时时间
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // 请求拦截器
    this.instance.interceptors.request.use(
      (config) => {
        // 在发送请求之前可以做一些处理，比如添加 token
        const token = localStorage.getItem('token');
        if (token) {
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      },
      (error) => {
        return Promise.reject(error);
      }
    );

    // 响应拦截器
    this.instance.interceptors.response.use(
      (response: AxiosResponse<ApiResponse>) => {
        // 对响应数据做一些处理
        if (response.data.code !== 200) {
          console.error('请求失败:', response.data.message);
          return Promise.reject(response.data.message);
        }
        return response.data.data; // 直接返回实际数据
      },
      (error) => {
        // 对错误响应做一些处理
        console.error('请求出错:', error.message);
        return Promise.reject(error);
      }
    );
  }

  // 封装 GET 请求
  public async get<T>(url: string, config?: AxiosRequestConfig): Promise<T> {
    return this.instance.get(url, config);
  }

  // 封装 POST 请求
  public async post<T>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T> {
    return this.instance.post(url, data, config);
  }

  // 封装 PUT 请求
  public async put<T>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T> {
    return this.instance.put(url, data, config);
  }

  // 封装 DELETE 请求
  public async delete<T>(url: string, config?: AxiosRequestConfig): Promise<T> {
    return this.instance.delete(url, config);
  }

  // 封装获取二进制数据的请求（不经过响应拦截器）
  public async getBinary(url: string, config?: AxiosRequestConfig): Promise<ArrayBuffer> {
    try {
      // 完整URL
      const fullUrl = `${baseURL}${url}`;

      const binaryConfig: AxiosRequestConfig = {
        ...config,
        responseType: 'arraybuffer' as ResponseType,
        headers: {
          ...config?.headers,
        },
      };

      // 创建新的axios实例，没有任何拦截器
      const binaryInstance = axios.create();

      // 直接使用新实例发起请求，完全绕过拦截器
      const response = await binaryInstance.get(fullUrl, binaryConfig);
      return response.data;
    } catch (error) {
      console.error('获取二进制数据失败:', error);
      throw error;
    }
  }
}

// 创建实例并导出
const baseURL = (import.meta as any).env?.VITE_API_URL ? `${(import.meta as any).env.VITE_API_URL}/api` : '/api';
const http = new HttpClient(baseURL);

export default http;
