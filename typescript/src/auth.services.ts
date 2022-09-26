import { User } from "./user.entities";

type JwtHeader = {
  alg: string;
  typ: string;
};

type JwtPayload = {
  sub: string; // (subject) = Entidade à quem o token pertence, normalmente o ID do usuário;
  iss: string; // (issuer) = Emissor do token;
  exp: string; // (expiration) = Timestamp de quando o token irá expirar;
  iat: number; // (issued at) = Timestamp de quando o token foi criado;
  aud: string; // (audience) = Destinatário do token, representa a aplicação que irá usá-lo.
  name: string;
  roles: string;
  permissions: string;
};

export class AuthService {
  //   constructor(parameters) {}
  auth(username: string, password: string) {}
  authVerify(token: string): boolean {
    return false;
  }
}

function makeToken(user: User): string {
  return "";
}
