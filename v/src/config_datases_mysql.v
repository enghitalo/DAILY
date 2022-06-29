module databases

import mysql

pub fn create_db_connection() &mysql.Connection {
	mut db := &mysql.Connection{
		host: 'localhost'
		port: 3306
		username: 'hitalo'
		password: 'password'
		dbname: 'daily-homolog'
	}

	db.connect() or { panic(err) }

	return db
}