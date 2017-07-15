package gps

import (
	"os/exec"
	"time"

	"github.com/Masterminds/vcs"
	"github.com/golang/dep/internal/fs"
)

// PrivateRepo implements required functions to use the registry spec.
type PrivateRepo struct {
	remote, local string
}

// Vcs returns the VCS type.
func (r PrivateRepo) Vcs() vcs.Type {
	return PrivateType
}

// Remote retrieves the remote location for a repo.
func (r PrivateRepo) Remote() string {
	return r.remote
}

// LocalPath retrieves the local file system location for a repo.
func (r PrivateRepo) LocalPath() string {
	return r.local
}

// Get is used to perform an initial clone/checkout of a repository.
func (r PrivateRepo) Get() error {
	return r.Update()
}

// Init a new repository locally.
func (r PrivateRepo) Init() error {
	// does nothing
	return nil
}

// Update performs an update to an existing checkout of a repository.
func (r PrivateRepo) Update() error {
	// TODO: download latest version from REST API and write out the version name to a ".version" file.
	return nil
}

// UpdateVersion sets the version of a package of a repository.
func (r PrivateRepo) UpdateVersion(string) error {
	// TODO: download specified version from REST API and write out the version name to a ".version" file.
	return nil
}

// Version retrieves the current version.
func (r PrivateRepo) Version() (string, error) {
	// TODO: this is the latest version available from the REST API
	return "unkown", nil
}

// Current retrieves the current version-ish. This is different from the
// Version method. The output could be a branch name if on the tip of a
// branch (git), a tag if on a tag, a revision if on a specific revision
// that's not the tip of the branch. The values here vary based on the VCS.
func (r PrivateRepo) Current() (string, error) {
	// TODO: current version from ".version" file
	return "unknown", nil
}

// Date retrieves the date on the latest commit.
func (r PrivateRepo) Date() (time.Time, error) {
	// TODO: call into REST API
	return time.Now(), nil
}

// CheckLocal verifies the local location is of the correct VCS type
func (r PrivateRepo) CheckLocal() bool {
	// does nothing
	return true
}

// Branches returns a list of available branches on the repository.
func (r PrivateRepo) Branches() ([]string, error) {
	// TODO: call into REST API
	return nil, nil
}

// Tags returns a list of available tags on the repository.
func (r PrivateRepo) Tags() ([]string, error) {
	// TODO: call into REST API
	return nil, nil
}

// IsReference returns if a string is a reference. A reference can be a
// commit id, branch, or tag.
func (r PrivateRepo) IsReference(string) bool {
	// does nothing
	return true
}

// IsDirty returns if the checkout has been modified from the checked
// out reference.
func (r PrivateRepo) IsDirty() bool {
	// does nothing
	return false
}

// CommitInfo retrieves metadata about a commit.
func (r PrivateRepo) CommitInfo(string) (*vcs.CommitInfo, error) {
	// TODO: call into REST API
	return nil, nil
}

// TagsFromCommit retrieves tags from a commit id.
func (r PrivateRepo) TagsFromCommit(string) ([]string, error) {
	return []string{}, nil
}

// Ping returns if remote location is accessible.
func (r PrivateRepo) Ping() bool {
	// TODO
	return true // for now
}

// RunFromDir executes a command from repo's directory.
func (r PrivateRepo) RunFromDir(cmd string, args ...string) ([]byte, error) {
	// does nothing
	return nil, nil
}

// CmdFromDir creates a new command that will be executed from repo's
// directory.
func (r PrivateRepo) CmdFromDir(cmd string, args ...string) *exec.Cmd {
	// does nothing
	return nil
}

// ExportDir exports the current revision to the passed in directory.
func (r PrivateRepo) ExportDir(path string) error {
	return fs.CopyDir(r.local, path)
}
