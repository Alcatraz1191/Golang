pipeline {
  agent {
   label createDynamicAnkaNode(
      masterVmId: 'd4c7bd5a-4885-4bcd-80b3-ab7c1aba30c2',
      tag: 'v1',
      nameTemplate: 'simple-example'
    )
  }
   stages {
     stage("hello") {
       steps {
         sh "echo hello"
       }
     }
  }
}