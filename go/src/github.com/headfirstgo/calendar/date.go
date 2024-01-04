package calendar

import "errors"

type Date struct{
	year int
	month int
	day int
}

func(d *Date) SetYear(year int) error{
	if year < 1{
		return errors.New("Year must be > 0")
	}
	d.year = year
	return nil
}

func(d *Date) SetMonth(month int) error{
	if month <1 || month > 12{
		return errors.New("Month must be from 1 to 12")
	}
	d.month = month
	return nil
}
func(d *Date) SetDay(day int) error{
	if day < 1 || day > 31{
		return errors.New("Daq must be from 1 to 31")
	}
	d.day = day
	return nil
}
func(d *Date) Year() int{
	return d.year
}
func(d *Date) Month() int{
	return d.month
}
func(d *Date) Day() int{
	return d.day
}
