module main

import crypto.bcrypt
import databases
import json

fn (mut app App) service_add_user(username string, password string) ?User {
	mut db := databases.create_db_connection()
	hashed_password := bcrypt.generate_from_password(password.bytes(), bcrypt.min_cost) or {
		panic(err)
	}

	db.query("
		INSERT INTO usersx(username, password)
		VALUES
		('$username', '$hashed_password');
	") or {
		println(err)
		return none
	}

	results := db.query("
		SELECT *
		FROM usersx
		WHERE username = '$username'
	") or {
		println(err)
		return none
	}

	json_string := '${results.maps()[0]}'.replace("'", '"')

	mut user := json.decode(User, json_string) or {
		eprintln('Failed to decode user json, error: $err')
		return none
	}

	db.free()
	db.close()
	return user
}

fn (mut app App) service_get_user(username string, password string) User {
	user := User{
		// username: username
		// password: password
		// created_at: time.now()
	}

	// sql app.db {
	// 	insert user into User
	// }
	return user
}

// fn (mut app App) service_get_user_by_id(user_id int) []User {
// 	mut users := sql app.db {
// 		select from User where id == user_id
// 	}

// 	for i, user in users {
// 		users[i].username = app.find_username_by_id(user.id)
// 	}

// 	return users
// }

// fn (mut app App) service_get_all_user(user_id int) []User {
// 	// mut users := sql app.db {
// 	// 	select from User where id == user_id
// 	// }

// 	// for i, user in users {
// 	// 	users[i].username = app.find_username_by_id(user.id)
// 	// }

// 	// return users

// 	db = databases.create_db_connection()

// 	find_user := db.query('
// 		SELECT *
// 		FROM usersx
// 		WHERE username LIKE "dada";
// 	')?

// 	mut users := []User{}

// 	for user in find_user.rows() {
// 	// 		users.push(user)
// 		users << user
// 		// Access the name of user
// 		println(user)
// 	}

// 	for user in find_user.maps() {
// 	// 		users.push(user)
// 		users << user
// 		// Access the name of user
// 		println(user)
// 	}

// 	// Free the query result
// 	find_user.free()
// 	db.close()
// 	return users
// }

// fn (user User) relative() string {
// 	return time.unix(user.created_at).relative()
// }
