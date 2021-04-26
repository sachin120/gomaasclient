package entity_test

import (
	"testing"

	. "github.com/ionutbalutoiu/gomaasclient/maas/entity"

	"github.com/ionutbalutoiu/gomaasclient/test/helper"
)

func TestResourcePoolt(t *testing.T) {
	pool := new(ResourcePool)
	pools := new([]ResourcePool)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/resource_pool.json", pool); err != nil {
		t.Fatal(err)
	}
	if err := helper.TestdataFromJSON("maas/resource_pools.json", pools); err != nil {
		t.Fatal(err)
	}
}