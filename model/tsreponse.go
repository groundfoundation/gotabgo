package model

import "encoding/xml"

// TsResponse is the wrapper that Tableau Server wraps each response with
type TsResponse struct {
	XMLName    xml.Name    `json:"-"            xml:"http://tableau.com/api tsResponse"`
	ServerInfo ServerInfo  `json:"serverInfo"   xml:"serverInfo"`
	Workbooks  []Workbooks `json:"workbooks"   xml:"workbooks"`
	Users      struct {
		User []User `json:"user"   xml:"user"`
	} `json:"users"   xml:"users"`
	Credentials Credentials `json:"credentials"  xml:"credentials"`
	Error       ErrorType   `json:"error"        xml:"error"`
	Site        SiteType    `json:"site"         xml:"site"`
}

// ServerInfo contains information about product version and api version for the server
type ServerInfo struct {
	XMLName        xml.Name       `json:"-" xml:"serverInfo"`
	ProductVersion ProductVersion `json:"productVersion,omitempty" xml:"productVersion"`
	RestApiVersion string         `json:"restApiVersion,omitempty" xml:"restApiVersion,omitempty"`
}

type SiteType struct {
	XMLName      xml.Name `json:"-"                       xml:"site"`
	ID           string   `json:"id,omitempty"            xml:"id,attr,omitempty"`
	Name         string   `json:"name,omitempty"          xml:"name,attr,omitempty"`
	ContentUrl   string   `json:"contentUrl,omitempty"    xml:"contentUrl,attr,omitempty"`
	AdminMode    string   `json:"adminMode,omitempty"     xml:"adminMode,attr,omitempty"`
	UserQuota    string   `json:"userQuota,omitempty"     xml:"userQuota,attr,omitempty"`
	StorageQuota int      `json:"storageQuota,omitempty"  xml:"storageQuota,attr,omitempty"`
	State        string   `json:"state,omitempty"         xml:"state,attr,omitempty"`
	StatusReason string   `json:"statusReason,omitempty"  xml:"statusReason,attr,omitempty"`
	Usage        struct {
		NumUsers     uint `json:"numUsers"                xml:"numUsers,attr"`
		NumCreators  uint `json:"numCreators,omitempty"   xml:"numCreators,omitempty,attr"`
		NumExplorers uint `json:"numExplorers,omitempty"  xml:"numExplorers,omitempty,attr"`
		NumViewers   uint `json:"numViewers,omitempty"    xml:"numViewers,omitempty,attr"`
		Storage      uint `json:"storage"                 xml:"storage,attr"`
	} `json:"usage,omitempty"  xml:"usage,omitempty"`
}

type Credentials struct {
	XMLName     xml.Name  `json:"-"                   xml:"credentials"`
	Name        string    `json:"name,omitempty"      xml:"name,attr,omitempty"`
	Password    string    `json:"password,omitempty"  xml:"password,attr,omitempty"`
	Token       string    `json:"token,omitempty"     xml:"token,attr,omitempty"`
	Site        *SiteType `json:"site,omitempty"      xml:"site,omitempty"`
	Impersonate *User     `json:"user,omitempty"      xml:"user,omitempty"`
}
type ProductVersion struct {
	XMLName xml.Name `json:"-"      xml:"productVersion"`
	Value   string   `json:"value"  xml:",chardata"`
	Build   string   `json:"build"  xml:"build,attr"`
}

type User struct {
	XMLName  xml.Name `json:"-"                   xml:"user"`
	ID       string   `json:"id,omitempty"        xml:"id,attr,omitempty"`
	Name     string   `json:"name,omitempty"      xml:"name,attr,omitempty"`
	SiteRole string   `json:"siteRole,omitempty"  xml:"siteRole,attr,omitempty"`
	FullName string   `json:"fullName,omitempty"  xml:"fullName,attr,omitempty"`
}

type Workbook struct {
	XMLName     xml.Name `json:"-"                   xml:"workbook"`
	ID          string   `json:"id,omitempty"        xml:"id,attr,omitempty"`
	Name        string   `json:"name,omitempty"      xml:"name,attr,omitempty"`
	Description string   `json:"description,omitempty"  xml:"description,attr,omitempty"`
	WebPageUrl  string   `json:"webpageurl,omitempty"  xml:"webpageurl,attr,omitempty"`
	ContentUrl  string   `json:"contentUrl,omitempty"  xml:"contentUrl,attr,omitempty"`
}

type Workbooks struct {
	XMLName       xml.Name       `json:"-"                   xml:"workbooks"`
	Workbook      *Workbook      `json:"workbook,omitempty"      xml:"workbook,omitempty"`
	ShowTabs      string         `json:"showTabs,omitempty"        xml:"showTabs,attr,omitempty"`
	Size          string         `json:"size,omitempty"        xml:"size,attr,omitempty"`
	CreatedAt     string         `json:"createdAt,omitempty"        xml:"createdAt,attr,omitempty"`
	UpdatedAt     string         `json:"updatedAt,omitempty"        xml:"updatedAt,attr,omitempty"`
	DefaultViewId *DefaultViewId `json:"defaultViewId,omitempty"        xml:"defaultViewId,attr,omitempty"`
}

type DefaultViewId struct {
	XMLName xml.Name `json:"-"                   xml:"defaultViewId"`
	Project struct {
		ID   string `json:"id,omitempty"        xml:"id,attr,omitempty"`
		Name string `json:"name,omitempty"      xml:"name,attr,omitempty"`
	}
	Owner struct {
		ID string `json:"id,omitempty"        xml:"id,attr,omitempty"`
	}
}

type ErrorType struct {
	XMLName xml.Name `json:"-"        xml:"error"`
	Summary string   `json:"summary"  xml:"summary"`
	Detail  string   `json:"detail"   xml:"detail"`
	Code    uint     `json:"code"     xml:"code,attr"`
}
