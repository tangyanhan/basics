/* Database: account */
CREATE TYPE urole AS ENUM('admin', 'user', 'tester', 'internal');
CREATE EXTENSION pgcrypto; /* Load extension so we can use random uuid*/
CREATE TABLE IF NOT EXISTS accounts(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(), /* Use random uuid so it become even harder to guess a specified uid, and impossible to guess total count of users*/
	role urole DEFAULT 'user',
	username VARCHAR(20),
	nickname VARCHAR(20),
	email VARCHAR(40),
	secret CHAR(32), /* a combined hash with SHA256*/
	regcode CHAR(32)
);
