export interface AuthData {
  joined_at: number;
}

export interface AuthRefresh {
  id: number;
  name: string;
  slug: string;
  email: string;
  settings: {
    avatar: string;
    theme: string;
  };
  type: number;
}
