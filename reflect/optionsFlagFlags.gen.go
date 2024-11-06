package reflect

// Code generated by ../bin/flags - DO NOT EDIT.

// Returns the supplied flags status
func (o optionsFlag) GetFlag(flag optionsFlag) bool {
	return o&flag > 0
}

// Description: set to true if the hash value should be calculated by
// following pointer values rather than using the pointers value itself
//
// Used by: [StructHash]
//
// Default: true
//
//gen:enum string followPntrs
func (o optionsFlag) FollowPntrs(b bool) optionsFlag {
	if b {
		o |= followPntrs
	} else {
		o &= ^followPntrs
	}
	return o
}

// Description: set to true if the hash value should be calculated by
// following interface value
//
// Used by: [StructHash]
//
// Default: true
//
//gen:enum string followInterfaces
func (o optionsFlag) FollowInterfaces(b bool) optionsFlag {
	if b {
		o |= followInterfaces
	} else {
		o &= ^followInterfaces
	}
	return o
}

// Description: set to true if the hash value should be calculated by
// including all sub-struct fields
//
// Used by: [StructHash]
//
// Default: true
//
//gen:enum string recurseStructs
func (o optionsFlag) RecurseStructs(b bool) optionsFlag {
	if b {
		o |= recurseStructs
	} else {
		o &= ^recurseStructs
	}
	return o
}

// Description: set to true to include map key value pairs in the hash
// calculation. If false the address of the map will be used when
// calculating the hash.
//
// Used by: [StructHash]
//
// Default: true
//
//gen:enum string includeMapVals
func (o optionsFlag) IncludeMapVals(b bool) optionsFlag {
	if b {
		o |= includeMapVals
	} else {
		o &= ^includeMapVals
	}
	return o
}

// Description: set to true to include slice values in the hash calculation.
// If false the address of the slice will be used when calculating the hash.
//
// Used by: [StructHash]
//
// Default: true
//
//gen:enum string includeSliceVals
func (o optionsFlag) IncludeSliceVals(b bool) optionsFlag {
	if b {
		o |= includeSliceVals
	} else {
		o &= ^includeSliceVals
	}
	return o
}

// Description: set to true to include array values in the hash
// calculation. If false the address of the slice will be used when
// calculating the hash.
//
// Used by: [StructHash]
//
// Default: true
//
//gen:enum string includeArrayVals
func (o optionsFlag) IncludeArrayVals(b bool) optionsFlag {
	if b {
		o |= includeArrayVals
	} else {
		o &= ^includeArrayVals
	}
	return o
}
