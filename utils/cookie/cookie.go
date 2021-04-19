package cookie

type Cookie struct {
	Path     string
	Domain   string
	Secure   bool
	HttpOnly bool
}
