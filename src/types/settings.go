package types

// Settings represents the available settings of the service
type Settings struct {
	workingDir  string
	lsfUsername string
	lsfPassword string // TODO remove asap!
}

// NewSettings returns a settings object
func NewSettings(workingDir string, lsfUsername string, lsfPassword string) *Settings {
	return &Settings{
		workingDir:  workingDir,
		lsfUsername: lsfUsername,
		lsfPassword: lsfPassword,
	}
}

// WorkingDir returns the working directory
func (s *Settings) WorkingDir() string {
	return s.workingDir
}

// LSFUsername returns the lsf username
func (s *Settings) LSFUsername() string {
	return s.lsfUsername
}

// LSFPassword returns the lsf password
// TODO remove asap!
func (s *Settings) LSFPassword() string {
	return s.lsfPassword
}
