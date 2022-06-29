module main

struct AuthRequestDto {
	// Adding a [required] attribute will make decoding fail, if that field is not present in the input.
	// If a field is not [required], but is missing, it will be assumed to have its default value, like 0 for numbers, or '' for strings, and decoding will not fail.
	username string [required]
	password string [required]
	// // Use the `skip` attribute to skip certain fields
	// foo Foo [skip]
	// // If the field name is different in JSON, it can be specified
	// last_name string [json: lastName]
}
