package json

var defCase = func(pkgPath, name string) string {
	return name
}

func DefCaseFn(f func(tag string) func(pkgPath, name string) string) {
	defCase = f("json")
}
