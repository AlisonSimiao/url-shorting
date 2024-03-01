package qrcode

import "url-shorting/repository"

type QrCodeRepository struct {
	repository.Repository
}

func NewQrCodeRepository() *repository.Repository {
	qcr := &QrCodeRepository{}

	qcr.Repository.Super("qr_codes")
	return &qcr.Repository
}
