package main

import (
	ocpp16 "github.com/lorenzodonini/ocpp-go/ocpp1.6"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocppj"
	"github.com/lorenzodonini/ocpp-go/ws"
	"github.com/sirupsen/logrus"
)

const (
	chargePointID    = "test-client"
	centralSystemUrl = "ws://45.32.156.36/steve/websocket/CentralSystemService"
	clientIdTag      = "test-tag"
	connector        = 1
)

var log *logrus.Logger

func makeAvailableAndCharge(chargePoint ocpp16.ChargePoint, stateHandler *ChargePointHandler) {
	// Boot

	// Update local connector status for available

	// Notify chargepoint status available

	// Notify connector status available

	// Wait some time ...

	// Simulate charging for connector
	// Authorize

	// Update connector status for preparing

	// Start transaction

	//update the transaction id (in confirmation) locally on the status handler

	// Update connector status to charging

	// Periodically send meter values
	// for i := 0; i < 5; i++ {
	// 	sampleInterval := 5
	// 	time.Sleep(time.Second * time.Duration(sampleInterval))
	// 	stateHandler.meterValue += 10
	// 	sampledValue := types.SampledValue{Value: fmt.Sprintf("%v", stateHandler.meterValue), Unit: types.UnitOfMeasureWh, Format: types.ValueFormatRaw, Measurand: types.MeasurandEnergyActiveExportRegister, Context: types.ReadingContextSamplePeriodic, Location: types.LocationOutlet}
	// 	meterValue := types.MeterValue{Timestamp: types.NewDateTime(time.Now()), SampledValue: []types.SampledValue{sampledValue}}
	// 	meterConfirmation, err := chargePoint.MeterValues(connector, []types.MeterValue{meterValue})
	// 	if err != nil {
	// 		log.Error("error sending meter values: ", err)
	// 	}
	// 	log.Info("Confirmed : ", meterConfirmation, " meter update: ", sampledValue.Value)
	// }

	// Stop charging for connector
	//update status to finishing

	//send the stop confirmation
	// stopConfirmation, err := chargePoint.StopTransaction(stateHandler.meterValue, types.NewDateTime(time.Now()), startConfirmation.TransactionId, func(request *core.StopTransactionRequest) {
	// 	sampledValue := types.SampledValue{Value: fmt.Sprintf("%v", stateHandler.meterValue), Unit: types.UnitOfMeasureWh, Format: types.ValueFormatRaw, Measurand: types.MeasurandEnergyActiveExportRegister, Context: types.ReadingContextSamplePeriodic, Location: types.LocationOutlet}
	// 	meterValue := types.MeterValue{Timestamp: types.NewDateTime(time.Now()), SampledValue: []types.SampledValue{sampledValue}}
	// 	request.TransactionData = []types.MeterValue{meterValue}
	// 	request.Reason = core.ReasonEVDisconnected
	// })
	// if err != nil {
	// 	log.Error("error sending update status: ", err)
	// 	return
	// }
	// log.Info("Stop confirmed : ", stopConfirmation)

	//make the connector available

}

func main() {
	chargePoint := ocpp16.NewChargePoint(chargePointID, nil, nil)

	connectors := map[int]*ConnectorInfo{
		1: {status: core.ChargePointStatusAvailable, availability: core.AvailabilityTypeOperative, currentTransaction: 0},
	}
	handler := &ChargePointHandler{
		status:     core.ChargePointStatusAvailable,
		connectors: connectors,
	}

	chargePoint.SetCoreHandler(handler)
	chargePoint.SetFirmwareManagementHandler(handler)
	chargePoint.SetLocalAuthListHandler(handler)
	chargePoint.SetReservationHandler(handler)
	chargePoint.SetRemoteTriggerHandler(handler)
	chargePoint.SetSmartChargingHandler(handler)
	ocppj.SetLogger(log.WithField("logger", "ocppj"))
	ws.SetLogger(log.WithField("logger", "websocket"))

	err := chargePoint.Start(centralSystemUrl)
	if err != nil {
		log.Error("unable to start chargepoint", err)
	} else {
		log.Info("connected to central system at ", centralSystemUrl)
		makeAvailableAndCharge(chargePoint, handler)
		// Disconnect
		chargePoint.Stop()
		log.Infof("disconnected from central system")
	}
}

func init() {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	log.SetLevel(logrus.InfoLevel)
}
