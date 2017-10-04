// package b imports package a.
package b

import a "github.com/benmai/dep-test-a"

func B() {
	a.A()
}
