package erratum

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
			if frob, ok := rec.(FrobError); ok {
				r.Defrob(frob.defrobTag)
			}
			err = rec.(error)
		}
	}()
	r.Frob(input)

	return err
}
