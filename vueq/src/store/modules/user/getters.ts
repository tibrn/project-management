import { UserState } from '../../../types/store/modules/user';


export function name(state:UserState) {
  return state.name;
}

export function surname(state:UserState) {
  return state.surname;
}

export function full_name(state:UserState) {
  return `${state.name} ${state.surname}`;
}

export function email(state:UserState) {
  return state.email;
}

export function type(state:UserState) {
  return state.type;
}

export function token(state:UserState) {
  return state.token;
}

export function joined_at(state:UserState) {
  return state.joined_at;
}

export function created_at(state:UserState) {
  return state.created_at;
}
