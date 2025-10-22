package types

// Product related types
type AvailabilityType string

const (
	AvailabilityTypeStartTime    AvailabilityType = "START_TIME"
	AvailabilityTypeOpeningHours AvailabilityType = "OPENING_HOURS"
)

type DeliveryFormat string

const (
	DeliveryFormatPdfUrl          DeliveryFormat = "PDF_URL"
	DeliveryFormatQRCode          DeliveryFormat = "QRCODE" // Fixed: was QR_CODE
	DeliveryFormatCode128         DeliveryFormat = "CODE128"
	DeliveryFormatPkpassUrl       DeliveryFormat = "PKPASS_URL"
	DeliveryFormatAztecCode       DeliveryFormat = "AZTECCODE" // Fixed: was AZTEC_CODE
	DeliveryFormatGoogleWalletUrl DeliveryFormat = "GOOGLE_WALLET_URL"
)

type DeliveryMethod string

const (
	DeliveryMethodVoucher DeliveryMethod = "VOUCHER"
	DeliveryMethodTicket  DeliveryMethod = "TICKET"
)

type RedemptionMethod string

const (
	RedemptionMethodDigital  RedemptionMethod = "DIGITAL"
	RedemptionMethodPrint    RedemptionMethod = "PRINT"
	RedemptionMethodManifest RedemptionMethod = "MANIFEST"
)

type PricingPer string

const (
	PricingPerUnit    PricingPer = "UNIT"
	PricingPerBooking PricingPer = "BOOKING"
)

// Question related types
type InputType string

const (
	InputTypeDate          InputType = "date"
	InputTypeDatetimeLocal InputType = "datetime-local"
	InputTypeEmail         InputType = "email"
	InputTypeFile          InputType = "file"
	InputTypeNumber        InputType = "number"
	InputTypeTel           InputType = "tel"
	InputTypeText          InputType = "text"
	InputTypeTime          InputType = "time"
	InputTypeUrl           InputType = "url"
	InputTypeSelect        InputType = "select"
	InputTypeTextarea      InputType = "textarea"
)

// Option/Unit related types
type ContactField string

const (
	ContactFieldFirstName      ContactField = "firstName"      // Fixed: camelCase como en OCTO
	ContactFieldLastName       ContactField = "lastName"       // Fixed
	ContactFieldEmailAddress   ContactField = "emailAddress"   // Fixed: era EMAIL
	ContactFieldPhoneNumber    ContactField = "phoneNumber"    // Fixed
	ContactFieldCountry        ContactField = "country"        // Fixed
	ContactFieldNotes          ContactField = "notes"          // Fixed
	ContactFieldLocales        ContactField = "locales"        // Fixed
	ContactFieldAllowMarketing ContactField = "allowMarketing" // Fixed
	ContactFieldPostalCode     ContactField = "postalCode"     // Fixed
)

// Duration related types
type DurationUnit string

const (
	DurationUnitHour   DurationUnit = "hour"
	DurationUnitMinute DurationUnit = "minute"
	DurationUnitDay    DurationUnit = "day"
)

type NetDiscount string

const (
	NetDiscountNone     NetDiscount = "NONE"
	NetDiscountFull     NetDiscount = "FULL"
	NetDiscountSplit    NetDiscount = "SPLIT"
	NetDiscountProrated NetDiscount = "PRORATED"
)

type ResellerStatus string

const (
	ResellerStatusActive   ResellerStatus = "ACTIVE"
	ResellerStatusDisabled ResellerStatus = "DISABLED"
	ResellerStatusDraft    ResellerStatus = "DRAFT"
)

// Unit related types
type UnitType string

const (
	UnitTypeAdult    UnitType = "ADULT"
	UnitTypeYouth    UnitType = "YOUTH"
	UnitTypeChild    UnitType = "CHILD"
	UnitTypeInfant   UnitType = "INFANT"
	UnitTypeFamily   UnitType = "FAMILY"
	UnitTypeSenior   UnitType = "SENIOR"
	UnitTypeStudent  UnitType = "STUDENT"
	UnitTypeMilitary UnitType = "MILITARY"
	UnitTypeOther    UnitType = "OTHER"
)

// Booking related types
type BookingStatus string

const (
	BookingStatusOnHold    BookingStatus = "ON_HOLD"
	BookingStatusConfirmed BookingStatus = "CONFIRMED"
	BookingStatusExpired   BookingStatus = "EXPIRED"
	BookingStatusCancelled BookingStatus = "CANCELLED"
	BookingStatusRedeemed  BookingStatus = "REDEEMED"
	BookingStatusPending   BookingStatus = "PENDING"
	BookingStatusRejected  BookingStatus = "REJECTED"
)

// Order related types
type OrderStatus string

const (
	OrderStatusOnHold    OrderStatus = "ON_HOLD"
	OrderStatusConfirmed OrderStatus = "CONFIRMED"
	OrderStatusCancelled OrderStatus = "CANCELLED"
	OrderStatusExpired   OrderStatus = "EXPIRED"
	OrderStatusPending   OrderStatus = "PENDING"
)

// Availability related types
type AvailabilityStatus string

const (
	AvailabilityStatusAvailable AvailabilityStatus = "AVAILABLE"
	AvailabilityStatusFreesale  AvailabilityStatus = "FREESALE"
	AvailabilityStatusSoldOut   AvailabilityStatus = "SOLD_OUT"
	AvailabilityStatusLimited   AvailabilityStatus = "LIMITED"
	AvailabilityStatusClosed    AvailabilityStatus = "CLOSED"
)
