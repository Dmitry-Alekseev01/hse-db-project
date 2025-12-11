package domain

import "context"

// Club
type ClubRepository interface {
	Create(ctx context.Context, club *Club) (int64, error)
	GetByID(ctx context.Context, id int64) (*Club, error)
	Update(ctx context.Context, club *Club) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Club, error)
}
type ClubUsecase interface {
	Create(ctx context.Context, club *Club) (int64, error)
	GetByID(ctx context.Context, id int64) (*Club, error)
	Update(ctx context.Context, club *Club) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Club, error)
}

// Team
type TeamRepository interface {
	Create(ctx context.Context, team *Team) (int64, error)
	GetByID(ctx context.Context, id int64) (*Team, error)
	Update(ctx context.Context, team *Team) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Team, error)
}
type TeamUsecase interface {
	Create(ctx context.Context, team *Team) (int64, error)
	GetByID(ctx context.Context, id int64) (*Team, error)
	Update(ctx context.Context, team *Team) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Team, error)
}

// Staff
type StaffRepository interface {
	Create(ctx context.Context, staff *Staff) (int64, error)
	GetByID(ctx context.Context, id int64) (*Staff, error)
	Update(ctx context.Context, staff *Staff) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Staff, error)
}
type StaffUsecase interface {
	Create(ctx context.Context, staff *Staff) (int64, error)
	GetByID(ctx context.Context, id int64) (*Staff, error)
	Update(ctx context.Context, staff *Staff) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Staff, error)
}

// Stadium
type StadiumRepository interface {
	Create(ctx context.Context, stadium *Stadium) (int64, error)
	GetByID(ctx context.Context, id int64) (*Stadium, error)
	Update(ctx context.Context, stadium *Stadium) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Stadium, error)
}
type StadiumUsecase interface {
	Create(ctx context.Context, stadium *Stadium) (int64, error)
	GetByID(ctx context.Context, id int64) (*Stadium, error)
	Update(ctx context.Context, stadium *Stadium) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Stadium, error)
}

// Game
type GameRepository interface {
	Create(ctx context.Context, game *Game) (int64, error)
	GetByID(ctx context.Context, id int64) (*Game, error)
	Update(ctx context.Context, game *Game) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Game, error)
}
type GameUsecase interface {
	Create(ctx context.Context, game *Game) (int64, error)
	GetByID(ctx context.Context, id int64) (*Game, error)
	Update(ctx context.Context, game *Game) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Game, error)
}

// Player
type PlayerRepository interface {
	Create(ctx context.Context, player *Player) (int64, error)
	GetByID(ctx context.Context, id int64) (*Player, error)
	Update(ctx context.Context, player *Player) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Player, error)
}
type PlayerUsecase interface {
	Create(ctx context.Context, player *Player) (int64, error)
	GetByID(ctx context.Context, id int64) (*Player, error)
	Update(ctx context.Context, player *Player) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Player, error)
}

// Coach
type CoachRepository interface {
	Create(ctx context.Context, coach *Coach) (int64, error)
	GetByID(ctx context.Context, id int64) (*Coach, error)
	Update(ctx context.Context, coach *Coach) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Coach, error)
}
type CoachUsecase interface {
	Create(ctx context.Context, coach *Coach) (int64, error)
	GetByID(ctx context.Context, id int64) (*Coach, error)
	Update(ctx context.Context, coach *Coach) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Coach, error)
}
