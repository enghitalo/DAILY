module main
// import time
// import mysql

[table: 'usersx']
struct User {
mut:
	id          int    [primary; sql: serial]
	username string 
	password string 
	name string 
	created_at  string
	updated_at  string
	deleted_at string
	active  bool
}

