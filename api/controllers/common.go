package controllers

import  (
	"github.com/astaxie/beego"
	"net/url"
	"strconv"
	"reflect"
	"html/template"
	"encoding/json"
)

type Request struct {
	Query  map[string]string                `json:"query"`
	Fields []string                        `json:"fields"`
	Sortby []string                        `json:"sortby"`
	Order  []string                        `json:"order"`
	Offset int64                        `json:"offset"`
	Limit  int64                        `json:"limit"`
}

type CommonController struct {
	beego.Controller
}

func (c *CommonController) ErrHan(code int, message string) {
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.Data["json"] = &map[string]interface{}{"status":"error", "message": message}
	c.ServeJSON()
}

func (c *CommonController) pagination() (query map[string]string, fields []string, sortby []string, order []string,
offset int64, limit int64) {

	link, _ := url.ParseRequestURI(c.Ctx.Request.URL.String())
	q := link.Query()

	query = map[string]string{}
	fields = []string{}
	sortby = []string{}
	order = []string{}

	values, _ := url.ParseQuery(q.Encode())
	if val, ok := values["query"]; ok {
		for _, v := range val {
			json.Unmarshal([]byte(v), &query)
		}
	}

	if val, ok := q["fields"]; ok {
		for _, v := range val {
			fields = append(fields, v)
		}
	}

	if val, ok := q["sortby"]; ok {
		for _, v := range val {
			sortby = append(sortby, v)
		}
	}

	if val, ok := q["order"]; ok {
		for _, v := range val {
			order = append(order, v)
		}
	}

	if val, ok := q["offset"]; ok {
		offset, _ = strconv.ParseInt(val[0], 10, 0)
	}

	if val, ok := q["limit"]; ok {
		limit, _ = strconv.ParseInt(val[0], 10, 0)
	}

	return
}

func (c *CommonController) GetTemplate() string {

	templatetype := beego.AppConfig.String("template_type")
	if templatetype == "" {
		templatetype = "default"
	}
	return templatetype
}

func (c *CommonController) Prepare() {


}

func init() {
	beego.AddFuncMap("safeHtml", func(s string) template.HTML {return template.HTML(s)})
	beego.AddFuncMap("safeCss", func(s string) template.CSS {return template.CSS(s)})
	beego.AddFuncMap("safeUrl", func(s string) template.URL {return template.URL(s)})
	beego.AddFuncMap("safeJs", func(s string) template.JS {return template.JS(s)})
	beego.AddFuncMap("attr", func(s string) template.HTMLAttr {return template.HTMLAttr(s)})
	beego.AddFuncMap("last", func(i int, s interface{}) bool {return i == reflect.ValueOf(s).Len() - 1})
	beego.AddFuncMap("len", func(s interface{}) int {return reflect.ValueOf(s).Len()})
}