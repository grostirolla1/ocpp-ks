package main

import (
	"errors"

	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/firmware"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/localauth"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/remotetrigger"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/reservation"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/smartcharging"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/types"
)

// ConnectorInfo contains some simple state about a single connector.
type ConnectorInfo struct {
	status             core.ChargePointStatus
	availability       core.AvailabilityType
	currentTransaction int
	currentReservation int
}

// ChargePointHandler contains some simple state that a charge point needs to keep.
// In production this will typically be replaced by database/API calls.
type ChargePointHandler struct {
	status     core.ChargePointStatus
	connectors map[int]*ConnectorInfo
	meterValue int
}

func (handler *ChargePointHandler) isValidConnectorID(ID int) bool {
	_, ok := handler.connectors[ID]
	return ok || ID == 0
}

// ------------- Core profile callbacks -------------

func (handler *ChargePointHandler) OnChangeAvailability(request *core.ChangeAvailabilityRequest) (confirmation *core.ChangeAvailabilityConfirmation, err error) {
	return core.NewChangeAvailabilityConfirmation(core.AvailabilityStatusRejected), errors.New("not supported")
}

func (handler *ChargePointHandler) OnChangeConfiguration(request *core.ChangeConfigurationRequest) (confirmation *core.ChangeConfigurationConfirmation, err error) {
	return core.NewChangeConfigurationConfirmation(core.ConfigurationStatusRejected), errors.New("not supported")
}

func (handler *ChargePointHandler) OnClearCache(request *core.ClearCacheRequest) (confirmation *core.ClearCacheConfirmation, err error) {
	return core.NewClearCacheConfirmation(core.ClearCacheStatusRejected), errors.New("not supported")
}

func (handler *ChargePointHandler) OnDataTransfer(request *core.DataTransferRequest) (confirmation *core.DataTransferConfirmation, err error) {
	return core.NewDataTransferConfirmation(core.DataTransferStatusRejected), errors.New("not supported")
}

func (handler *ChargePointHandler) OnGetConfiguration(request *core.GetConfigurationRequest) (confirmation *core.GetConfigurationConfirmation, err error) {
	return core.NewGetConfigurationConfirmation(nil), errors.New("not supported")
}

func (handler *ChargePointHandler) OnRemoteStartTransaction(request *core.RemoteStartTransactionRequest) (confirmation *core.RemoteStartTransactionConfirmation, err error) {
	return core.NewRemoteStartTransactionConfirmation(types.RemoteStartStopStatusRejected), errors.New("not supported")
}

func (handler *ChargePointHandler) OnRemoteStopTransaction(request *core.RemoteStopTransactionRequest) (confirmation *core.RemoteStopTransactionConfirmation, err error) {
	return core.NewRemoteStopTransactionConfirmation(types.RemoteStartStopStatusRejected), errors.New("not supported")
}

func (handler *ChargePointHandler) OnReset(request *core.ResetRequest) (confirmation *core.ResetConfirmation, err error) {
	return core.NewResetConfirmation(core.ResetStatusRejected), errors.New("not supported")
}

func (handler *ChargePointHandler) OnUnlockConnector(request *core.UnlockConnectorRequest) (confirmation *core.UnlockConnectorConfirmation, err error) {
	return core.NewUnlockConnectorConfirmation(core.UnlockStatusNotSupported), errors.New("not supported")
}

// ------------- Local authorization list profile callbacks -------------

func (handler *ChargePointHandler) OnGetLocalListVersion(request *localauth.GetLocalListVersionRequest) (confirmation *localauth.GetLocalListVersionConfirmation, err error) {
	return localauth.NewGetLocalListVersionConfirmation(0), errors.New("not supported")
}

func (handler *ChargePointHandler) OnSendLocalList(request *localauth.SendLocalListRequest) (confirmation *localauth.SendLocalListConfirmation, err error) {
	return localauth.NewSendLocalListConfirmation(localauth.UpdateStatusNotSupported), errors.New("not supported")
}

// ------------- Firmware management profile callbacks -------------

func (handler *ChargePointHandler) OnGetDiagnostics(request *firmware.GetDiagnosticsRequest) (confirmation *firmware.GetDiagnosticsConfirmation, err error) {
	return firmware.NewGetDiagnosticsConfirmation(), errors.New("not supported")
}

func (handler *ChargePointHandler) OnUpdateFirmware(request *firmware.UpdateFirmwareRequest) (confirmation *firmware.UpdateFirmwareConfirmation, err error) {
	return firmware.NewUpdateFirmwareConfirmation(), errors.New("not supported")
}

// ------------- Remote trigger profile callbacks -------------

func (handler *ChargePointHandler) OnTriggerMessage(request *remotetrigger.TriggerMessageRequest) (confirmation *remotetrigger.TriggerMessageConfirmation, err error) {
	return remotetrigger.NewTriggerMessageConfirmation(remotetrigger.TriggerMessageStatusNotImplemented), errors.New("not supported")
}

// ------------- Reservation profile callbacks -------------

func (handler *ChargePointHandler) OnReserveNow(request *reservation.ReserveNowRequest) (confirmation *reservation.ReserveNowConfirmation, err error) {
	return reservation.NewReserveNowConfirmation(reservation.ReservationStatusRejected), errors.New("not supported")
}

func (handler *ChargePointHandler) OnCancelReservation(request *reservation.CancelReservationRequest) (confirmation *reservation.CancelReservationConfirmation, err error) {
	return reservation.NewCancelReservationConfirmation(reservation.CancelReservationStatusRejected), errors.New("not supported")
}

// ------------- Smart charging profile callbacks -------------

func (handler *ChargePointHandler) OnSetChargingProfile(request *smartcharging.SetChargingProfileRequest) (confirmation *smartcharging.SetChargingProfileConfirmation, err error) {
	return smartcharging.NewSetChargingProfileConfirmation(smartcharging.ChargingProfileStatusNotSupported), errors.New("not supported")
}

func (handler *ChargePointHandler) OnClearChargingProfile(request *smartcharging.ClearChargingProfileRequest) (confirmation *smartcharging.ClearChargingProfileConfirmation, err error) {
	return smartcharging.NewClearChargingProfileConfirmation(smartcharging.ClearChargingProfileStatusUnknown), errors.New("not supported")
}

func (handler *ChargePointHandler) OnGetCompositeSchedule(request *smartcharging.GetCompositeScheduleRequest) (confirmation *smartcharging.GetCompositeScheduleConfirmation, err error) {
	return smartcharging.NewGetCompositeScheduleConfirmation(smartcharging.GetCompositeScheduleStatusRejected), errors.New("not supported")
}
