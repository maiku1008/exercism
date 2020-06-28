package erratum

import "log"

// Use tries its best to frob a resource
func Use(o ResourceOpener, input string) (err error) {
	var r Resource

	r, err = o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		r, err = o()
	}
	defer r.Close()

	defer func() {
		if rec := recover(); rec != nil {
			switch rec.(type) {
			case FrobError:
				r.Defrob(rec.(FrobError).defrobTag)
				r.Close()
				err = rec.(FrobError)
			case error:
				r.Close()
				err = rec.(error)
			default:
				log.Fatal(rec)
			}
		}
	}()
	r.Frob(input)

	return err
}
