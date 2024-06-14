//go:build !goverter

package modelfeed

//go:generate go run github.com/jmattheis/goverter/cmd/goverter@v1.4.0 gen -g ignoreUnexported .

func NewConverter() Converter {
	return &ConverterImpl{}
}
