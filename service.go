package main

type UserService interface {
  Register(user string, password string) error
}

type userService struct {
  db *sql.DB
}

func (s *userService) Register(user string, password string) error {
  // Implement the Register method here, using the db field of the userService struct to access the database
}
