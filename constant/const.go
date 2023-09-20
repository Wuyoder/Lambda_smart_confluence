package constant

const (
	Query = "將以下的文件總結，並提取出8個關鍵字，回答不須要有關鍵字的解釋，回答字數越少越好，只需要回答出關鍵字，並用空格將關鍵字區隔，不需要說明與標點符號，確保最終只有8個關鍵字。"
)

var (
	GetPageContenAPI = "wiki/api/v2/pages/%s"
	PostPageLabelAPI = "wiki/rest/api/content/%s/label"
)
