// Code generated by "stringer -type=Access"; DO NOT EDIT.

package ast

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AccessNotSpecified-0]
	_ = x[AccessPublic-1]
	_ = x[AccessPublicSettable-2]
}

const _Access_name = "AccessNotSpecifiedAccessPublicAccessPublicSettable"

var _Access_index = [...]uint8{0, 18, 30, 50}

func (i Access) String() string {
	if i < 0 || i >= Access(len(_Access_index)-1) {
		return "Access(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Access_name[_Access_index[i]:_Access_index[i+1]]
}
