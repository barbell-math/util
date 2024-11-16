// Implmenents all the tests that verify the containers in the [container] package
// properly implmenent the interfaces defined in the [dynamicContainers] and
// [staticContainers] packages. These tests only operate on the interface values
// that are defined in the two afformentioned packages, meaning they only test
// the high level functionality of each collection. For tests that require
// knowing details internal to each collection refer to the tests in the
// [containers] package.
//
// The tests in this package are not standard go test tests. The functions in
// this package are called from templated test functions that were generated
// using go:generate in the [containers] package. This allows for many container
// types to be tested against the same set of interface functionality.
//
// Due to this package only testing interface values many tests will rely on
// multiple functions from the interface working correctly.
//
// There are several conventions that must be observed when writing these tests:
//  1. All values placed in the tested containers must be of type int. This
//     facilitates populating values in the containers as it can simply be done
//     with standard for loops.
package tests
