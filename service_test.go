package main

type mockUserService struct {
  registerFn func(user string, password string) error
}

func (s *mockUserService) Register(user string, password string) error {
  return s.registerFn(user, password)
}


func TestAPIWithSuccess(t *testing.T) {
  s := &mockUserService{
    registerFn: func(user string, password string) error {
      return nil
    },
  }
}

func TestAPIWithError(t *testing.T) {
  s := &mockUserService{
	registerFn: func(user string, password string) error {
		return fmt.Errorf("error")
	},
	}
}
