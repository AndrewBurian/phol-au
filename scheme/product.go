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
	FundCode      string `xml:",attr"`
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
	OnlyOnePerson bool `xml:",attr"`
	Coverage      struct {
		Dependants
	}
}

type Dependants struct {
	NumberOfAdults int `xml:",attr"`
	DependantCover []struct {
		Title   string `xml:",attr"`
		Covered bool   `xml:",attr"`
	}
}

type Excesses struct {
	ExcessType         string `xml:",attr"`
	ExcessPerPolicy    float32
	ExcessPerPerson    float32
	ExcessPerAdmission float32
	ExcessWaivers      []string `xml:">Waiver"`
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
	HospitalTier string
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
	Accommodation        string
	HospitalPercent      float32
	LimitHospitalDays    int
	MedicalServices      []MedicalService         `xml:">MedicalService"`
	WaitingPeriods       []HospitalWaitingPeriods `xml:">WaitingPeriod"`
	OtherProductFeatures string
	AccidentCover        bool   `xml:",attr"`
	BasedOnId            string `xml:",attr"`
}

type GeneralHealthCover struct {
	BasedOnId                        bool `xml:",attr"`
	ProductPreferredProviderServices struct {
		UseFund bool `xml:",attr"`
	}
	GeneralHealthServices struct {
		XMLName                   xml.Name `xml:"GeneralHealthServices"`
		AllStatesHaveSameBenefits bool     `xml:",attr"`
		GeneralHealthService      []*GeneralHealthServiceCategory
	}
	OtherServices string          `xml:",omitempty"`
	BenefitLimits []*BenefitLimit `xml:">BenefitLimit"`
}

type GeneralHealthServices struct {
	XMLName                   xml.Name `xml:"GeneralHealthServices"`
	AllStatesHaveSameBenefits bool     `xml:",attr"`
	GeneralHealthService      []*GeneralHealthServiceCategory
}

type GeneralHealthServiceCategory struct {
	XMLName            xml.Name       `xml:"GeneralHealthService"`
	Title              string         `xml:",attr"`
	Covered            bool           `xml:",attr"`
	HasSpecialFeatures bool           `xml:",attr"`
	WaitingPeriod      *WaitingPeriod `xml:",omitempty"`
	BenefitsForState   *BenefitState  `xml:"BenefitsList,omitempty"`
}

type BenefitState struct {
	XMLName xml.Name `xml:"BenefitsList"`
	State   string   `xml:",attr"`
	Benefit []*Benefit
}

type Benefit struct {
	XMLName xml.Name `xml:"Benefit"`
	Type    string   `xml:",attr"`
	Item    string   `xml:",attr"`
	Value   float32  `xml:",chardata"`
}

type BenefitLimit struct {
	Title            string `xml:",attr"`
	ServicesCombined struct {
		IncludeOthersUnlisted bool `xml:",attr"`
		Services              []*ServiceLimit
	}
	LimitPerPerson        int
	LimitPerService       int
	LimitPerPolicy        int
	AnnualLimit           int
	LifetimeLimit         int
	ServiceCountLimit     int
	OtherLimitDescription string `xml:"FreeTextLimit"`
}

type ServiceLimit struct {
	SubLimitsApply   bool   `xml:",attr"`
	IndLifetimeLimit int    `xml:",attr"`
	ServiceName      string `xml:",chardata"`
}

type AmbulanceService struct {
}

type AgeBasedDiscount struct {
	Available              bool `xml:",attr"`
	AvailableForTransferee bool `xml:",attr"`
	AgeRange               *struct {
		MinAge int
		MaxAge int
	} `xml:",omitempty"`
}

type MedicalService struct {
	Title string `xml:",attr"`
	Cover string `xml:",attr"`
}

type WaitingPeriod struct {
	Unit   string `xml:",attr"`
	Period int    `xml:",chardata"`
}

type HospitalWaitingPeriods struct {
	Unit   string `xml:",attr"`
	Title  string `xml:",attr"`
	Period int    `xml:",chardata"`
}
