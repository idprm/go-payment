package callback

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/idprm/go-payment/src/domain/entity"
	"github.com/idprm/go-payment/src/logger"
	"github.com/sirupsen/logrus"
)

type Callback struct {
	logger *logger.Logger
	app    *entity.Application
	order  *entity.Order
}

func NewCallback(
	logger *logger.Logger,
	app *entity.Application,
	order *entity.Order,
) *Callback {
	return &Callback{
		logger: logger,
		app:    app,
		order:  order,
	}
}

func (p *Callback) Hit() ([]byte, error) {
	l := p.logger.Init("callback", true)

	start := time.Now()

	// set transaction_id
	id := uuid.New()
	trxId := id.String()

	request := &entity.CallbackRequestBody{
		Number: p.order.GetNumber(),
		IsPaid: true,
		Time:   time.Now(),
	}
	payload, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", p.app.GetUrlCallback(), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		Proxy:              http.ProxyFromEnvironment,
		MaxIdleConns:       10,
		IdleConnTimeout:    60 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   60 * time.Second,
		Transport: tr,
	}

	p.logger.Writer(req)
	l.WithFields(logrus.Fields{"request": string(payload), "trx_id": trxId}).Info("CALLBACK_REQUEST")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	duration := time.Since(start).Milliseconds()
	p.logger.Writer(string(body))
	l.WithFields(logrus.Fields{"response": string(body), "trx_id": trxId, "duration": duration}).Info("CALLBACK_RESPONSE")

	return body, nil
}
