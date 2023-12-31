-- This is the schema file for our database, we are handling the uid for the "pastes" using the uuid and gorm package of GOlang as postgresql doesn't handle so

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE Pastes (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL,
	body VARCHAR,
	name VARCHAR(20),
	PRIMARY KEY (id)
);