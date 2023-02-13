import { User } from "./user.entities";
import { timingSafeEqual, createHmac } from "crypto";

type JwtHeader = {
  alg: string;
  typ: string;
};

type JwtPayload = {
  sub: string; // (subject) = Entidade à quem o token pertence, normalmente o ID do usuário;
  iss?: string; // (issuer) = Emissor do token;
  exp?: string; // (expiration) = Timestamp de quando o token irá expirar;
  iat: number; // (issued at) = Timestamp de quando o token foi criado;
  aud?: string; // (audience) = Destinatário do token, representa a aplicação que irá usá-lo.
  name: string;
  roles?: string;
  permissions?: string;
};

export class AuthService {
  //   constructor(parameters) {}
  auth(username: string, password: string) {}
  authVerify(token: string): boolean {
    if (token == "") {
      return false;
    }
    const secret = "SECRET_KEY"; // os.getenv('SECRET_KEY')
    const token_split = token.split(".");

    const signatureMirror: Buffer = Buffer.from(createHmac("sha256", secret).update(`${token_split[0]}.${token_split[1]}`).digest("base64url"));

    const signatureFromToken = Buffer.from(token_split[2]);
    return timingSafeEqual(signatureMirror, signatureFromToken);
  }
}

function makeToken(user: User): string {
  const secret = "SECRET_KEY"; //os.getenv('SECRET_KEY')

  const jwt_header: JwtHeader = { alg: "HS256", typ: "JWT" };
  const jwt_payload: JwtPayload = {
    sub: `${user.id}`,
    name: `${user.name}`,
    iat: Date.now(),
  };

  const header: string = Buffer.from(JSON.stringify(jwt_header)).toString("base64url");

  const payload: string = Buffer.from(JSON.stringify(jwt_payload)).toString("base64url");

  const signature = createHmac("sha256", secret).update(`${header}.${payload}`).digest("base64url");

  const jwt: string = `${header}.${payload}.${signature}`;

  return jwt;
}
