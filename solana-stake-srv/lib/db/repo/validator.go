package repo

import (
	"github.com/gocraft/dbr/v2"

	"github.com/truekupo/cluster/solana-stake-srv/lib/db/models"
)

var (
	validatorFields = []string{
		"validator.network",
		"validator.account",
		"validator.name",
		"validator.www_url",
		"validator.details",
		"validator.avatar_url",
		"validator.created_at",
		"validator.updated_at",
		"validator.total_score",
		"validator.root_distance_score",
		"validator.vote_distance_score",
		"validator.skipped_slot_score",
		"validator.software_version",
		"validator.software_version_score",
		"validator.stake_concentration_score",
		"validator.data_center_concentration_score",
		"validator.published_information_score",
		"validator.security_report_score",
		"validator.active_stake",
		"validator.commission",
		"validator.delinquent",
		"validator.data_center_key",
		"validator.data_center_host",
		"validator.autonomous_system_number",
		"validator.vote_account",
		"validator.epoch_credits",
		"validator.skipped_slots",
		"validator.skipped_slot_percent",
		"validator.ping_time",
		"validator.url",
	}
)

// Repos
type ValidatorRepo struct {
	Repository
}

// Constructors
func Validator(session *dbr.Session) *ValidatorRepo {
	return &ValidatorRepo{New(session, "validator", validatorFields)}
}

func (r *ValidatorRepo) FindById(id int64) (*models.Validator, error) {
	t := models.Validator{}

	count, err := r.Select().Where("validator.id=?", id).Load(&t)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, dbr.ErrNotFound
	}

	return &t, nil
}

func (r *ValidatorRepo) FindByOrder(From int32, Limit int32, Field string, Order string) ([]models.Validator, error) {
	t := []models.Validator{}

	var (
		limit int32 = Limit
		count int
		err   error
	)

	if limit == 0 {
		limit = 99999
	}

	b := r.Select().Offset(uint64(From)).Limit(uint64(limit))
	if Field != "" {
		if Order == "ASC" {
			b = b.OrderAsc("validator." + Field)
		}

		if Order == "DESC" {
			b = b.OrderDesc("validator." + Field)
		}
	}

	count, err = b.Load(&t)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, dbr.ErrNotFound
	}

	return t, nil
}

func (r *ValidatorRepo) FindByAddress(Address string) (*models.Validator, error) {
	t := models.Validator{}

	count, err := r.Select().Where("validator.account=?", Address).Load(&t)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, dbr.ErrNotFound
	}

	return &t, nil
}

func (r *ValidatorRepo) InsertOrUpdate(t *models.Validator) (*models.Validator, error) {
	tt, err := r.Insert(t)
	if err == nil {
		return tt, nil
	}

	return r.Update(t)
}

func (r *ValidatorRepo) Insert(t *models.Validator) (*models.Validator, error) {
	return t, r.InsertBuilder().Columns("network", "account", "name", "www_url", "details", "avatar_url", "created_at", "updated_at", "total_score", "root_distance_score", "vote_distance_score",
		"skipped_slot_score", "software_version", "software_version_score", "stake_concentration_score", "data_center_concentration_score", "published_information_score", "security_report_score",
		"active_stake", "commission", "delinquent", "data_center_key", "data_center_host", "autonomous_system_number", "vote_account", "epoch_credits", "skipped_slots", "skipped_slot_percent",
		"ping_time", "url").Record(t).Returning("id").Load(&t.Id)
}

func (r *ValidatorRepo) Update(t *models.Validator) (*models.Validator, error) {
	_, err := r.UpdateBuilder().Where("id=?", t.Id).
		Set("network", t.Network).
		Set("account", t.Account).
		Set("name", t.Name).
		Set("www_url", t.WwwUrl).
		Set("details", t.Details).
		Set("avatar_url", t.AvatarUrl).
		Set("created_at", t.CreatedAt).
		Set("updated_at", t.UpdatedAt).
		Set("total_score", t.TotalScore).
		Set("root_distance_score", t.RootDistanceScore).
		Set("vote_distance_score", t.VoteDistanceScore).
		Set("skipped_slot_score", t.SkippedSlotScore).
		Set("software_version", t.SoftwareVersion).
		Set("software_version_score", t.SoftwareVersionScore).
		Set("stake_concentration_score", t.StakeConcentrationScore).
		Set("data_center_concentration_score", t.DataCenterConcentrationScore).
		Set("published_information_score", t.PublishedInformationScore).
		Set("security_report_score", t.SecurityReportScore).
		Set("active_stake", t.ActiveStake).
		Set("commission", t.Commission).
		Set("delinquent", t.Delinquent).
		Set("data_center_key", t.DataCenterKey).
		Set("data_center_host", t.DataCenterHost).
		Set("autonomous_system_number", t.AutonomousSystemNumber).
		Set("vote_account", t.VoteAccount).
		Set("epoch_credits", t.EpochCredits).
		Set("skipped_slots", t.SkippedSlots).
		Set("skipped_slot_percent", t.SkippedSlotPercent).
		Set("ping_time", t.PingTime).
		Set("url", t.Url).Exec()

	return t, err
}
