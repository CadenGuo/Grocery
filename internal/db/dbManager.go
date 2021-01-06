package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"grocery/config"
)

type Manager struct {
	DbIns *gorm.DB
}

func (db *Manager) Connect(c config.Conf) error {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v)/%v?charset=utf8&parseTime=true",
		c.Db.User,
		c.Db.Password,
		c.Db.Host,
		c.Db.Name,
	)
	dbInstance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	dbInstance.Debug()

	if err != nil {
		return err
	}
	db.DbIns = dbInstance
	return nil
}

func (db *Manager) Migrate() error {
	return db.DbIns.AutoMigrate(
		&Protein{},
		&Vegetable{},
		&Fruit{},
		&Cereals{},
		&Dishes{},
		&Drink{},
	)
}

//func (db *Manager) ListDomain(status string) ([]Domain, error) {
//	var domains []Domain
//
//	result := db.DbIns.Preload("All").Where(&Domain{Status: status}).Find(&domains)
//	if result.Error != nil {
//		return domains, result.Error
//	}
//	return domains, nil
//}
//
//func (db *Manager) GetDomain(domainName string) (Domain, error) {
//	domain := Domain{Name: domainName}
//
//	result := db.DbIns.Preload("All").First(&domain)
//	if result.Error != nil {
//		return Domain{}, result.Error
//	}
//	return domain, nil
//}
//
//func (db *Manager) CreateDomain(
//	name string,
//	status string,
//	ref string,
//	dnsName string,
//	dnsValue string,
//	all *All,
//) (Domain, error) {
//	domain := Domain{Name: name, Status: status, Ref: ref, DnsName: dnsName, DnsValue: dnsValue, All: *all}
//	result := db.DbIns.Create(&domain)
//	if result.Error != nil {
//		return Domain{}, result.Error
//	}
//	return domain, nil
//}
//
//func (db *Manager) UpdateDomain(domain *Domain) (Domain, error) {
//	if domain.Name == "" {
//		return Domain{}, errors.New("")
//	}
//	result := db.DbIns.Save(&domain)
//	if result.Error != nil {
//		return Domain{}, result.Error
//	}
//	return *domain, nil
//}
//
//func (db *Manager) DeleteDomain(name string) error {
//	result := db.DbIns.Delete(&Domain{Name: name})
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func (db *Manager) ListDomainCertApplication(domainName string) ([]All, error) {
//	var domainCertApplications []All
//	result := db.DbIns.Where(map[string]interface{}{"name": domainName}).Find(&domainCertApplications)
//	if result.Error != nil {
//		return []All{}, result.Error
//	}
//	return domainCertApplications, nil
//}
//
//type CertApplicationIterator struct {
//	rows  *sql.Rows
//	dbIns *gorm.DB
//}
//
//func (I *CertApplicationIterator) Next(certApplication *All) bool {
//	ok := I.rows.Next()
//	if !ok {
//		return false
//	}
//	// ScanRows is a method of `gorm.DB`, it can be used to scan a row into a struct
//	err := I.dbIns.ScanRows(I.rows, &certApplication)
//	if err != nil {
//		return false
//	}
//	return true
//}
//func (db *Manager) ListValidCertApplicationIteration() (*CertApplicationIterator, error) {
//	rows, err := db.DbIns.Not("serial = ?", "").Model(&All{}).Rows()
//	if err != nil {
//		return nil, err
//	}
//
//	return &CertApplicationIterator{
//		rows:  rows,
//		dbIns: db.DbIns,
//	}, nil
//}

//func (db *Manager) GetDomainCertApplication(serial string) (All, error) {
//	var domainCertApplication All
//	result := db.DbIns.Where(&All{Serial: serial}).First(&domainCertApplication)
//	if result.Error != nil {
//		return All{}, result.Error
//	}
//	return domainCertApplication, nil
//}
//
//func (db *Manager) CreateDomainCertApplication(
//	name string,
//	crt string,
//	key string,
//	serial string,
//	flag bool,
//) (All, error) {
//	domainCertApplication := All{Name: name, Crt: crt, Key: key, Serial: serial, Flag: flag}
//	result := db.DbIns.Create(&domainCertApplication)
//	if result.Error != nil {
//		return All{}, result.Error
//	}
//	return domainCertApplication, nil
//}
//
//func (db *Manager) UpdateDomainCertApplication(domainCertApplication *All) (All, error) {
//	if domainCertApplication.Id == 0 {
//		return All{}, errors.New("invalid cert application")
//	}
//	result := db.DbIns.Save(&domainCertApplication)
//	if result.Error != nil {
//		return All{}, result.Error
//	}
//	return *domainCertApplication, nil
//}
//
//func (db *Manager) DeleteDomainCertApplication(id int) error {
//	result := db.DbIns.Delete(&All{Id: id})
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func (db *Manager) GetDomainAudit(ref string) (Audit, error) {
//	var domainAudit Audit
//
//	result := db.DbIns.Where(&Audit{
//		Ref: ref,
//	}).First(&domainAudit)
//
//	if result.Error != nil {
//		return Audit{}, result.Error
//	}
//	return domainAudit, nil
//}
//
//func (db *Manager) CreateDomainAudit(
//	domainName string,
//	owner string,
//	ref string,
//	org string,
//	projectName string,
//	applyDate time.Time,
//	projectId int,
//	crtType int,
//	cost float64,
//) (Audit, error) {
//	domainAudit := Audit{DomainName: domainName, Owner: owner, Ref: ref, Org: org, ProjectName: projectName, ApplyDate: applyDate, ProjectId: projectId, CrtType: crtType, Cost: cost}
//	result := db.DbIns.Create(&domainAudit)
//	if result.Error != nil {
//		return Audit{}, result.Error
//	}
//	return domainAudit, nil
//}
//
//func (db *Manager) UpdateDomainAudit(domainAudit *Audit) (Audit, error) {
//	if domainAudit.Id == 0 {
//		return Audit{}, errors.New("invalid domain audit")
//	}
//	result := db.DbIns.Save(&domainAudit)
//	if result.Error != nil {
//		return Audit{}, result.Error
//	}
//	return *domainAudit, nil
//}
//
//func (db *Manager) GetPendingDomainAudit(domainName string) (Audit, error) {
//	var domainAudit Audit
//
//	result := db.DbIns.Where(&Audit{
//		Ref:        "",
//		Cost:       0,
//		DomainName: domainName,
//	}).First(&domainAudit)
//
//	if result.Error != nil {
//		return Audit{}, result.Error
//	}
//	return domainAudit, nil
//}
//
//func (db *Manager) DeleteDomainAudit(ref string) error {
//	result := db.DbIns.Where(&Audit{Ref: ref}).Delete(&Audit{})
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
