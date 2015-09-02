package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasVersion(t *testing.T) {

	var has bool
	var err error

	has, err = HasVersion("tests/none")
	assert.Error(t, err)
	assert.False(t, has)

	has, err = HasVersion("tests/version_one")
	assert.NoError(t, err)
	assert.True(t, has)

}

func TestLoadVersion(t *testing.T) {

	var v *Version
	var err error

	v, err = LoadVersion("tests/no-such-dir")
	assert.Error(t, err)

	v, err = LoadVersion("tests/none")
	assert.Error(t, err)

	v, err = LoadVersion("tests/version_one")
	if assert.NoError(t, err) {
		assert.Equal(t, "v1.0.0", v.String())
	}

	v, err = LoadVersion("tests/version_two")
	if assert.NoError(t, err) {
		assert.Equal(t, "v2.4.0", v.String())
	}

}

func TestSaveVersion(t *testing.T) {

	var v *Version
	var err error

	// save a new version
	err = SaveVersion("tests/save", &Version{Major: 3, Minor: 2})

	if assert.NoError(t, err) {
		v, err = LoadVersion("tests/save")
		if assert.NoError(t, err) {
			assert.Equal(t, "v3.2.0", v.String())
		}
	}

	// save a new version
	err = SaveVersion("tests/save", &Version{Major: 4, Minor: 5, Build: 6})

	if assert.NoError(t, err) {
		v, err = LoadVersion("tests/save")
		if assert.NoError(t, err) {
			assert.Equal(t, "v4.5.6", v.String())
		}
	}

}

func TestIncreaseVersion(t *testing.T) {

	SaveVersion("tests/save", &Version{Major: 1, Minor: 0, Build: 0})
	oldVersion, newVersion, err := IncreaseVersion("tests/save", 1, 1, 1)

	if assert.NoError(t, err) {
		assert.Equal(t, "v1.0.0", oldVersion.String())
		assert.Equal(t, "v2.1.1", newVersion.String())
	}

}
