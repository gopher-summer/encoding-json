package json

type defCaseFn func(pkgPath, name string) string

var defCase = func(pkgPath, name string) string {
	return name
}

func DefCaseFn(f func(tag string) defCaseFn) {
	defCase = f("json")
}
