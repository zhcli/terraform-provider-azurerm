package consumption

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// BillingFrequency enumerates the values for billing frequency.
type BillingFrequency string

const (
	// Month ...
	Month BillingFrequency = "Month"
	// Quarter ...
	Quarter BillingFrequency = "Quarter"
	// Year ...
	Year BillingFrequency = "Year"
)

// PossibleBillingFrequencyValues returns an array of possible values for the BillingFrequency const type.
func PossibleBillingFrequencyValues() []BillingFrequency {
	return []BillingFrequency{Month, Quarter, Year}
}

// Bound enumerates the values for bound.
type Bound string

const (
	// Lower ...
	Lower Bound = "Lower"
	// Upper ...
	Upper Bound = "Upper"
)

// PossibleBoundValues returns an array of possible values for the Bound const type.
func PossibleBoundValues() []Bound {
	return []Bound{Lower, Upper}
}

// ChargeType enumerates the values for charge type.
type ChargeType string

const (
	// ChargeTypeActual ...
	ChargeTypeActual ChargeType = "Actual"
	// ChargeTypeForecast ...
	ChargeTypeForecast ChargeType = "Forecast"
)

// PossibleChargeTypeValues returns an array of possible values for the ChargeType const type.
func PossibleChargeTypeValues() []ChargeType {
	return []ChargeType{ChargeTypeActual, ChargeTypeForecast}
}

// Datagrain enumerates the values for datagrain.
type Datagrain string

const (
	// DailyGrain Daily grain of data
	DailyGrain Datagrain = "daily"
	// MonthlyGrain Monthly grain of data
	MonthlyGrain Datagrain = "monthly"
)

// PossibleDatagrainValues returns an array of possible values for the Datagrain const type.
func PossibleDatagrainValues() []Datagrain {
	return []Datagrain{DailyGrain, MonthlyGrain}
}

// EventType enumerates the values for event type.
type EventType string

const (
	// NewCredit ...
	NewCredit EventType = "NewCredit"
	// PendingAdjustments ...
	PendingAdjustments EventType = "PendingAdjustments"
	// PendingCharges ...
	PendingCharges EventType = "PendingCharges"
	// PendingExpiredCredit ...
	PendingExpiredCredit EventType = "PendingExpiredCredit"
	// PendingNewCredit ...
	PendingNewCredit EventType = "PendingNewCredit"
	// SettledCharges ...
	SettledCharges EventType = "SettledCharges"
	// UnKnown ...
	UnKnown EventType = "UnKnown"
)

// PossibleEventTypeValues returns an array of possible values for the EventType const type.
func PossibleEventTypeValues() []EventType {
	return []EventType{NewCredit, PendingAdjustments, PendingCharges, PendingExpiredCredit, PendingNewCredit, SettledCharges, UnKnown}
}

// Grain enumerates the values for grain.
type Grain string

const (
	// Daily ...
	Daily Grain = "Daily"
	// Monthly ...
	Monthly Grain = "Monthly"
	// Yearly ...
	Yearly Grain = "Yearly"
)

// PossibleGrainValues returns an array of possible values for the Grain const type.
func PossibleGrainValues() []Grain {
	return []Grain{Daily, Monthly, Yearly}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// KindLegacy ...
	KindLegacy Kind = "legacy"
	// KindModern ...
	KindModern Kind = "modern"
	// KindUsageDetail ...
	KindUsageDetail Kind = "UsageDetail"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{KindLegacy, KindModern, KindUsageDetail}
}

// KindBasicChargeSummary enumerates the values for kind basic charge summary.
type KindBasicChargeSummary string

const (
	// KindBasicChargeSummaryKindChargeSummary ...
	KindBasicChargeSummaryKindChargeSummary KindBasicChargeSummary = "ChargeSummary"
	// KindBasicChargeSummaryKindLegacy ...
	KindBasicChargeSummaryKindLegacy KindBasicChargeSummary = "legacy"
	// KindBasicChargeSummaryKindModern ...
	KindBasicChargeSummaryKindModern KindBasicChargeSummary = "modern"
)

// PossibleKindBasicChargeSummaryValues returns an array of possible values for the KindBasicChargeSummary const type.
func PossibleKindBasicChargeSummaryValues() []KindBasicChargeSummary {
	return []KindBasicChargeSummary{KindBasicChargeSummaryKindChargeSummary, KindBasicChargeSummaryKindLegacy, KindBasicChargeSummaryKindModern}
}

// KindBasicReservationRecommendation enumerates the values for kind basic reservation recommendation.
type KindBasicReservationRecommendation string

const (
	// KindBasicReservationRecommendationKindLegacy ...
	KindBasicReservationRecommendationKindLegacy KindBasicReservationRecommendation = "legacy"
	// KindBasicReservationRecommendationKindModern ...
	KindBasicReservationRecommendationKindModern KindBasicReservationRecommendation = "modern"
	// KindBasicReservationRecommendationKindReservationRecommendation ...
	KindBasicReservationRecommendationKindReservationRecommendation KindBasicReservationRecommendation = "ReservationRecommendation"
)

// PossibleKindBasicReservationRecommendationValues returns an array of possible values for the KindBasicReservationRecommendation const type.
func PossibleKindBasicReservationRecommendationValues() []KindBasicReservationRecommendation {
	return []KindBasicReservationRecommendation{KindBasicReservationRecommendationKindLegacy, KindBasicReservationRecommendationKindModern, KindBasicReservationRecommendationKindReservationRecommendation}
}

