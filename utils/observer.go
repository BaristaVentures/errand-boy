package utils

// Publisher has to implement the Publish method
type Publisher interface {
	Publish(value interface{})
}

// Observer is an entity which implements the Notify method.
type Observer interface {
	Notify(value interface{}) error
}

// ObserverFunc defines the type of function the observer should register.
type ObserverFunc func(value interface{}) error

// Notify is executed when an event happens.
func (fn ObserverFunc) Notify(value interface{}) error {
	return fn(value)
}

// Observers is an array of Observer instances.
type Observers []Observer

// AddObserver adds an observer to the Observers.
func (observers *Observers) AddObserver(a Observer) {
	*observers = append(*observers, a)
}

// Publish publishes allows the publisher to notify each of its subscriptors.
func (observers Observers) Publish(value interface{}) []error {
	var errs []error
	for _, obs := range observers {
		if err := obs.Notify(value); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
