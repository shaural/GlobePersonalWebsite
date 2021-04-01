package db

import (
	"context"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // import postgres
	"github.com/shaural/GlobePersonalWebsite/server/pkg/common"
)

// Database ...
type Database interface {
	Initialize() error
	UpdateCountry(*Country) error
	UpdateState(*State) error
	UpdateCard(*Card) error
	GetCountries() ([]*Country, error)
	GetStates() ([]*State, error)
	GetCards() ([]*Card, error)
}

type gormDb struct {
	Database *gorm.DB
	ctx      context.Context
}

// NewDatabase initializes a new instance of a Database
func NewDatabase(ctx context.Context) (Database, error) {
	config := common.Config()
	db, err := gorm.Open("postgres", config.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return &gormDb{
		Database: db,
		ctx:      ctx,
	}, nil
}

func (gdb *gormDb) Initialize() error {
	if err := gdb.Database.Transaction(func(tx *gorm.DB) error {
		log.Println("gorm: Automigrate tables")
		return tx.AutoMigrate(&Country{}, &State{}, &Card{}).Error
	}); err != nil {
		return err
	}
	return nil
}

func (gdb *gormDb) Close() error {
	return gdb.Database.Close()
}

func (gdb *gormDb) UpdateCountry(country *Country) error {
	return gdb.Database.
		Where(Country{ID: country.ID}).
		Assign(country).
		FirstOrCreate(&country).Error
}

func (gdb *gormDb) UpdateState(state *State) error {
	return gdb.Database.
		Where(State{ID: state.ID}).
		Assign(state).
		FirstOrCreate(&state).Error
}

func (gdb *gormDb) UpdateCard(card *Card) error {
	return gdb.Database.
		Where(Card{CountryID: card.CountryID, Title: card.Title}).
		Assign(card).
		FirstOrCreate(&card).Error
}

func (gdb *gormDb) GetCountries() (countries []*Country, err error) {
	return countries, gdb.Database.Find(&countries).Error
}

func (gdb *gormDb) GetStates() (states []*State, err error) {
	return states, gdb.Database.Find(&states).Error
}

func (gdb *gormDb) GetCards() (cards []*Card, err error) {
	return cards, gdb.Database.Find(&cards).Error
}
