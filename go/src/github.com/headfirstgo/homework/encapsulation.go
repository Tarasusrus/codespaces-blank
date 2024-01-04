package homework

import (
	"errors"
	
)

type Coordinates struct{
	latityde float64
	longitude float64
}

func (c *Coordinates) Latityde() float64{
	return c.latityde
}

func (c *Coordinates) Longitude() float64{
	return c.longitude
}

func (c *Coordinates) SetLatitude(latityde float64) error{
	if latityde < -90 || latityde > 90{
		return errors.New("latitude mast be from -90 to 90")
	}
	c.latityde = latityde
	return nil
}

func (c *Coordinates) SetLongitude(longitude float64) error{
	if longitude < -180 || longitude > 180{
		return errors.New("longitude mast be from -90 to 90")
	}
	c.longitude = longitude
	return nil
}
