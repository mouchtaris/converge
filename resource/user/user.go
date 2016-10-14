// Copyright © 2016 Asteris, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package user

import (
	"fmt"
	"os/user"

	"github.com/asteris-llc/converge/resource"
	"github.com/pkg/errors"
)

// State type for User
type State string

const (
	// StatePresent indicates the user should be present
	StatePresent State = "present"

	// StateAbsent indicates the user should be absent
	StateAbsent State = "absent"
)

// User manages user users
type User struct {
	Username  string
	UID       string
	GroupName string
	GID       string
	Name      string
	HomeDir   string
	State     State
	system    SystemUtils
}

// AddUserOptions are the options specified in the configuration to be used
// when adding a user
type AddUserOptions struct {
	UID       string
	Group     string
	Comment   string
	Directory string
}

// SystemUtils provides system utilities for user
type SystemUtils interface {
	AddUser(userName string, options *AddUserOptions) error
	DelUser(userName string) error
	Lookup(userName string) (*user.User, error)
	LookupID(userID string) (*user.User, error)
	LookupGroup(groupName string) (*user.Group, error)
	LookupGroupID(groupID string) (*user.Group, error)
}

// ErrUnsupported is used when a system is not supported
var ErrUnsupported = fmt.Errorf("user: not supported on this system")

// NewUser constructs and returns a new User
func NewUser(system SystemUtils) *User {
	return &User{
		system: system,
	}
}

