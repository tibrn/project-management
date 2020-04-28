import { AppState } from '../../../types/store/modules/app';

export function is_error(state:AppState) {
  return state.is_error;
}
