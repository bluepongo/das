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
	ErrHealthcheckUpdateOperationStatus                  = 402001
	ErrHealthcheckDefaultEngineEmpty                     = 402002
	ErrHealthcheckItemWeightItemInvalid                  = 402003
	ErrHealthcheckLowWatermarkItemInvalid                = 402004
	ErrHealthcheckHighWatermarkItemInvalid               = 402005
	ErrHealthcheckUnitItemInvalid                        = 402006
	ErrHealthcheckScoreDeductionPerUnitHighItemInvalid   = 402007
	ErrHealthcheckMaxScoreDeductionHighItemInvalid       = 402008
	ErrHealthcheckScoreDeductionPerUnitMediumItemInvalid = 402009
	ErrHealthcheckMaxScoreDeductionMediumItemInvalid     = 402010
	ErrHealthcheckItemWeightSummaryInvalid               = 402011
	ErrHealthcheckPmmVersionInvalid                      = 402012
	ErrHealthcheckSQLAdvisorAdvice                       = 402013
)

func initDefaultEngineDebugMessage() {

}

func initDefaultEngineInfoMessage() {

}

func initDefaultEngineErrorMessage() {
	message.Messages[ErrHealthcheckUpdateOperationStatus] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckUpdateOperationStatus, "got error when updating operation status\n%s")
	message.Messages[ErrHealthcheckDefaultEngineEmpty] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckDefaultEngineEmpty, "default engine config should not be empty")
	message.Messages[ErrHealthcheckItemWeightItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckItemWeightItemInvalid, "item weight of %s must be in [1, 100], %d is not valid")
	message.Messages[ErrHealthcheckLowWatermarkItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckLowWatermarkItemInvalid, "low watermark of %s must be higher than 0, %f is not valid")
	message.Messages[ErrHealthcheckHighWatermarkItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckHighWatermarkItemInvalid, "high watermark of %s  must be larger than low watermark, %f is not valid")
	message.Messages[ErrHealthcheckUnitItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckUnitItemInvalid, "unit of %s must be higher than 0, %f is not valid")
	message.Messages[ErrHealthcheckScoreDeductionPerUnitHighItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckScoreDeductionPerUnitHighItemInvalid, "score deduction per unit high of %s must be in [1, 100], %f is not valid")
	message.Messages[ErrHealthcheckMaxScoreDeductionHighItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckMaxScoreDeductionHighItemInvalid, "max score deduction high of %s must be in [1, 100], %f is not valid")
	message.Messages[ErrHealthcheckScoreDeductionPerUnitMediumItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckScoreDeductionPerUnitMediumItemInvalid, "score deduction per unit medium of %s must be in [1, 100], %f is not valid")
	message.Messages[ErrHealthcheckMaxScoreDeductionMediumItemInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckMaxScoreDeductionMediumItemInvalid, "max score deduction medium of %s must be in [1, 100], %f is not valid")
	message.Messages[ErrHealthcheckItemWeightSummaryInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckItemWeightSummaryInvalid, "summary of all item weights should be 100, %d is not valid")
	message.Messages[ErrHealthcheckPmmVersionInvalid] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckPmmVersionInvalid, "pmm version should be 1 or 2, %d is not valid")
	message.Messages[ErrHealthcheckSQLAdvisorAdvice] = config.NewErrMessage(message.DefaultMessageHeader, ErrHealthcheckSQLAdvisorAdvice, "sql advisor returned error. error:\n%s")
}
