-- +goose Up
-- +goose StatementBegin
CREATE TABLE coin
(
	id BIGSERIAL NOT NULL,
	symbol TEXT NOT NULL,
	name TEXT NOT NULL DEFAULT '',

	CONSTRAINT pk_coin_id PRIMARY KEY (id)
);

INSERT INTO coin(symbol, name) VALUES('BTC', 'Bitcoin');
INSERT INTO coin(symbol, name) VALUES('ETH', 'Ethereum');
INSERT INTO coin(symbol, name) VALUES('SOL', 'Solana');
INSERT INTO coin(symbol, name) VALUES('USDT', 'Tether');

CREATE TABLE pair
(
	id BIGSERIAL NOT NULL,
	active bool DEFAULT TRUE,
	coin_symbol TEXT NOT NULL,
	currency_symbol TEXT NOT NULL,

	CONSTRAINT pk_pair_id PRIMARY KEY (id)
);

INSERT INTO pair(coin_symbol, currency_symbol) VALUES('BTC', 'USDT');
INSERT INTO pair(coin_symbol, currency_symbol) VALUES('ETH', 'USDT');
INSERT INTO pair(coin_symbol, currency_symbol) VALUES('SOL', 'USDT');

CREATE TABLE kline
(
	id BIGSERIAL NOT NULL,
	pair_id BIGINT NOT NULL,
	open_time BIGINT NOT NULL,
	close_time BIGINT NOT NULL,
	open TEXT NOT NULL,
	close TEXT NOT NULL,
	high TEXT NOT NULL,
	low TEXT NOT NULL,
	
	CONSTRAINT pk_kline_id PRIMARY KEY (id),
	CONSTRAINT fk_kline_pair_id FOREIGN KEY (pair_id) REFERENCES pair (id)
);

CREATE INDEX in_kline_pair_id ON kline	USING btree (pair_id);
CREATE UNIQUE INDEX in_kline_pair_id_open_time_close_time_idx ON kline (pair_id, open_time, close_time);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS kline;
DROP TABLE IF EXISTS pair;
DROP TABLE IF EXISTS coin;
-- +goose StatementEnd
