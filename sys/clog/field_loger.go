// package clog
// @Author cuisi
// @Date 2023/10/31 15:13:00
// @Desc
package clog

import (
	"github.com/sirupsen/logrus"
)

type Entity struct {
	data map[string]interface{}
}

func WidthFields(m map[string]interface{}) *Entity {
	return ler.WidthFields(m)
}

func WidthField(k string, v interface{}) *Entity {
	return ler.WidthField(k, v)
}

func (e Entity) Info(i ...interface{}) {
	logrus.WithFields(e.data).Info(i...)
}
func (e Entity) Error(i ...interface{}) {
	logrus.WithFields(e.data).Error(i...)
}
func (e Entity) Panic(i ...interface{}) {
	logrus.WithFields(e.data).Panic(i...)
}
func (e Entity) Fatal(i ...interface{}) {
	logrus.WithFields(e.data).Fatal(i...)
}
func (e Entity) Warning(i ...interface{}) {
	logrus.WithFields(e.data).Warning(i...)
}
func (e Entity) Debug(i ...interface{}) {
	logrus.WithFields(e.data).Debug(i...)
}
func (e Entity) Trace(i ...interface{}) {
	logrus.WithFields(e.data).Trace(i...)
}
