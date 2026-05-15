package core

type ArcgisHubWorldCountriesGeneralizedError struct {
	IsArcgisHubWorldCountriesGeneralizedError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewArcgisHubWorldCountriesGeneralizedError(code string, msg string, ctx *Context) *ArcgisHubWorldCountriesGeneralizedError {
	return &ArcgisHubWorldCountriesGeneralizedError{
		IsArcgisHubWorldCountriesGeneralizedError: true,
		Sdk:              "ArcgisHubWorldCountriesGeneralized",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *ArcgisHubWorldCountriesGeneralizedError) Error() string {
	return e.Msg
}
