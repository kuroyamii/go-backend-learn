package repository

import (
	"database/sql"

	"github.com/kuroyamii/go-backend-learn/internal/valorant/entity"
)

const (
	GET_VALORANT_MEMBERS = `SELECT ingame_name,role,nationality
	FROM valorant;`

	GET_VALORANT_MEMBER_BY_ID = `SELECT ingame_name,role,nationality
	FROM valorant
	WHERE member_id = ?`
)

type valorantRepository struct {
	DB *sql.DB
}

func ProvideValorantRepository(DB *sql.DB) *valorantRepository {
	return &valorantRepository{DB: DB}
}

func (vr *valorantRepository) getValorantMember(id int) (entity.ValorantMember, error) {
	return entity.ValorantMember{}, nil
}

func (vr *valorantRepository) getValorantMembers() (entity.ValorantMembers, error) {
	return entity.ValorantMembers{}, nil
}
