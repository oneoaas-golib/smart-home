package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego"
)

type Workflow struct {
	Id   		int64  			`orm:"pk;auto;column(id)" json:"id"`
	Name		string			`orm:"" json:"name"`
	Description	string			`orm:"" json:"description"`
	Status		string			`orm:"" json:"status"`
	Scenario	*WorkflowScenario       `orm:"rel(fk);column(workflow_scenario_id);null" json:"scenario"`
	Scripts		[]*Script		`orm:"rel(m2m);rel_through(github.com/e154/smart-home/api/models.WorkflowScript)" json:"scripts"`
	Flows		[]*Flow			`orm:"-" json:"flows"`
	Created_at	time.Time		`orm:"auto_now_add;type(datetime);column(created_at)" json:"created_at"`
	Update_at	time.Time		`orm:"auto_now;type(datetime);column(update_at)" json:"update_at"`
}

func (m *Workflow) TableName() string {
	return beego.AppConfig.String("db_workflows")
}

func init() {
	orm.RegisterModel(new(Workflow))
}

// AddWorkflow insert a new Workflow into database and returns
// last inserted Id on success.
func AddWorkflow(m *Workflow) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetWorkflowById retrieves Workflow by Id. Returns error if
// Id doesn't exist
func GetWorkflowById(id int64) (v *Workflow, err error) {
	o := orm.NewOrm()
	v = &Workflow{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllWorkflow retrieves all Workflow matches certain condition. Returns empty list if
// no records exist
func GetAllWorkflow(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, meta *map[string]int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Workflow))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Workflow
	qs = qs.OrderBy(sortFields...)
	objects_count, err := qs.Count()
	if err != nil {
		return
	}
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		meta = &map[string]int64{
			"objects_count": objects_count,
			"limit":         limit,
			"offset":        offset,
		}
		return ml, meta, nil
	}
	return nil, nil, err
}

// UpdateWorkflow updates Workflow by Id and returns error if
// the record to be updated doesn't exist
func UpdateWorkflowById(m *Workflow) (err error) {
	o := orm.NewOrm()
	v := Workflow{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteWorkflow deletes Workflow by Id and returns error if
// the record to be deleted doesn't exist
func DeleteWorkflow(id int64) (err error) {
	o := orm.NewOrm()
	v := Workflow{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Workflow{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetAllEnabledWorkflow() (wfs []*Workflow, err error) {
	o := orm.NewOrm()
	wfs = []*Workflow{}
	_, err = o.QueryTable(&Workflow{}).Filter("status", "enabled").All(&wfs)

	return
}

func (wf *Workflow) GetAllEnabledWorkers() ([]*Worker, error) {
	return GetAllEnabledWorkersByWorkflow(&Workflow{Id:wf.Id})
}

func (wf *Workflow) GetAllEnabledFlows() ([]*Flow, error) {
	return GetAllEnabledFlowsByWf(wf)
}

func (wf *Workflow) GetScripts() (int64, error) {
	o := orm.NewOrm()
	return o.LoadRelated(wf, "Scripts")
}

func (wf *Workflow) GetScenario() (int64, error) {

	o := orm.NewOrm()
	if wf.Scenario != nil {
		return o.LoadRelated(wf, "Scenario")
	}

	return 0, errors.New("scenario_id is nil")
}