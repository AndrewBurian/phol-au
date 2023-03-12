package scheme

import (
	"encoding/xml"
)

type ProductList struct {
	XMLName xml.Name `xml:"Products"`
	Product []*Product
	Count   string `xml:",attr"`
}

type Product struct {
	ProductType   ProductType
	ProductID     string `xml:",attr,omitempty"`
	ProductCode   string `xml:",attr"`
	ProductUrl    *URLorEmpty
	Name          string
	ProductStatus string
	State         string
	Corporate     struct {
		IsCorporate bool `xml:",attr"`
	}
	WhoIsCovered                WhoIsCovered
	Excesses                    Excesses
	CoPayments                  CoPayments `xml:",omitempty"`
	MedicareLevySurchargeExempt bool
	PremiumNoRebate             float32
	AddOns                      []*ProductAddOns
	HospitalCover               *HospitalCover
	GeneralHealthCover          *GeneralHealthCover `xml:",omitempty"`
	ProductAmbulance            struct {
		Ambulance *AmbulanceService
		UseFund   bool `xml:",attr"`
	}
}

type ProductType struct {
}

type WhoIsCovered struct {
	OnlyOnePerson string `xml:",attr"`
	Coverage      struct {
		Dependants
	}
}

type Dependants struct {
	DependantCover []struct {
		Title   string `xml:",attr"`
		Covered bool   `xml:",attr"`
	}
}

type Excesses struct {
	ExcessType         string `xml:",attr"`
	ExcessPerPolicy    Dollars
	ExcessPerPerson    Dollars
	ExcessPerAdmission Dollars
	ExcessWaivers      []struct {
		// The schema of waiver is referenced, but undocumented
		Waiver []string `xml:",innerxml"`
	}
}

type CoPayments struct {
	Shared        uint
	SharedMax     uint
	Private       uint
	PrivateMax    uint
	DaySurgery    uint
	AnnualMax     uint
	CoPaymentType string `xml:",attr"`
}

type ProductAddOns struct {
	ShortName struct {
		Value string `xml:",chardata"`
	}
}

type HospitalCover struct {
	HospitalTier TierHospital
	/*
		If Available is true this policy provides an age-based discount. As per Private Health Insurance (Reforms) Amendment Rules 2018.
		If Available is true a sequence of 1 up to a maximum of 4 non-contiguous age-ranges between the ages of 18 to 29 inclusive must be specified.
		For example a range of 18-19 and 20-21 would be deemed contiguous (and equivalent to 18-21).
		As such, maximum possible ranges would be 18-19, 21-22, 24-25, 27-28.
		If Available is true and AvailableForTransferee is true the policy is deemed to a "retained age-based discount policy" and as such provides
		that if a person transfers from another age-based discount policy (the old policy) to the policy,
		the person is entitled to retain the age-based discount (if any) they were receiving under the old policy .
	*/
	AgeBasedDiscount     AgeBasedDiscount
	Accommodation        Accommodation
	HospitalPercent      Percent
	LimitHospitalDays    Days
	MedicalServices      []MedicalService
	WaitingPeriods       []WaitingPeriods
	OtherProductFeatures string
	AccidentCover        bool
}

type GeneralHealthCover struct {
	BasedOnId bool `xml:",attr"`
	ProductPreferredProviderServices struct{
		UseFund bool `xml:",attr"`
	}
	GeneralHealthServices []*GeneralHealthService
}

type GeneralHealthService struct {
	Title string `xml:",attr"`
	Covered bool `xml:",attr"`
	HasSpecialFeatures bool `xml:,attr`

}

type AmbulanceService struct {
}

type TierHospital struct {
}

type AgeBasedDiscount struct {
}

type Accommodation struct {
}

type MedicalService struct {
}

type WaitingPeriods struct {
}
