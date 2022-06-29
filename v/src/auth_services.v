module main

import crypto.hmac
import crypto.sha256
import crypto.bcrypt
import encoding.base64
import json
import databases
import time

struct JwtHeader {
	alg string
	typ string
}

struct JwtPayload {
	sub         string    // (subject) = Entidade à quem o token pertence, normalmente o ID do usuário;
	iss         string    // (issuer) = Emissor do token;
	exp         string    // (expiration) = Timestamp de quando o token irá expirar;
	iat         time.Time // (issued at) = Timestamp de quando o token foi criado;
	aud         string    // (audience) = Destinatário do token, representa a aplicação que irá usá-lo.
	name        string
	roles       string
	permissions string
}

fn (mut app App) service_auth(username string, password string) ?string {
	mut db := databases.create_db_connection()

	results := db.query("
		SELECT *
		FROM usersx
		WHERE username = '$username';
	") or {
		println(err)
		return none
	}

	json_string := '${results.maps()[0]}'.replace("'", '"')

	// FIXME - active and id is wrong
	mut user := json.decode(User, json_string) or {
		eprintln('Failed to decode user json, error: $err')
		return none
	}

	// // TODO * Uncomment when fix
	// if !user.active {
	// 	println("usuário não está ativo")
	// 	return none
	// }

	db.free()
	db.close()

	hash := user.password

	bcrypt.compare_hash_and_password(password.bytes(), hash.bytes()) or {
		eprintln('Failed to decode user json, error: $err')
		return none
	}

	token := make_token(user)

	return token
}

fn make_token(user User) string {
	// TODO - get secret from .env
	secret := 'your-256-bit-secret'

	jwt_header := JwtHeader{'HS256', 'JWT'}
	jwt_payload := JwtPayload{
		sub: '$user.id'
		name: '$user.username'
		iat: time.now()
	}

	header := base64.url_encode(json.encode(jwt_header).bytes())
	payload := base64.url_encode(json.encode(jwt_payload).bytes())
	signature := base64.url_encode(hmac.new(secret.bytes(), '${header}.$payload'.bytes(),
		sha256.sum, sha256.block_size).bytestr().bytes())

	jwt := '${header}.${payload}.$signature'

	return jwt
}
