package models

import (
	"time"
)

// Temperature struct
type Temperature struct {
	ID        uint   	
	Value     float32	
	SavedAt   time.Time 
}

type Humidity struct {
	ID        uint   	
	Value     float32	
	SavedAt   time.Time 
}
