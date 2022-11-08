-- +goose Up
-- +goose StatementBegin

CREATE TABLE account
(
	id BIGSERIAL NOT NULL,
	active BOOL NOT NULL DEFAULT TRUE,
	uuid TEXT NOT NULL DEFAULT '',
	
	CONSTRAINT pk_account_id PRIMARY KEY (id),
	CONSTRAINT uk_account_uuid UNIQUE (uuid)
);

INSERT INTO account(uuid) VALUES('a27f9278-3834-4f92-a1e7-d158fa113370');

CREATE TABLE address (
	id BIGSERIAL NOT NULL,
	active BOOL NOT NULL DEFAULT TRUE,
	account_id BIGINT NOT NULL DEFAULT 0,
	addr TEXT NOT NULL DEFAULT '',
	request_history BOOL NOT NULL DEFAULT false,
	
	CONSTRAINT pk_address_id PRIMARY KEY (id)
);

CREATE INDEX in_address_account_id ON address ((account_id != 0)) WHERE account_id != 0;

INSERT INTO address(account_id, addr, request_history) VALUES(1, '------------------------------------------', true);

CREATE TABLE tx (
	id BIGSERIAL NOT NULL,
	created_at BIGINT NOT NULL,
	hash TEXT NOT NULL UNIQUE,
	block_num BIGINT NOT NULL,
	from_address_id BIGINT NOT NULL,
	to_address_id BIGINT NOT NULL,
	status TEXT NOT NULL,
	amount     TEXT NOT NULL,
	fee     TEXT NOT NULL,
	direction TEXT NOT NULL,

	CONSTRAINT pk_tx_id PRIMARY KEY (id),
	CONSTRAINT fk_tx_from_address_id FOREIGN KEY (from_address_id) REFERENCES address (id),
	CONSTRAINT fk_tx_to_address_id FOREIGN KEY (to_address_id) REFERENCES address (id)
);

CREATE INDEX in_from_address_id ON tx USING btree (from_address_id);
CREATE INDEX in_to_address_id ON tx USING btree (to_address_id);

CREATE TABLE block (
	 id BIGSERIAL NOT NULL,
	 symbol TEXT NOT NULL,
	 num BIGINT NOT NULL,

	 CONSTRAINT pk_block_id PRIMARY KEY (id)
);

INSERT INTO block(symbol, num) VALUES('ETH', 0);
INSERT INTO block(symbol, num) VALUES('SOL', 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tx;
DROP TABLE IF EXISTS address;
DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS block;
-- +goose StatementEnd
