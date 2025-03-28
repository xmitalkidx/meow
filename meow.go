package meow

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Request interface {
	GetType() string
	GetMeowVersion() string
	GetUser() string
	GetPassword() string
}

type MeowHeader struct {
	Type        string `json:"type"`
	MeowVersion string `json:"meowVersion"`
	User        string `json:"user"`
	Password    string `json:"password"`
}

func (mh MeowHeader) GetType() string {
	return mh.Type
}

func (mh MeowHeader) GetMeowVersion() string {
	return mh.MeowVersion
}

func (mh MeowHeader) GetUser() string {
	return mh.User
}

func (mh MeowHeader) GetPassword() string {
	return mh.Password
}

type RefDataRequest struct {
	MeowHeader
}

type NotificationsListRequest struct {
	MeowHeader
}

type NotificationsListResponse struct {
	MeowHeader
	Notifications []Notification
}

type RefDataResponse struct {
	Users      []string
	Currencies []string
}

type DebtsRequest struct {
	//User     string `json:"user"`
	//Password string `json:"password"`
	MeowHeader
	Currency string `json:"currency"`
	Offeree  string `json:"offeree"`
}
type DebtsResponse struct {
	DebtsList []Debt
	Stats     map[string]map[string]int64
}

type Notification struct {
	Debt
	NotificationID   string `json:"notification_id"`
	NotificationType string `json:"notification_type"`
}

type Debt struct {
	Id       string `json:"id"`
	Lender   string `json:"lender"`
	Borrower string `json:"borrower"`
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
	Date     string `json:"date"`
	Status   string `json:"status"`
	Text     string `json:"text"`
}

type NewDebtRequest struct {
	MeowHeader
	Lender   string `json:"lender"`
	Borrower string `json:"borrower"`
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
	Text     string `json:"text"`
}

type CancelDebtRequest struct {
	MeowHeader
	Lender   string `json:"lender"`
	Borrower string `json:"borrower"`
	DebtID   string `json:"debtid"`
}

func ReadAndParse(r *http.Request) (Request, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}
	defer r.Body.Close()

	var base MeowHeader
	if err := json.Unmarshal(body, &base); err != nil {
		return nil, fmt.Errorf("failed to unmarshal base message: %w", err)
	}
	switch base.GetType() {
	case "RefDataRequest":
		var data RefDataRequest
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal base message: %w", err)
		}
		log.Println(data)
		return data, nil
	case "DebtsRequest":
		var data DebtsRequest
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal base message: %w", err)
		}
		return data, nil
	case "NewDebtRequest":
		var data NewDebtRequest
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal base message: %w", err)
		}
		return data, nil
	case "CancelDebtRequest":
		var data CancelDebtRequest
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal base message: %w", err)
		}
		return data, nil
	case "NotificationsListRequest":
		var data NotificationsListRequest
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal base message: %w", err)
		}
		return data, nil

	default:
		return nil, fmt.Errorf("message of unknown type")
	}

}
