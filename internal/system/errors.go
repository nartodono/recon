package system

type ResolveError struct {
	Target string
}

func (e ResolveError) Error() string {
	return "failed to resolve target: " + e.Target
}
