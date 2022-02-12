package reservation

const (
	RoleDeveloper = 1
	RoleAdmin     = 2
	RoleCustomer  = 3
	RolePIC       = 4
	RoleOwner     = 5
	RoleUser      = 6
)

// HstReservation constat
const (
	ReasonNewReservation = "New Reservation"
)

// reservation status constant
const (
	StatusRegistered              = 1
	StatusAccepted                = 2
	StatusProcessing              = 3
	StatusFinished                = 4
	StatusDone                    = 5
	StatusRejected                = 6
	StatusWaitingCustomerApproval = 7
	StatusCancelledByShop         = 8
	StatusCancelledByCustomer     = 9
)

var (
	MapAllowedStatusProgress = map[int][]int{
		StatusRegistered: {
			StatusRejected,
			StatusCancelledByShop,
			StatusCancelledByCustomer,
		},
		StatusAccepted: {
			StatusRegistered,
		},
		StatusProcessing: {
			StatusAccepted,
			StatusWaitingCustomerApproval,
		},
		StatusFinished: {
			StatusProcessing,
		},
		StatusDone: {
			StatusFinished,
		},
		StatusRejected: {
			StatusRegistered,
			StatusAccepted,
		},
		StatusWaitingCustomerApproval: {
			StatusProcessing,
			StatusFinished,
		},
		StatusCancelledByShop: {
			StatusProcessing,
		},
		StatusCancelledByCustomer: {
			StatusRegistered,
			StatusAccepted,
		},
	}
)
