package version

import (
	"io/ioutil"
	"os"
	"path"
)

const (
	versionFilename string = ".version"
)

// IncreaseVersion jumps to the next version by increasing each component by the
// specified amount and updates the directory.
//
// Returns the old version, the new version and any errors that may have occurred
// while reading and writing files.
func IncreaseVersion(directory string, majorInc, minorInc, buildInc uint64) (oldVersion *Version, newVersion *Version, err error) {

	oldVersion, err = LoadVersion(directory)
	if err != nil {
		return nil, nil, err
	}
	newVersion = oldVersion.Increase(majorInc, minorInc, buildInc)
	err = SaveVersion(directory, newVersion)
	return

}

// HasVersion gets whether the specified directory has a
// version or not.
func HasVersion(directory string) (bool, error) {

	var err error

	// make sure the directory exists
	_, err = os.Stat(directory)
	if os.IsNotExist(err) {
		return false, err
	}

	_, err = os.Stat(getVersionFilename(directory))
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil

}

// LoadVersion loads the version for the specified directory.
func LoadVersion(directory string) (*Version, error) {

	// if no version - use default
	has, err := HasVersion(directory)
	if err != nil {
		return nil, err
	}
	if !has {
		return &Version{}, nil
	}

	verBytes, err := ioutil.ReadFile(getVersionFilename(directory))
	if err != nil {
		return nil, err
	}

	return Parse(string(verBytes))

}

// SaveVersion saves a new version for the specified directory.
func SaveVersion(directory string, version *Version) error {
	return ioutil.WriteFile(getVersionFilename(directory), []byte(version.String()), 0644)
}

// getVersionFilename gets the filename of the version file for the
// specified directory.
func getVersionFilename(directory string) string {
	return path.Join(directory, versionFilename)
}
