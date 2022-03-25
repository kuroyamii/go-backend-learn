package repository

import "github.com/kuroyamii/go-backend-learn/internal/valorant/entity"

type ValorantRepository interface {
	getValorantMember(id int) (entity.ValorantMember, error)
	getValorantMembers() (entity.ValorantMembers, error)
}
