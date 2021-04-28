package endpoint_test

import (
	"testing"

	. "github.com/ionutbalutoiu/gomaasclient/api/endpoint"
	"github.com/ionutbalutoiu/gomaasclient/test/helper"
)

func TestDomaint(t *testing.T) {
	domains := new([]Domain)
	if err := helper.TestdataFromJSON("maas/domains.json", domains); err != nil {
		t.Fatal(err)
	}
}