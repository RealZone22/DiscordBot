package handlers

import (
	"github.com/RealZone22/DiscordBot/pkg/utils"
	"time"
)

type Warn struct {
	ID        int64     `json:"id" xorm:"pk autoincr"`
	UserID    string    `json:"user_id" xorm:"not null"`
	Reason    string    `json:"reason" xorm:"not null"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
}

func CreateWarn(userID string, reason string) error {
	warn := &Warn{
		UserID: userID,
		Reason: reason,
	}

	_, err := utils.DB.Insert(warn)
	if err != nil {
		utils.Logger.Error().Err(err).Msg("Failed to create warn")
		return err
	}

	utils.Logger.Debug().Str("user_id", userID).Str("reason", reason).Msg("Warn created successfully")
	return nil
}

func GetWarns(userID string) ([]Warn, error) {
	var warns []Warn
	err := utils.DB.Where("user_id = ?", userID).Find(&warns)
	if err != nil {
		utils.Logger.Error().Err(err).Msg("Failed to retrieve warns")
		return nil, err
	}

	utils.Logger.Debug().Str("user_id", userID).Int("count", len(warns)).Msg("Retrieved warns successfully")
	return warns, nil
}

func DeleteWarn(id int64) error {
	_, err := utils.DB.Delete(&Warn{ID: id})
	if err != nil {
		utils.Logger.Error().Err(err).Msg("Failed to delete warn")
		return err
	}

	utils.Logger.Debug().Int64("warn_id", id).Msg("Warn deleted successfully")
	return nil
}

func ClearWarns(userID string) error {
	_, err := utils.DB.Where("user_id = ?", userID).Delete(&Warn{})
	if err != nil {
		utils.Logger.Error().Err(err).Msg("Failed to clear warns")
		return err
	}

	utils.Logger.Debug().Str("user_id", userID).Msg("All warns cleared successfully")
	return nil
}

func GetWarnCount(userID string) (int, error) {
	var count int64
	count, err := utils.DB.Where("user_id = ?", userID).Count(&Warn{})
	if err != nil {
		utils.Logger.Error().Err(err).Msg("Failed to get warn count")
		return 0, err
	}

	utils.Logger.Info().Str("user_id", userID).Int64("count", count).Msg("Retrieved warn count successfully")
	return int(count), nil
}
