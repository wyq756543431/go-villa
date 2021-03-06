package villa

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// Path is a wrapper for a path in the OS.
// Some commonly used functions are wrapped as methods of Path,
// and results, if any, are converted back to Path
type Path string

// Join connect elems to the tail of path
func (p Path) Join(elem ...interface{}) Path {
	els := make(StringSlice, 0, len(elem)+1)
	els.Add(p).Add(elem...)
	return Path(filepath.Join(els...))
}

// Exists checks whether the file exists
func (p Path) Exists() bool {
	_, err := p.Stat()
	return err == nil
}

// IsDir returns true only if the path exists and indicates a directory
func (p Path) IsDir() bool {
	info, err := p.Stat()
	if err != nil {
		// the path does not exist
		return false
	}
	return info.Mode().IsDir()
}

// S converts Path back to string. This is sometimes more concise than string(p)
func (p Path) S() string {
	return string(p)
}

// AbsPath returns the absolute path returned by Path.Abs() if no error found,
// panic otherwise.
func (p Path) AbsPath() (ap Path) {
	ap, err := p.Abs()
	if err != nil {
		panic(err)
	}
	return ap
}

/*
	wrappers of path/filepath package
*/

// Abs is a wrapper to filepath.Abs
// It returns an absolute representation of path. If the path is not absolute it
// will be joined with the current working directory to turn it into an absolute
// path. The absolute path name for a given file is not guaranteed to be unique.
func (p Path) Abs() (pth Path, err error) {
	pt, err := filepath.Abs(string(p))
	return Path(pt), err
}

// Base returns the last element of path.
// Trailing path separators are removed before extracting the last element.
// If the path is empty, Base returns ".".
// If the path consists entirely of separators, Base returns a single separator.
func (p Path) Base() Path {
	return Path(filepath.Base(string(p)))
}

// Clean is a wrapper to filepath.Clean.
//
// It returns the shortest path name equivalent to path by purely lexically
// processing.
func (p Path) Clean() Path {
	return Path(filepath.Clean(string(p)))
}

// Dir is a wrapper to filepath.Dir.
// It returns all but the last element of path, typically the path's directory.
// Trailing path separators are removed before processing. If the path is empty,
// Dir returns ".". If the path consists entirely of separators, Dir returns a
// single separator. The returned path does not end in a separator unless it is
// the root directory.
func (p Path) Dir() Path {
	return Path(filepath.Dir(string(p)))
}

// EvalSymlinks is a wrapper to filepath.EvalSymlinks
func (p Path) EvalSymlinks() (Path, error) {
	pt, err := filepath.EvalSymlinks(string(p))
	return Path(pt), err
}

// Ext is a wrapper to filepath.Ext
func (p Path) Ext() string {
	return filepath.Ext(string(p))
}

// FromSlash is a wrapper to filepath.FromSlash
func (p Path) FromSlash() Path {
	return Path(filepath.FromSlash(string(p)))
}

// IsAbs is a wrapper to filepath.IsAbs
//
// IsAbs returns true if the path is absolute.
func (p Path) IsAbs() bool {
	return filepath.IsAbs(string(p))
}

// Rel is a wrapper to filepath.Rel.
// It returns a relative path that is lexically equivalent to targpath when
// joined to basepath with an intervening separator. That is, Join(basepath,
// Rel(basepath, targpath)) is equivalent to targpath itself. On success, the
// returned path will always be relative to basepath, even if basepath and
// targpath share no elements. An error is returned if targpath can't be made
// relative to basepath or if knowing the current working directory would be
// necessary to compute it.
func (p Path) Rel(targetpath Path) (Path, error) {
	rel, err := filepath.Rel(string(p), string(targetpath))
	return Path(rel), err
}

// Split is a wrapper to filepath.Split
func (p Path) Split() (dir, file Path) {
	d, f := filepath.Split(string(p))
	return Path(d), Path(f)
}

// SplitList is a wrapper to filepath.SplitList
func (p Path) SplitList() (lst []Path) {
	l := filepath.SplitList(string(p))
	lst = make([]Path, len(l))
	for i, el := range l {
		lst[i] = Path(el)
	}
	return
}

// ToSlash is a wrapper to filepath.ToSlash
func (p Path) ToSlash() string {
	return filepath.ToSlash(string(p))
}

// VolumeName is a wrapper to filepath.VolumeName
func (p Path) VolumeName() string {
	return filepath.VolumeName(string(p))
}

// WalkFunc is a wrapper to filepath.WalkFunc
// WalkFunc is the type of the function called for each file or directory
// visited by Walk. The path argument contains the argument to Walk as a prefix;
// that is, if Walk is called with "dir", which is a directory containing the
// file "a", the walk function will be called with argument "dir/a". The info
// argument is the os.FileInfo for the named path.
//
// If there was a problem walking to the file or directory named by path, the
// incoming error will describe the problem and the function can decide how to
// handle that error (and Walk will not descend into that directory). If an
// error is returned, processing stops. The sole exception is that if path is a
// directory and the function returns the special value filepath.SkipDir, the
// contents of the directory are skipped and processing continues as usual on
// the next file.
type WalkFunc func(path Path, info os.FileInfo, err error) error

