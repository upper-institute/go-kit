package exceptions

import "encoding/json"

type DomainException interface {
	Code() string
}

type domainException[E DomainException] struct {
	err E
}

func ThrowException[E DomainException](err E) error {
	return &domainException[E]{err}
}

func IsException[E DomainException](exception error) (err E, ok bool) {
	err, ok = exception.(E)
	return
}

func (e *domainException[E]) Error() string {
	errJson, err := json.Marshal(e.err)
	if err != nil {
		panic(err)
	}
	return string(errJson)
}

func Catch(errGens ...func() error) error {
	for _, errGen := range errGens {
		if err := errGen(); err != nil {
			return err
		}
	}
	return nil
}
