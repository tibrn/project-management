import { AppState } from './modules/app';
import { UserState } from './modules/user';

export interface RootState {
  user:UserState
  app:AppState
}
