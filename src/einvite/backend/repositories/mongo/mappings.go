package mongo

import (
	"einvite/common/contracts"
	"einvite/framework"
)

func toUserDtos(entities []*dbUser) []*contracts.User {

	dtos := make([]*contracts.User, len(entities))

	for i, entity := range entities {

		dtos[i] = toUserDto(entity)
	}

	return dtos
}

func toUserDto(entity *dbUser) *contracts.User {
	dto := &contracts.User{}

	dto.Email = entity.Email
	dto.Name = entity.Name

	return dto
}

func toEventDto(entity *dbEvent) *contracts.Event {

	return nil
}

func toDbChoosables(dtos []*contracts.Choosable, getNextChoosableIdFunc func() int) []*dbChoosable {

	choosables := make([]*dbChoosable, len(dtos))

	for i, choosableDto := range dtos {
		choosable := &dbChoosable{
			Id:       getNextChoosableIdFunc(),
			Owner:    choosableDto.Owner,
			Type:     choosableDto.Type,
			Voters:   choosableDto.Voters,
			DataType: choosableDto.DataType,
		}

		choosable.Data = choosableDto.Data //TODO: missing bson attributes

		choosables[i] = choosable
	}

	return choosables
}

func toSessionDto(entity *sessionEntity) *contracts.SessionInfo {

	dto := &contracts.SessionInfo{
		Id:     entity.Id.Hex(),
		Values: entity.Values,
		Expiry: entity.Expiry,
	}

	if entity.User != nil {
		dto.User = &framework.SessionUser{
			UserId:   entity.User.Id,
			AuthType: framework.AuthType(entity.User.AuthType),
			AuthData: entity.User.AuthData,
		}
	}

	return dto
}
