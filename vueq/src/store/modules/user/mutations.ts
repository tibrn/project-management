
import { TokenStorage } from 'src/services/token'
import { instance as axios } from 'src/boot/axios'
import { UserState } from '../../../types/store/modules/user'

export function SET_NAME (state: UserState, name: string) {
  state.name = name
}

export function SET_SURNAME (state: UserState, surname: string) {
  state.surname = surname
}

export function SET_TYPE (state: UserState, type: number) {
  state.type = type
}

export function SET_EMAIL (state: UserState, email: string) {
  state.email = email
}

export function SET_USER (state: UserState, payload: any) {
  if (payload.user) {
    state.name = payload.user.name
    state.surname = payload.user.surname
    state.email = payload.user.email
    state.type = payload.user.type
    state.joined_at = payload.user.joined_at
  }

  if (payload.token) {
    TokenStorage.storeToken(payload.token)

    state.token = TokenStorage.getAuthentication().headers.Authorization

    Object.assign(axios.defaults.headers.common, TokenStorage.getAuthentication().headers)
  }
}
