/**
* The Problem:
*
* We have a list of tasks. Each task can depend on other tasks.
* For example if task A depends on task B then B should run before A.
*
* Implement the "getTaskWithDependencies" method such that it returns a list of task names in the correct order.
*
* For example if I want to execute task "application A", the method should return a list with "storage,mongo,application A".
*
* List of Tasks:
*
*   - name: application A
*     dependsOn:
*       - mongo
*
*   - name: storage
*
*   - name: mongo
*     dependsOn:
*       - storage
*
*   - name: memcache
*
*   - name: application B
*     dependsOn:
*       - memcache
*
* The Golang program is expected to be executed succesfully.
 */

 package main

 import (
	 "fmt"
	 "reflect"
 )
 
 func getTaskWithDependencies(tasks []task, dependsOn string) []string {
	 // TODO: please implement logic here
	 var taskOrder []string
	 for _, v := range tasks {
		 if dependsOn != "" {
			 for _, u := range tasks {
				 if dependsOn == u.Name {
					 for _, x := range u.DependsOn {
						 getTaskWithDependencies(tasks, x)
					 }
				 }
			 }
		 }
		 for _, w := range taskOrder {
			 if v.Name != w {
				 taskOrder = append(taskOrder, v.Name)
			 }
		 }
	 }
	 return taskOrder
 }
 
 func main() {
	 verify(
		 []string{"storage", "mongo", "application A"},
		 getTaskWithDependencies(getTasks(), "application A"),
	 )
 }
 
 type task struct {
	 Name      string
	 DependsOn []string
 }
 
 func newTask(name string, dependsOn []string) task {
	 return task{
		 Name:      name,
		 DependsOn: dependsOn,
	 }
 }
 
 func getTasks() []task {
	 return []task{
		 newTask("application A", []string{"mongo"}),
		 newTask("storage", []string{}),
		 newTask("mongo", []string{"storage"}),
		 newTask("memcache", []string{}),
		 newTask("application B", []string{"memcache"}),
	 }
 }
 
 // ===== helper methods =====
 
 func verify(expected, actual []string) {
	 if !reflect.DeepEqual(expected, actual) {
		 fmt.Printf("❌ Failed test!\n")
		 fmt.Printf("  expected: %s\n", expected)
		 fmt.Printf("  actual: %s\n", actual)
		 return
	 }
 
	 fmt.Println("✅ Passed test!")
 }
 