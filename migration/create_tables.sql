CREATE TABLE IF NOT EXISTS article (
	id serial4 NOT NULL,
	author text NOT NULL,
	title text NOT NULL,
	body text NOT NULL,
	created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT article_pk PRIMARY KEY (id)
);