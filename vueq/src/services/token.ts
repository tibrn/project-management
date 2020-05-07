import axios, { AxiosRequestConfig } from 'axios'
import Cookies from 'js-cookie'
export class TokenStorage {
  private static readonly LOCAL_STORAGE_TOKEN = 'token';

  private static readonly LOCAL_STORAGE_REFRESH_TOKEN = 'refresh_token';

  private static token: null | string

  public static isAuthenticated (): boolean {
    return this.getToken() !== null
  }

  public static getAuthentication () {
    return { Authorization: `Bearer ${this.getToken()}` }
  }

  public static storeToken (token: string): void {
    TokenStorage.token = token
    axios.defaults.headers.common.Authorization = TokenStorage.getAuthentication().Authorization
    localStorage.setItem(TokenStorage.LOCAL_STORAGE_TOKEN, token)
  }

  public static clear (): void {
    TokenStorage.token = null
    axios.defaults.headers.common.Authorization = undefined
  }

  public static getToken (): string | null {
    if (TokenStorage.token === undefined) {
      // TokenStorage.token = Cookies.get('token') || null

      TokenStorage.token = localStorage.getItem(TokenStorage.LOCAL_STORAGE_TOKEN) || null
    }

    return TokenStorage.token
  }
}
