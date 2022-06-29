module main

import vweb
import json

['/user/:id/get'; get]
pub fn (mut app App) controller_get_user_by_id(id string) vweb.Result {
	response := app.service_get_user("usernamexid", "passwordx")
	return app.json(response)
}

['/user/create'; post]
pub fn (mut app App) controller_create_user() vweb.Result {

	body := json.decode(User, app.req.data) or {
		eprintln('Failed to decode json, error: $err')
		app.set_status(400, '')
		return app.text( 'Failed to decode json, error: $err')
	}
	
	response := app.service_add_user(body.username, body.password) or {
		app.set_status(400, '')
		return app.text( 'Failed to decode json, error: $err')
	}
	return app.json(response)
}

['/user/get_all'; get]
pub fn (mut app App) controller_get_all_user() vweb.Result {
	response := app.service_get_user("usernamexgetAll", "passwordx")
	return app.json(response)
}