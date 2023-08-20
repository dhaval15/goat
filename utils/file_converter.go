package utils

func FileConverterRegistry() map[string]FileConverter {
	return map[string]FileConverter{
		"pdf_to_html": &PDFToHTMLConverter{},
		"image_to_html": &ImageToHTMLConverter{},
	}
}

type FileConverter interface {
	Convert(input []byte) ([]byte, error)
}

type PDFToHTMLConverter struct{}

func (c *PDFToHTMLConverter) Convert(input []byte) ([]byte, error) {
	return input, nil
}

type ImageToHTMLConverter struct{}

func (c *ImageToHTMLConverter) Convert(input []byte) ([]byte, error) {
	return input, nil
}

