import { TokenStorage } from 'src/services/token'
import { MiddlewareCallback } from '../../types/router/middleware'

export default (({ redirect }) => {
  if (TokenStorage.isAuthenticated()) {
    redirect('/')
  }
}) as MiddlewareCallback
