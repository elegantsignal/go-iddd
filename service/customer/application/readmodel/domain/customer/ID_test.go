package customer_test

import (
	"go-iddd/service/customer/application/readmodel/domain/customer"
	"go-iddd/service/lib"
	"sync"
	"testing"

	"github.com/cockroachdb/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateID(t *testing.T) {
	Convey("When an ID is generated", t, func() {
		customerID := customer.GenerateID()

		Convey("It should succeed", func() {
			So(customerID, ShouldNotBeZeroValue)
			So(customerID, ShouldHaveSameTypeAs, customer.ID{})
			So(customerID.ID(), ShouldNotBeBlank)
		})
	})

	Convey("When many IDs are generated", t, func() {
		group := sync.WaitGroup{}
		mutex := sync.Mutex{}
		customerIDs := make(map[string]int)
		amountPerRoutine := 100
		numRoutines := 500
		totalAmount := 0

		for i := 0; i < numRoutines; i++ {
			group.Add(1)
			go generateManyIDs(customerIDs, &group, &mutex, amountPerRoutine)
			totalAmount += amountPerRoutine
		}

		group.Wait()

		Convey("They should have unique values", func() {
			So(customerIDs, ShouldHaveLength, totalAmount)
		})
	})
}

func generateManyIDs(ids map[string]int, group *sync.WaitGroup, mutex *sync.Mutex, amountPerRoutine int) {
	generatedIDs := make(map[string]int)

	for i := 0; i < amountPerRoutine; i++ {
		id := customer.GenerateID()
		generatedIDs[id.ID()] = i
	}

	mutex.Lock()
	for key, value := range generatedIDs {
		ids[key] = value
	}
	mutex.Unlock()

	group.Done()
}

func TestBuildID(t *testing.T) {
	Convey("Given that the supplied id is valid", t, func() {
		idValue := "b5f1a1b1-5d03-4e08-8365-259791228be3"

		Convey("When an ID is built", func() {
			customerID, err := customer.BuildID(idValue)

			Convey("It should succeed", func() {
				So(err, ShouldBeNil)
				So(customerID, ShouldNotBeZeroValue)
				So(customerID, ShouldHaveSameTypeAs, customer.ID{})
				So(customerID.ID(), ShouldEqual, idValue)
			})
		})
	})

	Convey("Given that the supplied id is not valid", t, func() {
		idValue := ""

		Convey("When an ID is built", func() {
			customerID, err := customer.BuildID(idValue)

			Convey("It should fail", func() {
				So(err, ShouldBeError)
				So(errors.Is(err, lib.ErrInputIsInvalid), ShouldBeTrue)
				So(customerID, ShouldBeZeroValue)
			})
		})
	})
}

func TestRebuildID(t *testing.T) {
	Convey("When an ID is rebuilt", t, func() {
		idValue := "b5f1a1b1-5d03-4e08-8365-259791228be3"
		customerID := customer.RebuildID(idValue)

		Convey("It should succeed", func() {
			So(customerID, ShouldNotBeZeroValue)
			So(customerID, ShouldHaveSameTypeAs, customer.ID{})
			So(customerID.ID(), ShouldEqual, idValue)
		})
	})
}

func TestCustomerIDEquals(t *testing.T) {
	Convey("Given an Identifier of type ID", t, func() {
		idValue := "64bcf656-da30-4f5a-b0b5-aead60965aa3"
		customerID, err := customer.BuildID(idValue)
		So(err, ShouldBeNil)

		Convey("And given an equal ID", func() {
			equalId, err := customer.BuildID(idValue)
			So(err, ShouldBeNil)

			Convey("When they are compared", func() {
				isEqual := customerID.Equals(equalId)

				Convey("They should be equal", func() {
					So(isEqual, ShouldBeTrue)
				})
			})
		})

		Convey("And given an ID with different value", func() {
			differentIdValue := "5b6e0bc9-aa69-4dd9-be1c-d54bee80f565"
			differentId, err := customer.BuildID(differentIdValue)
			So(err, ShouldBeNil)

			Convey("When they are compared", func() {
				isEqual := customerID.Equals(differentId)

				Convey("They should not be equal", func() {
					So(isEqual, ShouldBeFalse)
				})
			})
		})
	})
}