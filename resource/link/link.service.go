package link

import (
	"fmt"
	"math/rand"
	"url-shorting/repository"
	"url-shorting/resource/click"
	qrcode "url-shorting/resource/qr_code"
	rest_error "url-shorting/restError"
)

type LinkService struct {
	lr  *repository.Repository
	cs  *click.ClickService
	qcs *qrcode.QrCodeService
}

func NewLinkService() *LinkService {
	return &LinkService{
		lr:  NewLinkRepository(),
		cs:  click.NewClickService(),
		qcs: qrcode.NewQrCodeService(),
	}
}

type object map[string]interface{}

func (ls *LinkService) update(hash string, link LinkUpdate) *rest_error.Err {
	var l Link

	ls.lr.FindOne("hash = @hash", object{"hash": hash}, &l)
	if l.Id == 0 {
		return rest_error.NewNotFoundError(
			fmt.Sprintf("Link com hash '%s' não foi encontrado.", hash),
		)
	}

	if link.Original != "" {
		qrCode, err := ls.qcs.Create(link.Original)
		if qrCode.Id == 0 {
			return err
		}
		link.IdQrCode = qrCode.Id
	}

	ls.lr.Update("hash = ?", hash, link)

	return nil
}

func generateRandomString(length int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var bytes string

	tam := len(str)

	for i := 0; i < length; i++ {
		numeroAleatorio := rand.Intn(tam)
		bytes = bytes + string(str[numeroAleatorio])
	}

	return bytes
}

func createHash() string {
	return generateRandomString(6)
}

func (ls *LinkService) create(idUser int, link Link) (*LinkResponse, *rest_error.Err) {
	var hash string
	for {
		var l Link
		hash = createHash()
		ls.lr.FindOne("hash = @hash", object{"hash": hash}, &l)
		if l.Id == 0 {
			break
		}
	}

	click := ls.cs.Create()
	qrCode, errQrCode := ls.qcs.Create(link.Original)
	if errQrCode != nil {
		return nil, errQrCode
	}

	linkCreate := LinkCreate{
		IdUser:   int64(idUser),
		Hash:     hash,
		Original: link.Original,
		IdClick:  click.Id,
		IdQrCode: qrCode.Id,
		Ative:    link.Ative,
	}

	ls.lr.Create(&linkCreate)
	if linkCreate.Id == 0 {
		return nil, rest_error.NewInternalError()
	}

	return &LinkResponse{
		Id:       linkCreate.Id,
		Hash:     hash,
		Original: link.Original,
		Clicks:   click.Value,
		QrCode:   qrCode.Link,
		Ative:    link.Ative,
	}, nil
}

func (ls *LinkService) findOne(id int) *rest_error.Err {

	return nil
}
