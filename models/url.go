package models

/*
this file define the DSA so i the system is like that
Long Link--req---> [ Shortener ] --res---> Short Link
*/
type ShortenerRequest struct {
	URL string `json:"url"`
}

type ShortenerResponse struct {
	ShortURL string `json:"short_url"`
}
