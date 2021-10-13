package healthcheck

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initDefaultEngineDebugMessage()
	initDefaultEngineInfoMessage()
	initDefaultEngineErrorMessage()
}

const (
	// debug

	// info

	// error
	ErrHealthcheckUpdateOperationStatus       = 402001
	ErrDefaultEngineEmpty                     = 402002
	ErrItemWeightItemInvalid                  = 402003
	ErrLowWatermarkItemInvalid                = 402004
	ErrHighWatermarkItemInvalid               = 402005
	ErrUnitItemInvalid                        = 402006
	ErrScoreDeductionPerUnitHighItemInvalid   = 402007
	ErrMaxScoreDeductionHighItemInvalid       = 402008
	ErrScoreDeductionPerUnitMediumItemInvalid = 402009
	ErrMaxScoreDeductionMediumItemInvalid     = 402010
	ErrItemWeightSummaryInvalid               = 402011
	ErrPmmVersionInvalid                      = 402012
)

func initDefaultEngineDebugMessage() {

}

func initDefaultEngineInfoMessage() {

}

func initDefaultEngineErrorMessage() {
	message.Messages[ErrHealthcheckUpdateOperationStatus] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckUpdateOperationStatus, "got error when updating operation status\n%s")
	message.Messages[ErrDefaultEngineEmpty] = config.NewErrMessage(message.DefaultMessageHeader, ErrDefaultEngineEmpty, "default engine config should not be empty")
	message.Messages[ErrItemWeightItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrItemWeightItemInvalid, "item weight of %s must be in [1, 100], %d is not valid")
	message.Messages[ErrLowWatermarkItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrLowWatermarkItemInvalid, "low watermark of %s must be higher than 0, %f is not valid")
	message.Messages[ErrHighWatermarkItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrHighWatermarkItemInvalid, "high watermark of %s  must be larger than low watermark, %f is not valid")
	message.Messages[ErrUnitItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrUnitItemInvalid, "unit of %s must be higher than 0, %f is not valid")
	message.Messages[ErrScoreDeductionPerUnitHighItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrScoreDeductionPerUnitHighItemInvalid, "score deduction per unit high of %s must be in [1, 100], %f is not valid")
	message.Messages[ErrMaxScoreDeductionHighItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrMaxScoreDeductionHighItemInvalid, "max score deduction high of %s must be in [1, 100], %f is not valid")
	message.Messages[ErrScoreDeductionPerUnitMediumItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrScoreDeductionPerUnitMediumItemInvalid, "score deduction per unit medium of %s must be in [1, 100], %f is not valid")
	message.Messages[ErrMaxScoreDeductionMediumItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrMaxScoreDeductionMediumItemInvalid, "max score deduction medium of %s must be in [1, 100], %f is not valid")
	message.Messages[ErrItemWeightSummaryInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrItemWeightSummaryInvalid, "summary of all item weights should be 100, %d is not valid")
	message.Messages[ErrPmmVersionInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrPmmVersionInvalid, "pmm version should be 1 or 2, %d is not valid")
}
