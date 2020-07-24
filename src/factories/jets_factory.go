package factories

import (
	"fmt"

	"github.com/bluele/factory-go/factory"

	"gorp-tips/db"
	"gorp-tips/models"
)

var JetFactory = factory.NewFactory(
	&models.Jet{},
).SeqInt("ID", func(n int) (interface{}, error) {
	return n, nil
}).Attr("Age", func(args factory.Args) (interface{}, error) {
	return uint8(20), nil
}).SeqInt("Name", func(n int) (interface{}, error) {
	return fmt.Sprintf("Jet-%d", n), nil
}).Attr("Color", func(args factory.Args) (interface{}, error) {
	return "White", nil
})

// MakeJet Jetのファクトリを作る
func MakeJet(fields Fields, deps []db.Dependency) (*models.Jet, []db.Dependency) {
	m := JetFactory.MustCreateWithOption(fields).(*models.Jet)
	if m.PilotID == 0 {
		pilot, _deps := MakePilot(nil, nil)
		m.PilotID = pilot.ID
		deps = append(deps, _deps...)
	}
	deps = append(deps, m)
	return m, deps
}
