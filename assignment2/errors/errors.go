package errors

type TemporaryError struct{ Message string }

func (m TemporaryError) Error() string { return m.Message }
