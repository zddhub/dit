package dit

type Type string // object type: object, blob, tree, commit

type object struct {
	sha1   string // or []byte
	type_  Type
	size   int64
	flags  uint
	used   bool
	parsed bool
}
