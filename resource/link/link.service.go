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

func (ls *LinkService) updateClick(hash string) *rest_error.Err {
	var link Link

	ls.lr.FindOne("hash = @hash", object{"hash": hash}, &link)
	if link.Id == 0 {
		return rest_error.NewNotFoundError(
			fmt.Sprintf("Link com hash '%s' não foi encontrado.", hash),
		)
	}

	err := ls.cs.AddClick(link.IdClick)
	if err != nil {
		return err
	}
	return nil
}

func (ls *LinkService) findOne(hash string) (*LinkResponse, *rest_error.Err) {

	var link LinkResponse
	ls.lr.FindOneWithJoin(`
	ORIGINAL,
	links.ID,
  	HASH,
  	ATIVE,
	clicks.VALUE AS clicks,
  	qr_codes.link AS qr_code
	`,
	`
	LEFT JOIN clicks ON links.id_click = clicks.id
	LEFT JOIN qr_codes ON links.id_qr_code = qr_codes.id`,
		"Hash=?", hash, &link)

	if link.Id == 0 {
		return nil, rest_error.NewNotFoundError(
			fmt.Sprintf("Link com Hash '%s' não foi encontrado.", hash),
		)
	}

	return &link, nil
}

func (ls *LinkService) findAll(idUser int, page int, limit int) (*repository.PaginateData, *rest_error.Err) {
	//trazer os links do banco e colocar no array
	var links []LinkResponse
	
	res := ls.lr.PaginateWithJoin(`
	ORIGINAL,
  	HASH,
  	ATIVE,
	clicks.VALUE AS clicks,
  	qr_codes.link AS qr_code
	`,
		`
	LEFT JOIN clicks ON links.id_click = clicks.id
	LEFT JOIN qr_codes ON links.id_qr_code = qr_codes.id`,
		`id_user = ?`, idUser, &links, page, limit)

	return res, nil

}
