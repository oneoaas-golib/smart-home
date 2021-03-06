package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type MapDeviceStates_20170121_013247 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &MapDeviceStates_20170121_013247{}
	m.Created = "20170121_013247"
	migration.Register("MapDeviceStates_20170121_013247", m)
}

// Run the migrations
func (m *MapDeviceStates_20170121_013247) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE map_device_states (
	id Int( 32 ) AUTO_INCREMENT NOT NULL,
	device_state_id Int( 32 ) NOT NULL,
	map_device_id Int( 32 ) NOT NULL,
	image_id Int( 32 ) NULL,
	created_at DateTime NOT NULL,
	update_at DateTime NOT NULL,
	style Text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
	PRIMARY KEY ( id ) )
	CHARACTER SET = utf8
	COLLATE = utf8_general_ci
	ENGINE = InnoDB;`)
}

// Reverse the migrations
func (m *MapDeviceStates_20170121_013247) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS `map_device_states` CASCADE")
}
