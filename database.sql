/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */


/** Install uuid module. */
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

/** This is user table to store user information. */
CREATE TABLE users (
	id uuid DEFAULT uuid_generate_v4(),
	phone VARCHAR (13) UNIQUE NOT NULL,
  fullname VARCHAR (60) NOT NULL,
  salt VARCHAR NOT NULL,
  password VARCHAR NOT NULL,
  successful_login INT DEFAULT 0,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

/** This is function to set current timestamp to updated_at column. */
CREATE FUNCTION update_updated_at_users()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ language 'plpgsql';

/** This is the trigger to set current timestamp to updated_at column. */
CREATE TRIGGER update_users_updated_at
  BEFORE UPDATE
  ON
    users
  FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_users();