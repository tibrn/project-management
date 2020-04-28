export interface UserState {
  name: string;
  surname: string;
  email: string;
  type: number;
  token: string | null;
  joined_at: Date | null;
  created_at: Date | null;
}
