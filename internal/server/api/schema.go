package api

type ApplySslCertSchema struct {
	Domain      string `json:"domain" binding:"required"`
	ZoneId      string `json:"zoneId" binding:"required"`
	Period      int    `json:"period" binding:"required"`
	ProjectName string `json:"projectName"`
	ProjectId   int    `json:"projectId"`
	ApplyDate   int    `json:"applyDate"`
	Owner       string `json:"owner"`
	Org         string `json:"org"`
}

type ListSslCertSchema struct {
	Domain string `form:"domain" binding:"required"`
}

type GtsCallbackSchema struct {
	Domain string `json:"domain" binding:"required"`
	Crt    string `json:"crt" binding:"required"`
	Key    string `json:"key" binding:"required"`
	Ref    string `json:"ref" binding:"required"`
	Cost   string `json:"cost" binding:"required"`
}

type GetSslPrivateKeySchema struct {
	Serial string `form:"serial" binding:"required"`
}

type CreateSanCertSchema struct {
	Domain      string `json:"domain" binding:"required"`
	Crt         string `json:"crt" binding:"required"`
	Key         string `json:"key" binding:"required"`
	Ref         string `json:"ref" binding:"required"`
	Cost        string `json:"cost" binding:"required"`
	ProjectName string `json:"projectName" binding:"required"`
	ProjectId   int    `json:"projectId" binding:"required"`
	Period      int    `json:"period" binding:"required"`
	Owner       string `json:"owner" binding:"required"`
}





