package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound      = DictionaryErr("could not find the word you were looking for")
	ErrAlreadyExists = DictionaryErr("cannot add word because it already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (definition string, err error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	if err == ErrNotFound {
		d[word] = definition
	}
	if err == nil {
		return ErrAlreadyExists
	}
	return nil
}

func (d Dictionary) Update(word, new_definition string) error {
	_, err := d.Search(word)

	if err == ErrNotFound {
		return ErrNotFound
	}

	if err == nil {
		d[word] = new_definition
	}

	return nil
}