// Check if a user user exists
func (u *User) Check(resource.Renderer) (resource.TaskStatus, error) {
	var (
		userByID *user.User
		uidErr   error
	)

	// lookup the user by name and lookup the user by uid
	// the lookups return ErrUnsupported if the system is not supported
	// Lookup returns user.UnknownUserError if the user is not found
	// LookupID returns user.UnknownUserIdError if the uid is not found
	userByName, nameErr := u.system.Lookup(u.Username)
	if u.UID != "" {
		userByID, uidErr = u.system.LookupID(u.UID)
	}

	status := &resource.Status{}

	if nameErr == ErrUnsupported {
		status.RaiseLevel(resource.StatusFatal)
		return status, ErrUnsupported
	}

	switch u.State {
	case StatePresent:
		switch {
		case u.UID == "":
			_, nameNotFound := nameErr.(user.UnknownUserError)

			switch {
			case userByName != nil:
				status.Output = append(status.Output, fmt.Sprintf("user %s already exists", u.Username))
			case nameNotFound:
				switch {
				case u.GroupName != "":
					_, err := u.system.LookupGroup(u.GroupName)
					if err != nil {
						status.RaiseLevel(resource.StatusCantChange)
						status.Output = append(status.Output, fmt.Sprintf("group %s does not exist", u.GroupName))
						return status, fmt.Errorf("cannot add user %s", u.Username)
					}
				case u.GID != "":
					_, err := u.system.LookupGroupID(u.GID)
					if err != nil {
						status.RaiseLevel(resource.StatusCantChange)
						status.Output = append(status.Output, fmt.Sprintf("group gid %s does not exist", u.GID))
						return status, fmt.Errorf("cannot add user %s", u.Username)
					}
				}
				status.RaiseLevel(resource.StatusWillChange)
				status.Output = append(status.Output, "user does not exist")
				status.AddDifference("user", string(StateAbsent), fmt.Sprintf("user %s", u.Username), "")
			}
		case u.UID != "":
			_, nameNotFound := nameErr.(user.UnknownUserError)
			_, uidNotFound := uidErr.(user.UnknownUserIdError)

			switch {
			case nameNotFound && uidNotFound:
				switch {
				case u.GroupName != "":
					_, err := u.system.LookupGroup(u.GroupName)
					if err != nil {
						status.RaiseLevel(resource.StatusCantChange)
						status.Output = append(status.Output, fmt.Sprintf("group %s does not exist", u.GroupName))
						return status, fmt.Errorf("cannot add user %s with uid %s", u.Username, u.UID)
					}
				case u.GID != "":
					_, err := u.system.LookupGroupID(u.GID)
					if err != nil {
						status.RaiseLevel(resource.StatusCantChange)
						status.Output = append(status.Output, fmt.Sprintf("group gid %s does not exist", u.GID))
						return status, fmt.Errorf("cannot add user %s with uid %s", u.Username, u.UID)
					}
				}
				status.RaiseLevel(resource.StatusWillChange)
				status.Output = append(status.Output, "user name and uid do not exist")
				status.AddDifference("user", string(StateAbsent), fmt.Sprintf("user %s with uid %s", u.Username, u.UID), "")
			case nameNotFound:
				status.RaiseLevel(resource.StatusCantChange)
				status.Output = append(status.Output, fmt.Sprintf("user uid %s already exists", u.UID))
				return status, fmt.Errorf("cannot add user %s with uid %s", u.Username, u.UID)
			case uidNotFound:
				status.RaiseLevel(resource.StatusCantChange)
				status.Output = append(status.Output, fmt.Sprintf("user %s already exists", u.Username))
				return status, fmt.Errorf("cannot add user %s with uid %s", u.Username, u.UID)
			case userByName != nil && userByID != nil && userByName.Name != userByID.Name || userByName.Uid != userByID.Uid:
				status.RaiseLevel(resource.StatusCantChange)
				status.Output = append(status.Output, fmt.Sprintf("user %s and uid %s belong to different users", u.Username, u.UID))
				return status, fmt.Errorf("cannot add user %s with uid %s", u.Username, u.UID)
			case userByName != nil && userByID != nil && *userByName == *userByID:
				status.RaiseLevel(resource.StatusCantChange)
				status.Output = append(status.Output, fmt.Sprintf("user %s with uid %s already exists", u.Username, u.UID))
				return status, fmt.Errorf("cannot add user %s with uid %s", u.Username, u.UID)
			}
		}
	case StateAbsent:
		switch {
		case u.UID == "":
			_, nameNotFound := nameErr.(user.UnknownUserError)

			switch {
			case nameNotFound:
				status.Output = append(status.Output, fmt.Sprintf("user %s does not exist", u.Username))
			case userByName != nil:
				status.RaiseLevel(resource.StatusWillChange)
				status.AddDifference("user", fmt.Sprintf("user %s", u.Username), string(StateAbsent), "")
			}
		case u.UID != "":
			_, nameNotFound := nameErr.(user.UnknownUserError)
			_, uidNotFound := uidErr.(user.UnknownUserIdError)

			switch {
			case nameNotFound && uidNotFound:
				status.Output = append(status.Output, "user name and uid do not exist")
			case nameNotFound:
				status.RaiseLevel(resource.StatusCantChange)
				status.Output = append(status.Output, fmt.Sprintf("user %s does not exist", u.Username))
				return status, fmt.Errorf("cannot delete user %s with uid %s", u.Username, u.UID)
			case uidNotFound:
				status.RaiseLevel(resource.StatusCantChange)
				status.Output = append(status.Output, fmt.Sprintf("user uid %s does not exist", u.UID))
				return status, fmt.Errorf("cannot delete user %s with uid %s", u.Username, u.UID)
			case userByName != nil && userByID != nil && userByName.Name != userByID.Name || userByName.Uid != userByID.Uid:
				status.RaiseLevel(resource.StatusCantChange)
				status.Output = append(status.Output, fmt.Sprintf("user %s and uid %s belong to different users", u.Username, u.UID))
				return status, fmt.Errorf("cannot delete user %s with uid %s", u.Username, u.UID)
			case userByName != nil && userByID != nil && *userByName == *userByID:
				status.RaiseLevel(resource.StatusWillChange)
				status.AddDifference("user", fmt.Sprintf("user %s with uid %s", u.Username, u.UID), string(StateAbsent), "")
			}
		}
	default:
		status.RaiseLevel(resource.StatusFatal)
		return status, fmt.Errorf("user: unrecognized state %v", u.State)
	}

	return status, nil
}

