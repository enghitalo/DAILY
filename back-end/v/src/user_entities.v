module main

[table: 'usersxqa']
struct User {
mut:
	id         int    [primary; sql: serial]
	username   string [required; sql_type: 'varchar(191)']
	password   string [required; sql_type: 'longtext']
	name       string [sql_type: 'varchar(191)']
	created_at string [default: 'CURRENT_TIMESTAMP'; sql_type: 'datetime(3)']
	updated_at string [default: 'CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP'; sql_type: 'datetime(3)']
	deleted_at string [default: 'CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP'; sql_type: 'datetime(3)']
	active     bool
}
