package model

type DefaultEvent[T interface{}] struct {
	Key  string
	Data T
}

// account created

func AccountCreated(data AccountCreatedEventData) AccountCreatedEvent {
	return AccountCreatedEvent{
		DefaultEvent[AccountCreatedEventData]{
			Key:  "ACCOUNT_CREATED",
			Data: data,
		},
	}
}

type AccountCreatedEvent struct {
	DefaultEvent[AccountCreatedEventData]
}

type AccountCreatedEventData struct {
}

// account updated

func AccountUpdated(data AccountUpdatedEventData) AccountUpdatedEvent {
	return AccountUpdatedEvent{
		DefaultEvent[AccountUpdatedEventData]{
			Key:  "ACCOUNT_UPDATED",
			Data: data,
		},
	}
}

type AccountUpdatedEvent struct {
	DefaultEvent[AccountUpdatedEventData]
}

type AccountUpdatedEventData struct {
}
