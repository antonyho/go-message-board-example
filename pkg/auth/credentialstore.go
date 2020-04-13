package auth

// CredentialStore stores user accounts' credential
type CredentialStore interface {
	// Check if provide username and password matches on any user account
	Check(username, password string) bool
}

// MockCredentialStore is a hardcoded CredentialStore to act like a credential service
type MockCredentialStore struct {
	credentials map[string]string
}

// NewMockCredentialStore creates a new dummy CredentialStore
func NewMockCredentialStore() CredentialStore {
	return &MockCredentialStore{
		credentials: map[string]string{
			"admin": "back-challenge", // Hardcoded credential just lives in coding example
		},
	}
}

// Check against user's credential in credential store
func (s *MockCredentialStore) Check(username, password string) bool {
	userPwd, exist := s.credentials[username]
	if !exist {
		return false
	}
	return password == userPwd
}
