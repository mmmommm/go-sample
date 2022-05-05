package handler

func ProvideHandler(
	userGet UserGetHandler,
) *Handler {
	return &Handler{
		UserGetHandler: userGet,
	}
}

type Handler struct {
	UserGetHandler UserGetHandler
}