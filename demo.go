package main

import "fmt"

type task struct {
	id   string
	name string
	desc string
}

type Demo struct {
	id         string
	name       string
	pm         *benchMember
	supervisor *benchMember
	status     string
	homepage   string
	tasks      []*task
}

func NewDemo(name string, pm *benchMember, sv *benchMember) *Demo {
	return &Demo{name: name, pm: pm, supervisor: sv}
}

func (d *Demo) Id() string {
	return d.id
}
func (d *Demo) Name() string {
	return d.name
}
func (d *Demo) PM() *benchMember {
	return d.pm
}
func (d *Demo) Supervisor() *benchMember {
	return d.supervisor
}
func (d *Demo) Status() string {
	return d.status
}
func (d *Demo) Homepage() string {
	return d.homepage
}
func (d *Demo) task() []*task {
	return d.tasks
}
func (d *Demo) addTask(t *task) []*task {
	return append(d.tasks, t)
}

// TODO 需要确认
func (d *Demo) updateTask(ta *task) ([]*task, error) {
	for _, t := range d.tasks {
		if t.name == ta.name {
			//d.tasks = append(d.tasks[:i], ta, d.tasks[i:])
		}
		return d.tasks, nil
	}
	return nil, fmt.Errorf("task name not exsit.")
}

//TODO
func (d *Demo) deleteTask(ta *task) ([]*task, error) {
	return d.tasks, nil
}
