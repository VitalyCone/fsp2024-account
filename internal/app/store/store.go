package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Store struct {
	config                *Config
	db                    *sql.DB
	usersRepository       *UsersRepository
	tagRepository         *TagRepository
	serviceTypeRepository *ServiceTypeRepository
	reviewRepository      *ReviewRepository
}

func NewStore(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	log.Println("Database is working!")

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UsersRepository {
	if s.usersRepository != nil {
		return s.usersRepository
	}

	s.usersRepository = &UsersRepository{
		store: s,
	}

	return s.usersRepository
}

func (s *Store) Tag() *TagRepository {
	if s.tagRepository != nil {
		return s.tagRepository
	}

	s.tagRepository = &TagRepository{
		store: s,
	}

	return s.tagRepository
}

func (s *Store) ServiceType() *ServiceTypeRepository {
	if s.serviceTypeRepository != nil {
		return s.serviceTypeRepository
	}

	s.serviceTypeRepository = &ServiceTypeRepository{
		store: s,
	}

	return s.serviceTypeRepository
}

func (s *Store) Review() *ReviewRepository {
	if s.reviewRepository != nil {
		return s.reviewRepository
	}

	s.reviewRepository = &ReviewRepository{
		store: s,
	}

	return s.reviewRepository
}
