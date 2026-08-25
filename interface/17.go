package _interface

type Storage interface {
	Save(k, v string) bool
	Load(k string) (string, bool)
	Delete(k string) bool
	Keys() []string
	Len() int
	Clear()
}

type Keyser interface {
	Keys() []string
}

type Loader interface {
	Load(k string) (string, bool)
}

type Deleter interface {
	Delete(k string) bool
}

func Backup(s interface {
	Keyser
	Loader
}) error {
	return nil
}

func Exists(s Loader, k string) bool {
	return false
}

func Purge(s interface {
	Keyser
	Deleter
}) {
}
