// +build !windows
// +build cgo,!osusergo

package luser

import "os/user"

func currentUser() (*User, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	return &User{User: u}, nil
}
