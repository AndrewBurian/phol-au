package scheme

type Fund struct {
	FundCode FundCode
	FundName Name
	FundID IDString
	FundDescription TextorXHTML
	Website URLorEmpty
	PreferredProviderServices []PreferredProviderServices
	AgreementHospitals []AgreementHospitals
	CorporateStructure Paragraph
	Ambulance []AmbulanceService
}

type FundCode struct {

}

type PreferredProviderServices struct {

}

type AgreementHospitals struct {

}