// Walk is a wrapper to filepath.Walk
// Walk walks the file tree rooted at root, calling walkFn for each file or
// directory in the tree, including root. All errors that arise visiting files
// and directories are filtered by walkFn. The files are walked in lexical
// order, which makes the output deterministic but means that for very large
// directories Walk can be inefficient. Walk does not follow symbolic links.
func (p Path) Walk(walkFn WalkFunc) error {
	return filepath.Walk(string(p), func(path string, info os.FileInfo, err error) error {
		return walkFn(Path(path), info, err)
	})
}

/*
	wrappers of os package
*/

// Create is a wrapper to os.Create.
//
// It creates the named file mode 0666 (before umask), truncating it if it 
// already exists. If successful, methods on the returned File can be used for 
// I/O; the associated file descriptor has mode O_RDWR. If there is an error, 
// it will be of type *PathError.
func (p Path) Create() (file *os.File, err error) {
	return os.Create(string(p))
}

// Open is a wrapper to os.Open.
//
// It opens the named file for reading. If successful, methods on the returned
// file can be used for reading; the associated file descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.
func (p Path) Open() (file *os.File, err error) {
	return os.Open(string(p))

}

// OpenFile is a wrapper to os.OpenFile.
//
// It is the generalized open call; most users will use Open or Create instead.
// It opens the named file with specified flag (O_RDONLY etc.) and perm,
// (0666 etc.) if applicable. If successful, methods on the returned File can be
// used for I/O. If there is an error, it will be of type *PathError.
func (p Path) OpenFile(flag int, perm os.FileMode) (file *os.File, err error) {
	return os.OpenFile(string(p), flag, perm)
}

// Mkdir is a wrappter to os.Mkdir.
//
// It creates a new directory with the specified name and permission bits. If 
// there is an error, it will be of type *PathError.
func (p Path) Mkdir(perm os.FileMode) error {
	return os.Mkdir(string(p), perm)
}

// MkdirAll is a wrappter to os.MkdirAll.
//
// It creates a directory named path, along with any necessary parents, and 
// returns nil, or else returns an error. The permission bits perm are used for 
// all directories that MkdirAll creates. If path is already a directory, 
// MkdirAll does nothing and returns nil.
func (p Path) MkdirAll(perm os.FileMode) error {
	return os.MkdirAll(string(p), perm)
}

// Remove is a wrappter to os.Remove
//
// Remove removes the named file or directory. If there is an error, it will be
// of type *PathError.
func (p Path) Remove() error {
	return os.Remove(string(p))
}

// RemoveAll is a wrappter to os.RemoveAll
// RemoveAll removes path and any children it contains. It removes everything it
// can but returns the first error it encounters. If the path does not exist,
// RemoveAll returns nil (no error).
func (p Path) RemoveAll() error {
	return os.RemoveAll(string(p))
}

// Rename is a wrappter to os.Rename
func (p Path) Rename(newname Path) error {
	return os.Rename(string(p), string(newname))
}

// Stat is a wrappter to os.Stat
func (p Path) Stat() (fi os.FileInfo, err error) {
	return os.Stat(string(p))
}

// Symlink is a wrappter to os.Symlink
func (p Path) Symlink(dst Path) error {
	return os.Symlink(string(p), string(dst))
}

/*
	wrappers of ioutil package
*/

// ReadDir is a wrappter to ioutil.ReadDir
//
// ReadDir reads the directory named by dirname and returns a list of sorted
// directory entries.
func (p Path) ReadDir() (fi []os.FileInfo, err error) {
	return ioutil.ReadDir(string(p))
}

// ReadFile is a wrappter to ioutil.ReadFile
func (p Path) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(string(p))
}

// WriteFile is a wrappter to ioutil.WriteFile
func (p Path) WriteFile(data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(string(p), data, perm)
}

// TempDir is a wrappter to ioutil.TempDir.
// It creates a new temporary directory in the directory dir with a name
// beginning with prefix and returns the path of the new directory. If dir is
// the empty string, TempDir uses the default directory for temporary files
// (see os.TempDir). Multiple programs calling TempDir simultaneously will not
// choose the same directory. It is the caller's responsibility to remove the
// directory when no longer needed.
func (p Path) TempDir(prefix string) (name Path, err error) {
	nm, err := ioutil.TempDir(string(p), prefix)
	return Path(nm), err
}

/*
	wrapppers of exec package
*/

// Command is a wrappter to exec.Command
func (p Path) Command(arg ...string) *exec.Cmd {
	return exec.Command(string(p), arg...)
}