// Apply changes for user
func (u *User) Apply() (resource.TaskStatus, error) {
	var (
		userByID *user.User
		uidErr   error
	)

	// lookup the user by name and lookup the user by uid
	// the lookups return ErrUnsupported if the system is not supported
	// Lookup returns user.UnknownUserError if the user is not found
	// LookupID returns user.UnknownUserIdError if the uid is not found
	userByName, nameErr := u.system.Lookup(u.Username)
	if u.UID != "" {
		userByID, uidErr = u.system.LookupID(u.UID)
	}

	status := &resource.Status{}

	if nameErr == ErrUnsupported {
		status.RaiseLevel(resource.StatusFatal)
		return status, ErrUnsupported
	}

	switch u.State {
	case StatePresent:
		switch {
		case u.UID == "":
			_, nameNotFound := nameErr.(user.UnknownUserError)

			switch {
			case nameNotFound:
				options, err := SetAddUserOptions(u)
				if err != nil {
					status.RaiseLevel(resource.StatusCantChange)
					status.Output = append(status.Output, err.Error())
					return status, fmt.Errorf("will not attempt add: user %s", u.Username)
				}
				err = u.system.AddUser(u.Username, options)
				if err != nil {
					status.RaiseLevel(resource.StatusFatal)
					status.Output = append(status.Output, fmt.Sprintf("error adding user %s", u.Username))
					return status, errors.Wrap(err, "user add")
				}
				status.Output = append(status.Output, fmt.Sprintf("added user %s", u.Username))
			default:
				status.RaiseLevel(resource.StatusCantChange)
				return status, fmt.Errorf("will not attempt add: user %s", u.Username)
			}
		case u.UID != "":
			_, nameNotFound := nameErr.(user.UnknownUserError)
			_, uidNotFound := uidErr.(user.UnknownUserIdError)

			switch {
			case nameNotFound && uidNotFound:
				options, err := SetAddUserOptions(u)
				if err != nil {
					status.RaiseLevel(resource.StatusCantChange)
					status.Output = append(status.Output, err.Error())
					return status, fmt.Errorf("will not attempt add: user %s with uid %s", u.Username, u.UID)
				}
				err = u.system.AddUser(u.Username, options)
				if err != nil {
					status.RaiseLevel(resource.StatusFatal)
					status.Output = append(status.Output, fmt.Sprintf("error adding user %s with uid %s", u.Username, u.UID))
					return status, errors.Wrap(err, "user add")
				}
				status.Output = append(status.Output, fmt.Sprintf("added user %s with uid %s", u.Username, u.UID))
			default:
				status.RaiseLevel(resource.StatusCantChange)
				return status, fmt.Errorf("will not attempt add: user %s with uid %s", u.Username, u.UID)
			}
		}
	case StateAbsent:
		switch {
		case u.UID == "":
			_, nameNotFound := nameErr.(user.UnknownUserError)

			switch {
			case !nameNotFound && userByName != nil:
				err := u.system.DelUser(u.Username)
				if err != nil {
					status.RaiseLevel(resource.StatusFatal)
					status.Output = append(status.Output, fmt.Sprintf("error deleting user %s", u.Username))
					return status, errors.Wrap(err, "user delete")
				}
				status.Output = append(status.Output, fmt.Sprintf("deleted user %s", u.Username))
			default:
				status.RaiseLevel(resource.StatusCantChange)
				return status, fmt.Errorf("will not attempt delete: user %s", u.Username)
			}
		case u.UID != "":
			_, nameNotFound := nameErr.(user.UnknownUserError)
			_, uidNotFound := uidErr.(user.UnknownUserIdError)

			switch {
			case !nameNotFound && !uidNotFound && userByName != nil && userByID != nil && *userByName == *userByID:
				err := u.system.DelUser(u.Username)
				if err != nil {
					status.RaiseLevel(resource.StatusFatal)
					status.Output = append(status.Output, fmt.Sprintf("error deleting user %s with uid %s", u.Username, u.UID))
					return status, errors.Wrap(err, "user delete")
				}
				status.Output = append(status.Output, fmt.Sprintf("deleted user %s with uid %s", u.Username, u.UID))
			default:
				status.RaiseLevel(resource.StatusCantChange)
				return status, fmt.Errorf("will not attempt delete: user %s with uid %s", u.Username, u.UID)
			}
		}
	default:
		status.RaiseLevel(resource.StatusFatal)
		return status, fmt.Errorf("user: unrecognized state %s", u.State)
	}

	return status, nil
}

// SetAddUserOptions returns a AddUserOptions struct with the options
// specified in the configuration for adding a user
// If group information is provided and the group is not found, nil and an
// error indicating the group name or gid is not found is returned
func SetAddUserOptions(u *User) (*AddUserOptions, error) {
	options := new(AddUserOptions)

	if u.UID != "" {
		options.UID = u.UID
	}

	switch {
	case u.GroupName != "":
		_, err := u.system.LookupGroup(u.GroupName)
		if err != nil {
			return nil, fmt.Errorf("group %s does not exist", u.GroupName)
		}
		options.Group = u.GroupName
	case u.GID != "":
		_, err := u.system.LookupGroupID(u.GID)
		if err != nil {
			return nil, fmt.Errorf("group gid %s does not exist", u.GID)
		}
		options.Group = u.GID
	}

	if u.Name != "" {
		options.Comment = u.Name
	}

	if u.HomeDir != "" {
		options.Directory = u.HomeDir
	}

	return options, nil
}
