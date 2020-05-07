import { TokenStorage } from './../../../services/token'

import { instance as axios } from 'src/boot/axios'

export async function REFRESH () {
  try {
    const { data } = await axios.get('api/users/refresh')

    TokenStorage.storeToken(data.data.token)
  } catch (e) { }
}
