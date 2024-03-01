package link

type Link struct {
	Id       int64  `json:"id"`
	IdUser   int64  `json:"id_user"`
	IdClick  int64  `json:"id_click"`
	IdQrCode int64  `json:"id_qr_code"`
	Original string `json:"original"`
	Hash     string `json:"hash"`
	Ative    bool   `json:"ative"`
}

type LinkCreate struct {
	Id       int64  `json:"id"`
	IdUser   int64  `json:"id_user"`
	IdClick  int64  `json:"id_click"`
	IdQrCode int64  `json:"id_qr_code"`
	Original string `json:"original"`
	Hash     string `json:"hash"`
	Ative    bool   `json:"ative"`
}

type LinkResponse struct {
	Id       int64  `json:"id"`
	Clicks   int64  `json:"clicks"`
	QrCode   string `json:"qrCode"`
	Original string `json:"original"`
	Hash     string `json:"hash"`
	Ative    bool   `json:"ative"`
}
