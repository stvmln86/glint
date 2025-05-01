// Package sqls implements SQLite pragma, schema and query definitions.
package sqls

// Pragma is the default always-enabled database pragma.
const Pragma = `
	pragma encoding = 'utf-8';
	pragma foreign_keys = on;
`

// Schema is the default first-run database schema.
const Schema = `
	create table Notes (
		n_id integer primary key asc,
		init integer not null default (unixepoch()),
		name text    not null unique
	);

	create table Pages (
		p_id integer primary key asc,
		init integer not null default (unixepoch()),
		note integer not null references Notes(n_id),
		body text    not null
	);

	create view LiveNotes as
		select * from Notes where (select count(*) from Pages where note=n_id) > 0;

	create view DeadNotes as
		select * from Notes where (select count(*) from Pages where note=n_id) = 0;
`
