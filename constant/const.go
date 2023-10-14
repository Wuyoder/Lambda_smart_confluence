package constant

const (
	Query = "Extract 8 keywords from the following document. Answer with only the keywords separated by spaces, without any explanations or punctuation marks, to ensure that there are only 8 keywords."
)

var (
	GetPageContenAPI = "wiki/api/v2/pages/%s"
	PostPageLabelAPI = "wiki/rest/api/content/%s/label"
)
