package handlers

type Error struct {
	Code    string
	Message string
}

type CreateIndexReq struct {
	Name string
	Lang string
}
