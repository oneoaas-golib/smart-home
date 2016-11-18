package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"../crontab"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type Worker struct {
	Id           int64  		`orm:"pk;auto;column(id)" json:"id"`
	FlowId       int64  		`orm:"column(flow_id)" json:"flow_id" valid:"Required"`
	WorkflowId   int64		`orm:"column(workflow_id)" json:"workflow_id" valid:"Required"`
	DeviceAction *DeviceAction  	`orm:"rel(fk);column(device_action_id);null" json:"device_action"`
	Device       *Device		`orm:"-" json:"device"`
	Flow         *Flow		`orm:"-" json:"flow"`
	Message      *Message		`orm:"-" json:"-"`
	Node         *Node		`orm:"-" json:"-"`
	CronTask     *crontab.Task	`orm:"-" json:"-"`
	Status       string 		`orm:"size(254)" json:"status" valid:"Required"`
	Name         string 		`orm:"size(254)" json:"name" valid:"MaxSize(254);Required"`
	Time         string  		`orm:"size(254)" json:"time"`
	Created_at   time.Time		`orm:"auto_now_add;type(datetime);column(created_at)" json:"created_at"`
	Update_at    time.Time		`orm:"auto_now;type(datetime);column(update_at)" json:"update_at"`
}

func (m *Worker) TableName() string {
	return beego.AppConfig.String("db_workers")
}

func init() {
	orm.RegisterModel(new(Worker))
}

// AddWorker insert a new Worker into database and returns
// last inserted Id on success.
func AddWorker(m *Worker) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetWorkerById retrieves Worker by Id. Returns error if
// Id doesn't exist
func GetWorkerById(id int64) (v *Worker, err error) {
	o := orm.NewOrm()
	v = &Worker{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllWorker retrieves all Worker matches certain condition. Returns empty list if
// no records exist
func GetAllWorker(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, meta *map[string]int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Worker))
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

	var l []Worker
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

// UpdateWorker updates Worker by Id and returns error if
// the record to be updated doesn't exist
func UpdateWorkerById(m *Worker) (err error) {
	o := orm.NewOrm()
	v := Worker{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteWorker deletes Worker by Id and returns error if
// the record to be deleted doesn't exist
func DeleteWorker(id int64) (err error) {
	o := orm.NewOrm()
	v := Worker{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Worker{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetAllEnabledWorkersByWorkflow(workflow *Workflow) (workers []*Worker, err error) {

	o := orm.NewOrm()
	_, err = o.QueryTable(&Worker{}).Filter("workflow_id", workflow.Id).Filter("status", "enabled").All(&workers)
	for _,  worker := range workers {
		if _, err = o.LoadRelated(worker, "DeviceAction"); err != nil {
			return
		}

		worker.Device, _ = GetParentDeviceByChildId(worker.DeviceAction.DeviceId)
		worker.Device.Id = worker.DeviceAction.DeviceId
	}
	return
}

func GetAllEnabledWorkers() (workers []*Worker, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(&Worker{}).Filter("status", "enabled").All(&workers)
	for _,  worker := range workers {
		if _, err = o.LoadRelated(worker, "DeviceAction"); err != nil {
			return
		}

		worker.Device, _ = GetParentDeviceByChildId(worker.DeviceAction.DeviceId)
		worker.Device.Id = worker.DeviceAction.DeviceId
	}
	return
}

func (w *Worker) Run() bool {
	if w.CronTask == nil {
		return false
	}

	w.CronTask.Run()

	return true
}

func (w *Worker) Stop() bool {
	if w.CronTask == nil {
		return false
	}

	w.CronTask.Stop()

	return true
}

func (w *Worker) IsRun() bool {
	if w.CronTask == nil {
		return false
	}

	return w.CronTask.IsRun()
}