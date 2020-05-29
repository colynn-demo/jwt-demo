package user

import "github.com/gin-gonic/gin"

// Login ..
type Login struct {
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}

// Model user model defined, in real scnario you should store this into database
type Model struct {
	ID       int
	UserName string
	Email    string
}

// Info ..
func Info(c *gin.Context) {
	c.JSON(200, gin.H{
		"name":  "colynn",
		"email": "colynn.liu@xx.com",
	})
}

// Service ..
type Service struct{}

// NewService ..
func NewService() *Service {
	return &Service{}
}

// Authenticate ..
func (s *Service) Authenticate(req *Login) (*Model, error) {
	// Authenticate here
	// You can add authenticate method, ldap/ local  and others
	// Here, we just return login success
	return &Model{
		UserName: "colynn",
		Email:    "colynn.liu#xx.com",
	}, nil
}
