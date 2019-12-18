package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Clues struct {
	Id            int       `orm:"column(id);auto"`
	AcceptTime    time.Time `orm:"column(accept_time);type(datetime);null" description:"线索接受时间"`
	Source        string    `orm:"column(source);size(255);null" description:"来源"`
	SmugglerNature *SmugglerNature `orm:"rel(one);null"`	//性质	
	SmugglingChannels *SmugglingChannels `orm:"rel(one);null"`	//走私渠道
	SmugglingObject *SmugglingObject `orm:"rel(one);null"`	//走私对象
	CompanyName   string    `orm:"column(company_name);size(255);null" description:"企业名称"`
	SmugglerPerson []*SmugglerPerson `orm:"reverse(many);null"`	//嫌疑人
	CompanyCode   string    `orm:"column(company_code);size(50);null" description:"企业编码"`
	SmugglingWay *SmugglingWay `orm:"rel(one);null"`	//走私方式 	
	SmugglingCategory *SmugglingCategory `orm:"rel(one);null"`	//类别
	SmugglingCharge *SmugglingCharge `orm:"rel(one);null"`	//走私罪名
	Smuggle       string    `orm:"column(smuggle);size(255);null" description:"走私标的名称"`
	SmugglingNature *SmugglingNature `orm:"rel(one);null"`	//贸易性质
	BNum          string    `orm:"column(b_num);size(50);null" description:"报关单号"`
	JNum          string    `orm:"column(j_num);size(50);null" description:"集装箱号码"`
	KNum          string    `orm:"column(k_num);size(50);null" description:"快递单号"`
	HNum          string    `orm:"column(h_num);size(50);null" description:"海运提单号码"`
	YNum          string    `orm:"column(y_num);size(50);null" description:"运输工具号码"`
	CNum          string    `orm:"column(c_num);size(50);null" description:"舱单号码"`
	Address       string    `orm:"column(address);size(255);null" description:"地址"`
	Count         string    `orm:"column(count);size(50);null" description:"数量"`
	Amount        string    `orm:"column(amount);size(50);null" description:"金额"`
	Case          string    `orm:"column(case);size(50);null" description:"简要案情"`
	Opinion       string    `orm:"column(opinion);size(255);null" description:"处理意见"`
	Investigators string    `orm:"column(investigators);size(255);null" description:"主办侦查人员名字"`
	Note          string    `orm:"column(note);size(255);null" description:"备注"`
	SmugglerBank []*SmugglerBank `orm:"reverse(many);null"` //银行开户
	Email         string    `orm:"column(email);size(255);null" description:"电子邮件"`
	Phone         int       `orm:"column(phone);null" description:"联系电话"`
	Goods          string    `orm:"column(goods);size(50);null" description:"商品名称"`
	Rate          string    `orm:"column(rate);size(50);null" description:"税率"`
	Escape          string    `orm:"column(escape);size(50);null" description:"偷逃税款金额"`
	Country          string    `orm:"column(country);size(50);null" description:"进口国别"`
	Place          string    `orm:"column(place);size(50);null" description:"起运地"`
	Entry          string    `orm:"column(entry);size(255);null" description:"进境口岸"`
	Exit          string    `orm:"column(exit);size(255);null" description:"出境口岸"`
	SmugglingHandle *SmugglingHandle `orm:"rel(one);null"`	//办理状态
	State         int       `orm:"column(state);null" description:"0:未审批,1:通过审批,2:完结"`
	CreateTime    time.Time `orm:"column(create_time);type(datetime);null" description:"自己定义的创建时间"`
	ModiTime      time.Time `orm:"column(modi_time);type(datetime);null" description:"自己定义的最后一次修改时间"`
	ModiCount     int       `orm:"column(modi_count);null" description:"自己定义的修改次数"`
	Department *Department `orm:"rel(fk);null"` //部门
}

func (t *Clues) TableName() string {
	return "clues"
}

func init() {
	orm.RegisterModel(new(Clues))
}

// AddClues insert a new Clues into database and returns
// last inserted Id on success.
func AddClues(m *Clues) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCluesById retrieves Clues by Id. Returns error if
// Id doesn't exist
func GetCluesById(id int) (v *Clues, err error) {
	o := orm.NewOrm()
	v = &Clues{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllClues retrieves all Clues matches certain condition. Returns empty list if
// no records exist
func GetAllClues(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Clues))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Clues
	qs = qs.OrderBy(sortFields...)
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
		return ml, nil
	}
	return nil, err
}

// UpdateClues updates Clues by Id and returns error if
// the record to be updated doesn't exist
func UpdateCluesById(m *Clues) (err error) {
	o := orm.NewOrm()
	v := Clues{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteClues deletes Clues by Id and returns error if
// the record to be deleted doesn't exist
func DeleteClues(id int) (err error) {
	o := orm.NewOrm()
	v := Clues{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Clues{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
