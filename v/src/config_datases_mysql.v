module databases

import mysql
import os

pub fn create_db_connection() &mysql.Connection {
	mut db := &mysql.Connection{
		host: os.getenv('DB_HOST')
		port: 3306
		username: 'hitalo'
		password: 'password'
		dbname: 'daily-homolog'
	}

	db.connect() or { panic(err) }

	return db
}
