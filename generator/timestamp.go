package generator

import (
	"strconv"
	"strings"
	"time"

	"khqr/constants"
	"khqr/models"
)

type timestamp struct {
	createTime int64
	expTime    *int64
}

func NewTimestamp(v *int64) KHQRBuilder {
	return &timestamp{
		expTime:    v,
		createTime: time.Now().UnixMilli(),
	}
}

func (t *timestamp) String() string {

	if t.expTime == nil {
		return ""
	}

	expStr := strconv.FormatInt(*t.expTime, 10)
	createStr := strconv.FormatInt(t.createTime, 10)

	var subBuilder strings.Builder

	subBuilder.WriteString(models.NewTagLengthValue(constants.CreationTimestamp, &createStr).ToString())
	subBuilder.WriteString(models.NewTagLengthValue(constants.ExpirationTimestamp, &expStr).ToString())

	sub := subBuilder.String()

	if sub == "" {
		return ""
	}

	return models.NewTagLengthValue(constants.TimestampTag, &sub).ToString()
}

func (t *timestamp) Validate() error {

	if t.expTime == nil {
		return nil
	}

	if *t.expTime == 0 {
		return &constants.ErrExpirationTimestampRequired
	}

	strExpTime := strconv.FormatInt(*t.expTime, 10)

	if len(strExpTime) != constants.TimestampLength {
		return &constants.ErrExpirationTimestampLengthInvalid
	}

	if *t.expTime < t.createTime {
		return &constants.ErrExpirationTimestampInThePast
	}

	if *t.expTime < time.Now().UnixMilli() {
		return &constants.ErrKHQRExpired
	}

	return nil
}
