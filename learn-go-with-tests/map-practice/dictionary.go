package map_practice

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrExists           = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

func (d Dictionary) Add(key string, val string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = val
		return nil
	case nil:
		return ErrExists
	default:
		return err
	}
}

func (d Dictionary) Update(key string, val string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		d[key] = val
		return nil
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
