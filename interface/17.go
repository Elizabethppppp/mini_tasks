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
}) error

func Exists(s Loader, k string) bool

func Purge(s interface {
	Keyser
	Deleter
})