// LookBackPeriod enumerates the values for look back period.
type LookBackPeriod string

const (
	// Last07Days Use 7 days of data for recommendations
	Last07Days LookBackPeriod = "Last7Days"
	// Last30Days Use 30 days of data for recommendations
	Last30Days LookBackPeriod = "Last30Days"
	// Last60Days Use 60 days of data for recommendations
	Last60Days LookBackPeriod = "Last60Days"
)

// PossibleLookBackPeriodValues returns an array of possible values for the LookBackPeriod const type.
func PossibleLookBackPeriodValues() []LookBackPeriod {
	return []LookBackPeriod{Last07Days, Last30Days, Last60Days}
}

// LotSource enumerates the values for lot source.
type LotSource string

const (
	// PromotionalCredit ...
	PromotionalCredit LotSource = "PromotionalCredit"
	// PurchasedCredit ...
	PurchasedCredit LotSource = "PurchasedCredit"
)

// PossibleLotSourceValues returns an array of possible values for the LotSource const type.
func PossibleLotSourceValues() []LotSource {
	return []LotSource{PromotionalCredit, PurchasedCredit}
}

// Metrictype enumerates the values for metrictype.
type Metrictype string

const (
	// ActualCostMetricType Actual cost data.
	ActualCostMetricType Metrictype = "actualcost"
	// AmortizedCostMetricType Amortized cost data.
	AmortizedCostMetricType Metrictype = "amortizedcost"
	// UsageMetricType Usage data.
	UsageMetricType Metrictype = "usage"
)

// PossibleMetrictypeValues returns an array of possible values for the Metrictype const type.
func PossibleMetrictypeValues() []Metrictype {
	return []Metrictype{ActualCostMetricType, AmortizedCostMetricType, UsageMetricType}
}

// OperatorType enumerates the values for operator type.
type OperatorType string

const (
	// EqualTo ...
	EqualTo OperatorType = "EqualTo"
	// GreaterThan ...
	GreaterThan OperatorType = "GreaterThan"
	// GreaterThanOrEqualTo ...
	GreaterThanOrEqualTo OperatorType = "GreaterThanOrEqualTo"
)

// PossibleOperatorTypeValues returns an array of possible values for the OperatorType const type.
func PossibleOperatorTypeValues() []OperatorType {
	return []OperatorType{EqualTo, GreaterThan, GreaterThanOrEqualTo}
}

// Scope11 enumerates the values for scope 11.
type Scope11 string

const (
	// Shared ...
	Shared Scope11 = "Shared"
	// Single ...
	Single Scope11 = "Single"
)

// PossibleScope11Values returns an array of possible values for the Scope11 const type.
func PossibleScope11Values() []Scope11 {
	return []Scope11{Shared, Single}
}

// Scope9 enumerates the values for scope 9.
type Scope9 string

const (
	// Scope9Shared ...
	Scope9Shared Scope9 = "Shared"
	// Scope9Single ...
	Scope9Single Scope9 = "Single"
)

// PossibleScope9Values returns an array of possible values for the Scope9 const type.
func PossibleScope9Values() []Scope9 {
	return []Scope9{Scope9Shared, Scope9Single}
}

// Term enumerates the values for term.
type Term string

const (
	// P1Y 1 year reservation term
	P1Y Term = "P1Y"
	// P3Y 3 year reservation term
	P3Y Term = "P3Y"
)

// PossibleTermValues returns an array of possible values for the Term const type.
func PossibleTermValues() []Term {
	return []Term{P1Y, P3Y}
}

// ThresholdType enumerates the values for threshold type.
type ThresholdType string

const (
	// Actual ...
	Actual ThresholdType = "Actual"
)

// PossibleThresholdTypeValues returns an array of possible values for the ThresholdType const type.
func PossibleThresholdTypeValues() []ThresholdType {
	return []ThresholdType{Actual}
}

// TimeGrainType enumerates the values for time grain type.
type TimeGrainType string

const (
	// TimeGrainTypeAnnually ...
	TimeGrainTypeAnnually TimeGrainType = "Annually"
	// TimeGrainTypeBillingAnnual ...
	TimeGrainTypeBillingAnnual TimeGrainType = "BillingAnnual"
	// TimeGrainTypeBillingMonth ...
	TimeGrainTypeBillingMonth TimeGrainType = "BillingMonth"
	// TimeGrainTypeBillingQuarter ...
	TimeGrainTypeBillingQuarter TimeGrainType = "BillingQuarter"
	// TimeGrainTypeMonthly ...
	TimeGrainTypeMonthly TimeGrainType = "Monthly"
	// TimeGrainTypeQuarterly ...
	TimeGrainTypeQuarterly TimeGrainType = "Quarterly"
)

// PossibleTimeGrainTypeValues returns an array of possible values for the TimeGrainType const type.
func PossibleTimeGrainTypeValues() []TimeGrainType {
	return []TimeGrainType{TimeGrainTypeAnnually, TimeGrainTypeBillingAnnual, TimeGrainTypeBillingMonth, TimeGrainTypeBillingQuarter, TimeGrainTypeMonthly, TimeGrainTypeQuarterly}
}
