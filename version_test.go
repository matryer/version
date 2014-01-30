package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var parseTests = []struct {
	versionString                               string
	expectedMajor, expectedMinor, expectedBuild uint64
}{
	{"1.2.3", 1, 2, 3},
	{"1024.2048.4096", 1024, 2048, 4096},
	{"1.2.3", 1, 2, 3},
	{"1.2", 1, 2, 0},
	{"1", 1, 0, 0},
	{"0.2.3", 0, 2, 3},
	{"0.2048.4096", 0, 2048, 4096},
	{"0.0.3", 0, 0, 3},
	{"0.2", 0, 2, 0},
	{"1", 1, 0, 0},
	{"", 0, 0, 0},
	{"v1.2.3", 1, 2, 3},
	{"v1.2.3", 1, 2, 3},
	{"v1.2", 1, 2, 0},
	{"v1", 1, 0, 0},
	{"v", 0, 0, 0},
}

func assertVersion(t *testing.T, versionString string, expectedMajor, expectedMinor, expectedBuild uint64) bool {
	v, err := Parse(versionString)
	if assert.NoError(t, err) && assert.NotNil(t, v) {
		return assert.Equal(t, v.Major, expectedMajor, versionString+" (Major)") && assert.Equal(t, v.Minor, expectedMinor, versionString+" (Minor)") && assert.Equal(t, v.Build, expectedBuild, versionString+" (Build)")
	}
	return false
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		assertVersion(t, test.versionString, test.expectedMajor, test.expectedMinor, test.expectedBuild)
	}
}

func TestNewVersion(t *testing.T) {

	v := &Version{Major: 1, Minor: 2, Build: 3}

	assert.Equal(t, v.Major, uint64(1), "Major")
	assert.Equal(t, v.Minor, uint64(2), "Minor")
	assert.Equal(t, v.Build, uint64(3), "Build")

}

func TestIncrease(t *testing.T) {

	var v *Version
	var v2 *Version

	v = &Version{Major: 1, Minor: 2, Build: 3}
	v2 = v.Increase(0, 1, 0)
	assert.Equal(t, v2.String(), "v1.3.0")

	v = &Version{Major: 1, Minor: 2, Build: 3}
	v2 = v.Increase(0, 0, 1)
	assert.Equal(t, v2.String(), "v1.2.4")

	v = &Version{Major: 1, Minor: 2, Build: 3}
	v2 = v.Increase(1, 0, 0)
	assert.Equal(t, v2.String(), "v2.0.0")

	v = &Version{Major: 1, Minor: 2, Build: 3}
	v2 = v.Increase(2, 0, 0)
	assert.Equal(t, v2.String(), "v3.0.0")

}

func TestIncreaseHere(t *testing.T) {

	var v *Version

	v = &Version{Major: 1, Minor: 2, Build: 3}
	v.IncreaseHere(0, 1, 0)
	assert.Equal(t, v.String(), "v1.3.0")

	v = &Version{Major: 1, Minor: 2, Build: 3}
	v.IncreaseHere(0, 0, 1)
	assert.Equal(t, v.String(), "v1.2.4")

	v = &Version{Major: 1, Minor: 2, Build: 3}
	v.IncreaseHere(1, 0, 0)
	assert.Equal(t, v.String(), "v2.0.0")

	v = &Version{Major: 1, Minor: 2, Build: 3}
	v.IncreaseHere(2, 0, 0)
	assert.Equal(t, v.String(), "v3.0.0")

}

func TestString(t *testing.T) {

	assert.Equal(t, (&Version{Major: 1, Minor: 2, Build: 3}).String(), "v1.2.3")
	assert.Equal(t, (&Version{Major: 1, Minor: 2}).String(), "v1.2.0")
	assert.Equal(t, (&Version{Major: 1}).String(), "v1.0.0")
	assert.Equal(t, (&Version{}).String(), "v0.0.0")

}

func TestStringShort(t *testing.T) {

	assert.Equal(t, (&Version{Major: 1, Minor: 2, Build: 3}).StringShort(), "v1.2.3")
	assert.Equal(t, (&Version{Major: 1, Minor: 0, Build: 3}).StringShort(), "v1.0.3")
	assert.Equal(t, (&Version{Major: 0, Minor: 0, Build: 3}).StringShort(), "v0.0.3")
	assert.Equal(t, (&Version{Major: 1, Minor: 2}).StringShort(), "v1.2")
	assert.Equal(t, (&Version{Major: 1}).StringShort(), "v1")
	assert.Equal(t, (&Version{}).StringShort(), "v0")

}

func TestStringShortNumber(t *testing.T) {

	assert.Equal(t, (&Version{Major: 1, Minor: 2, Build: 3}).StringShortNumber(), "1.2.3")
	assert.Equal(t, (&Version{Major: 1, Minor: 2}).StringShortNumber(), "1.2")
	assert.Equal(t, (&Version{Major: 1}).StringShortNumber(), "1")
	assert.Equal(t, (&Version{}).StringShortNumber(), "0")

}

func TestStringNumber(t *testing.T) {

	assert.Equal(t, (&Version{Major: 1, Minor: 2, Build: 3}).StringNumber(), "1.2.3")
	assert.Equal(t, (&Version{Major: 1, Minor: 2}).StringNumber(), "1.2.0")
	assert.Equal(t, (&Version{Major: 1}).StringNumber(), "1.0.0")
	assert.Equal(t, (&Version{}).StringNumber(), "0.0.0")
	assert.Equal(t, (&Version{Build: 2}).StringNumber(), "0.0.2")

}
