package cmd

import (
	"database/sql"
	"go-iddd/service/customer/application/readmodel"
	"go-iddd/service/customer/application/writemodel"
	customercli "go-iddd/service/customer/infrastructure/primary/cli"
	customergrpc "go-iddd/service/customer/infrastructure/primary/grpc"
	eventstoreForReadModel "go-iddd/service/customer/infrastructure/secondary/forreadingcustomereventstreams/eventstore"
	eventstoreForWriteModel "go-iddd/service/customer/infrastructure/secondary/forstoringcustomerevents/eventstore"
	"go-iddd/service/lib"
	"go-iddd/service/lib/es"
	"go-iddd/service/lib/eventstore/postgres"

	"github.com/cockroachdb/errors"
)

const eventStoreTableName = "eventstore"

type DIContainer struct {
	postgresDBConn                    *sql.DB
	unmarshalDomainEventForWriteModel es.UnmarshalDomainEvent
	unmarshalDomainEventForReadModel  es.UnmarshalDomainEvent
	eventStoreForWriteModel           *postgres.EventStore
	eventStoreForReadModel            *postgres.EventStore
	customerEventStoreForWriteModel   *eventstoreForWriteModel.CustomerEventStore
	customerEventStoreForReadModel    *eventstoreForReadModel.CustomerEventStore
	customerCommandHandler            *writemodel.CustomerCommandHandler
	customerQueryHandler              *readmodel.CustomerQueryHandler
	customerServer                    customergrpc.CustomerServer
	customerApp                       *customercli.CustomerApp
}

func NewDIContainer(
	postgresDBConn *sql.DB,
	unmarshalDomainEventForWriteModel es.UnmarshalDomainEvent,
	unmarshalDomainEventForReadModel es.UnmarshalDomainEvent,
) (*DIContainer, error) {

	if postgresDBConn == nil {
		return nil, errors.Mark(errors.New("newContainer: postgres DB connection must not be nil"), lib.ErrTechnical)
	}

	container := &DIContainer{
		postgresDBConn:                    postgresDBConn,
		unmarshalDomainEventForWriteModel: unmarshalDomainEventForWriteModel,
		unmarshalDomainEventForReadModel:  unmarshalDomainEventForReadModel,
	}

	container.init()

	return container, nil
}

func (container DIContainer) init() {
	container.getEventStoreForWriteModel()
	container.getEventStoreForReadModel()
	container.GetCustomerEventStoreForWriteModel()
	container.GetCustomerEventStoreForReadModel()
	container.GetCustomerCommandHandler()
	container.GetCustomerQueryHandler()
	container.GetCustomerServer()
	container.GetCustomerApp()
}

func (container DIContainer) GetPostgresDBConn() *sql.DB {
	return container.postgresDBConn
}

func (container DIContainer) getEventStoreForWriteModel() *postgres.EventStore {
	if container.eventStoreForWriteModel == nil {
		container.eventStoreForWriteModel = postgres.NewEventStore(
			container.postgresDBConn,
			eventStoreTableName,
			container.unmarshalDomainEventForWriteModel,
		)
	}

	return container.eventStoreForWriteModel
}

func (container DIContainer) getEventStoreForReadModel() *postgres.EventStore {
	if container.eventStoreForReadModel == nil {
		container.eventStoreForReadModel = postgres.NewEventStore(
			container.postgresDBConn,
			eventStoreTableName,
			container.unmarshalDomainEventForReadModel,
		)
	}

	return container.eventStoreForReadModel
}

func (container DIContainer) GetCustomerEventStoreForWriteModel() *eventstoreForWriteModel.CustomerEventStore {
	if container.customerEventStoreForWriteModel == nil {
		container.customerEventStoreForWriteModel = eventstoreForWriteModel.NewCustomerEventStore(
			container.getEventStoreForWriteModel(),
		)
	}

	return container.customerEventStoreForWriteModel
}

func (container DIContainer) GetCustomerEventStoreForReadModel() *eventstoreForReadModel.CustomerEventStore {
	if container.customerEventStoreForReadModel == nil {
		container.customerEventStoreForReadModel = eventstoreForReadModel.NewCustomerEventStore(
			container.getEventStoreForReadModel(),
		)
	}

	return container.customerEventStoreForReadModel
}

func (container DIContainer) GetCustomerCommandHandler() *writemodel.CustomerCommandHandler {
	if container.customerCommandHandler == nil {
		container.customerCommandHandler = writemodel.NewCustomerCommandHandler(
			container.GetCustomerEventStoreForWriteModel(),
		)
	}

	return container.customerCommandHandler
}

func (container DIContainer) GetCustomerQueryHandler() *readmodel.CustomerQueryHandler {
	if container.customerQueryHandler == nil {
		container.customerQueryHandler = readmodel.NewCustomerQueryHandler(
			container.GetCustomerEventStoreForReadModel(),
		)
	}

	return container.customerQueryHandler
}

func (container DIContainer) GetCustomerServer() customergrpc.CustomerServer {
	if container.customerServer == nil {
		container.customerServer = customergrpc.NewCustomerServer(
			container.GetCustomerCommandHandler().RegisterCustomer,
			container.GetCustomerCommandHandler().ConfirmCustomerEmailAddress,
			container.GetCustomerCommandHandler().ChangeCustomerEmailAddress,
		)
	}

	return container.customerServer
}

func (container DIContainer) GetCustomerApp() *customercli.CustomerApp {
	if container.customerApp == nil {
		container.customerApp = customercli.NewCustomerApp(
			container.GetCustomerCommandHandler().RegisterCustomer,
			container.GetCustomerCommandHandler().ConfirmCustomerEmailAddress,
			container.GetCustomerCommandHandler().ChangeCustomerEmailAddress,
		)
	}

	return container.customerApp
}
