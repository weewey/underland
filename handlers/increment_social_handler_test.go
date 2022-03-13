package handlers

import (
	"github.com/stretchr/testify/assert"
	"github.com/underland/models"
	"testing"
)

func TestIncrementSocial(t *testing.T) {
	metadata := models.CharacterMetaData{
		Name:        "Varick",
		Symbol:      "",
		Image:       "imageUrl",
		SocialIndex: 100,
		Attributes: []*models.Attributes{
			{
				TraitType: models.SOCIAL,
				Value:     1,
			},
			{
				TraitType: models.SOCIAL_INDEX,
				Value:     100,
			},
			{
				TraitType: models.BEHAVIOUR,
				Value:     1,
			},
			{
				TraitType: models.CONTRIBUTION,
				Value:     1,
			},
			{
				TraitType: models.INFLUENCE,
				Value:     1,
			},
			{
				TraitType: models.RESPECT,
				Value:     1,
			},
		},
		Properties: &models.Properties{
			Files:    nil,
			Creators: nil,
		},
	}
	IncrementSocial(1, &metadata)
	assert.Equal(t, 120, metadata.SocialIndex)
}
