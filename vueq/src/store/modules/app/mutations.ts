import { AppState } from '../../../types/store/modules/app';


export function SET_ERROR(state:AppState, is_error : boolean) {
  state.is_error = Boolean(is_error).valueOf();
}
