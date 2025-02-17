package generatortests

//go:generate ../../bin/flags -type=optionsFlag -package=generatortests

type (
	optionsFlag int
)

const (
	// Description: set to true if the hash value should be calculated by
	// following pointer values rather than using the pointers value itself
	//
	// Used by: [StructHash]
	//
	// Default: true
	//gen:enum string followPntrs
	followPntrs optionsFlag = 1 << iota
	// Description: set to true if the hash value should be calculated by
	// following interface value
	//
	// Used by: [StructHash]
	//
	// Default: true
	//gen:enum string followInterfaces
	followInterfaces
	// Description: set to true if the hash value should be calculated by
	// including all sub-struct fields
	//
	// Used by: [StructHash]
	//
	// Default: true
	//gen:enum string recurseStructs
	recurseStructs
	// Description: set to true to include map key value pairs in the hash
	// calculation. If false the address of the map will be used when
	// calculating the hash.
	//
	// Used by: [StructHash]
	//
	// Default: true
	//gen:enum string includeMapVals
	includeMapVals
	// Description: set to true to include slice values in the hash calculation.
	// If false the address of the slice will be used when calculating the hash.
	//
	// Used by: [StructHash]
	//
	// Default: true
	//gen:enum string includeSliceVals
	includeSliceVals
	// Description: set to true to include array values in the hash
	// calculation. If false the address of the slice will be used when
	// calculating the hash.
	//
	// Used by: [StructHash]
	//
	// Default: true
	//gen:enum string includeArrayVals
	includeArrayVals
	//gen:flags noSetter
	//gen:enum string unknownOptionsFlag
	unknownOptionsFlag
)
