import { TokenStorage } from 'src/services/token';
import { UserState } from '../../../types/store/modules/user';

export default function (): UserState {
  return {
    name: '',
    surname: '',
    email: '',
    type: 0,
    joined_at: null,
    created_at: null,
    token: TokenStorage.getToken(),
  };
}
