package qrcode

import (
	"url-shorting/repository"
	rest_error "url-shorting/restError"
)

type QrCodeService struct {
	qcr *repository.Repository
}

func NewQrCodeService() *QrCodeService {
	return &QrCodeService{
		qcr: NewQrCodeRepository(),
	}
}

type object map[string]interface{}

func (qs *QrCodeService) Create(url string) (QrCode, *rest_error.Err) {
	var qrCode QrCode

	qs.qcr.FindOne("link = @link", object{"link": url}, &qrCode)
	if qrCode.Id != 0 {
		return qrCode, nil
	}
	qrCode.Link = url

	qs.qcr.Create(&qrCode)

	if qrCode.Id == 0 {
		return QrCode{}, rest_error.NewInternalError()
	}

	return qrCode, nil
}
