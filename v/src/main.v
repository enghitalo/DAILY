module main

import vweb
import databases
// import time

const (
	http_port = 8080
)

struct App {
	vweb.Context // pub mut:
	//     db sqlite.DB
}

fn main() {
	mut db := databases.create_db_connection()
	
	db.query('CREATE TABLE usersx (
		id BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(191) UNIQUE,
		password VARCHAR(191),
		name VARCHAR(191),
		created_at datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
		updated_at datetime(3) DEFAULT NULL,
		deleted_at datetime(3) DEFAULT NULL,
		active BOOLEAN NOT NULL DEFAULT TRUE
		)') or {
		println('error: $err')
	}

	db.free()
	db.close()

	vweb.run(new_app(), http_port)
}

fn new_app() &App {
	mut app := &App{
		// db: sqlite.connect('gitly.sqlite') or { panic(err) }		
	}

	/// TODO - init dotenv
	return app
}
