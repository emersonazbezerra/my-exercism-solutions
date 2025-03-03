package elon

import "fmt"

func (c *Car) Drive() {
	if c.battery-c.batteryDrain < 0 {
		return
	}
	c.battery -= c.batteryDrain
	c.distance += c.speed
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c *Car) CanFinish(trackDistance int) bool {
	tdf, sf, bdf, bf := float32(trackDistance), float32(c.speed), float32(c.batteryDrain), float32(c.battery)
	return tdf/sf*bdf <= bf
}
