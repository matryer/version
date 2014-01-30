package version

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	versionPrefix         string = "v"
	versionLevelSeparator string = "."
)

// Version represents the major, minor and build values.
//
// Making a new Version is as easy as:
//
// v := &Version{Major: 1, Minor: 2, Build: 3}
//
type Version struct {
	// Major represents the major version value.
	Major uint64
	// Minor represents the minor version value.
	Minor uint64
	// Build represents the build version value.
	Build uint64
}

// Parse parses a version string into a Version object.
func Parse(version string) (*Version, error) {

	v := new(Version)
	var majorStr, minorStr, buildStr string
	var err error

	// split the string
	segs := strings.Split(version, versionLevelSeparator)
	majorStr = segs[0]
	switch len(segs) {
	case 3: // v1.2.3
		buildStr = segs[2]
		fallthrough
	case 2: // v1.2
		minorStr = segs[1]
	case 1: // v1
	default: // something else unexpected
		return nil, errors.New("version: Version only supports up to three segments, major.minor.build")
	}

	if v.Major, err = parseUInt(majorStr); err != nil {
		return nil, err
	}
	if v.Minor, err = parseUInt(minorStr); err != nil {
		return nil, err
	}
	if v.Build, err = parseUInt(buildStr); err != nil {
		return nil, err
	}

	return v, nil
}

// parseUInt safely parses a string to produce a uint64.
func parseUInt(s string) (uint64, error) {

	// ignore the 'v' if there is one
	if strings.HasPrefix(s, versionPrefix) {
		s = strings.TrimLeft(s, versionPrefix)
	}

	// "" is OK - just zero
	if len(s) == 0 {
		return 0, nil
	}

	return strconv.ParseUint(s, 10, 64)
}

// Increase jumps to the next version by increasing each component by the
// specified amount.
//
// A new Version object will be returned.
func (v *Version) Increase(majorInc, minorInc, buildInc uint64) *Version {

	newV := &Version{v.Major, v.Minor, v.Build}
	newV.IncreaseHere(majorInc, minorInc, buildInc)
	return newV

}

// IncreaseHere jumps to the next version by increasing each component by the
// specified amount.  This Version object is modified.
func (v *Version) IncreaseHere(majorInc, minorInc, buildInc uint64) {
	if majorInc > 0 {
		v.Major += majorInc
		v.Minor = 0
		v.Build = 0
	}
	if minorInc > 0 {
		v.Minor += minorInc
		v.Build = 0
	}
	if buildInc > 0 {
		v.Build += buildInc
	}
}

// String gets the full version string representing this version.
func (v *Version) String() string {
	return versionPrefix + v.StringNumber()
}

// StringShort gets the short string representing this version.
// So v1.0.0 becomes v1.
func (v *Version) StringShort() string {
	return versionPrefix + v.StringShortNumber()
}

// StringNumber gets the short string representing this version
// without the prefix.  E.g. v1.1.0 becomes 1.1
func (v *Version) StringNumber() string {
	return strings.Join([]string{fmt.Sprintf("%d", v.Major),
		fmt.Sprintf("%d", v.Minor),
		fmt.Sprintf("%d", v.Build),
	}, versionLevelSeparator)
}

// StringNumber gets the short string representing this version
// without the prefix.  E.g. v1.1.0 becomes 1.1
func (v *Version) StringShortNumber() string {
	segs := make([]string, 1, 3)
	segs[0] = fmt.Sprintf("%d", v.Major)
	if v.Minor > 0 || v.Build > 0 {
		segs = append(segs, fmt.Sprintf("%d", v.Minor))
	}
	if v.Build > 0 {
		segs = append(segs, fmt.Sprintf("%d", v.Build))
	}
	return strings.Join(segs, versionLevelSeparator)
}
