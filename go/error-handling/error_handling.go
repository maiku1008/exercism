package erratum

const attempts = 5

// Use tries its best to frob a resource
func Use(o ResourceOpener, input string) error {
	r, err := tryToOpen(attempts, o)
	if err != nil {
		return err
	}
	defer r.Close()

	err = tryToFrob(r, input)
	if err != nil {
		return err
	}
	return nil
}

// helper function that attempts to open a resource n times,
// crashing and burning in the process
func tryToOpen(n int, o ResourceOpener) (Resource, error) {
	var r Resource
	var err error

	for i := 0; ; i++ {
		r, err = o()
		if err == nil {
			return r, nil
		}

		if i >= (n - 1) {
			break
		}

		switch err.(type) {
		case TransientError:
			continue
		default:
			return nil, err
		}
	}
	return nil, err
}

// a helper function which attempts to frob a frobbable resource
func tryToFrob(r Resource, input string) (err error) {
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
			}
		}
	}()
	r.Frob(input)

	return err
}
