import axios, { AxiosRequestConfig } from 'axios';


export class TokenStorage {
  private static readonly LOCAL_STORAGE_TOKEN = 'token';

  private static readonly LOCAL_STORAGE_REFRESH_TOKEN = 'refresh_token';

  private static refresh_token: undefined | null | string

  private static token: undefined | null | string

  public static isAuthenticated (): boolean {
    return this.getToken() !== null;
  }

  public static getAuthentication (): AxiosRequestConfig {
    return {
      headers: { Authorization: `Bearer ${this.getToken()}` },
    };
  }

  public static getNewToken (): Promise<string> {
    return new Promise((resolve: any, reject: any) => {
      axios
        .post(`${process.env.API}/api/token/refresh`, { refresh_token: this.getRefreshToken() })
        .then(({ data }) => {
          this.storeToken(data.token);
          this.storeRefreshToken(data.refresh_token);

          resolve(data.token);
        })
        .catch((error) => {
          reject(error);
        });
    });
  }

  public static storeToken (token: string): void {
    localStorage.setItem(TokenStorage.LOCAL_STORAGE_TOKEN, token);
    TokenStorage.token = token;
  }

  public static storeRefreshToken (refreshToken: string): void {
    localStorage.setItem(TokenStorage.LOCAL_STORAGE_REFRESH_TOKEN, refreshToken);
    TokenStorage.refresh_token = refreshToken;
  }

  public static clear (): void {
    localStorage.removeItem(TokenStorage.LOCAL_STORAGE_TOKEN);
    localStorage.removeItem(TokenStorage.LOCAL_STORAGE_REFRESH_TOKEN);
    TokenStorage.refresh_token = undefined;
    TokenStorage.token = undefined;
  }

  public static getRefreshToken (): string | null {
    if (TokenStorage.refresh_token === undefined) {
      TokenStorage.refresh_token = localStorage.getItem(TokenStorage.LOCAL_STORAGE_REFRESH_TOKEN);
    }
    return TokenStorage.refresh_token;
  }

  public static getToken (): string | null {
    if (TokenStorage.token === undefined) {
      TokenStorage.token = localStorage.getItem(TokenStorage.LOCAL_STORAGE_TOKEN);
    }

    return TokenStorage.token;
  }
}
