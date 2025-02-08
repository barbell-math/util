package db

// Code generated by ../../bin/flags - DO NOT EDIT.

// Returns the supplied flags status
func (o RequiredArgs) GetFlag(flag RequiredArgs) bool {
	return o&flag > 0
}

// A flag that will be set to one if the dbUser arg is required.
func (o RequiredArgs) UserRequired(b bool) RequiredArgs {
	if b {
		o |= UserRequired
	} else {
		o &= ^UserRequired
	}
	return o
}

// A flag that will be set to one if the dbPswd arg is required.
func (o RequiredArgs) PswdRequired(b bool) RequiredArgs {
	if b {
		o |= PswdRequired
	} else {
		o &= ^PswdRequired
	}
	return o
}

// A flag that will be set to one if the dbNetLoc arg is required.
func (o RequiredArgs) NetLocRequired(b bool) RequiredArgs {
	if b {
		o |= NetLocRequired
	} else {
		o &= ^NetLocRequired
	}
	return o
}

// A flag that will be set to one if the dbPort arg is required.
func (o RequiredArgs) PortRequired(b bool) RequiredArgs {
	if b {
		o |= PortRequired
	} else {
		o &= ^PortRequired
	}
	return o
}

// A flag that will be set to one if the dbName arg is required.
func (o RequiredArgs) DBNameRequired(b bool) RequiredArgs {
	if b {
		o |= DBNameRequired
	} else {
		o &= ^DBNameRequired
	}
	return o
}
