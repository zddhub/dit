package dit

func (repo *repository) CatFile(sha1 string) (*object, []byte, error) {
	object := &object{sha1: sha1}
	buffer, err := object.ReadAll()
	return object, buffer, err
}
