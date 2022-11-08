-- +goose Up
-- +goose StatementBegin
CREATE TABLE validator
(
	id BIGSERIAL NOT NULL,
	active BOOL NOT NULL DEFAULT TRUE,
	created_at TIMESTAMP WITH TIME ZONE,
	updated_at TIMESTAMP WITH TIME ZONE,
	network TEXT NOT NULL,
	account TEXT NOT NULL,
	name TEXT NOT NULL,
	www_url TEXT NOT NULL DEFAULT '',
	details TEXT NOT NULL DEFAULT '',
	avatar_url TEXT NOT NULL DEFAULT '',
	total_score BIGINT NOT NULL DEFAULT 0,
	root_distance_score BIGINT NOT NULL DEFAULT 0,
	vote_distance_score BIGINT NOT NULL DEFAULT 0,
	skipped_slot_score BIGINT NOT NULL DEFAULT 0,
	software_version TEXT NOT NULL DEFAULT '',
	software_version_score BIGINT NOT NULL DEFAULT 0,
	stake_concentration_score BIGINT NOT NULL DEFAULT 0,
	data_center_concentration_score BIGINT NOT NULL DEFAULT 0,
	published_information_score BIGINT NOT NULL DEFAULT 0,
	security_report_score BIGINT NOT NULL DEFAULT 0,
	active_stake BIGINT NOT NULL DEFAULT 0,
	commission BIGINT NOT NULL DEFAULT 0,
	delinquent BOOL NOT NULL,
	data_center_key TEXT NOT NULL DEFAULT '',
	data_center_host TEXT NOT NULL DEFAULT '',
	autonomous_system_number BIGINT NOT NULL DEFAULT 0,
	vote_account TEXT NOT NULL UNIQUE,
	epoch_credits BIGINT NOT NULL DEFAULT 0,
	skipped_slots BIGINT NOT NULL DEFAULT 0,
	skipped_slot_percent TEXT NOT NULL DEFAULT '',
	ping_time TEXT NOT NULL DEFAULT '',
	url TEXT NOT NULL DEFAULT '',

	CONSTRAINT pk_validator_id PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS validator;
-- +goose StatementEnd
