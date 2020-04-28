import { TokenStorage } from 'src/services/token';
import { MiddlewareCallback } from '../../types/router/middleware';


export default (({ redirect }) => {
  console.log(TokenStorage.isAuthenticated());
  if (TokenStorage.isAuthenticated()) {
    redirect('/');
  }
}) as MiddlewareCallback;
