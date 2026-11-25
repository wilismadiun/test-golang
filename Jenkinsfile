
pipeline {
    agent {
        kubernetes {
            defaultContainer 'golang'

            yaml '''
apiVersion: v1
kind: Pod
spec:
  containers:
    - name: golang
      image: golang:1.26-alpine
      command:
        - cat
      tty: true
'''
        }
    }

    stages {
        stage('Checkout') {
            steps {
                container('golang') {
                    checkout scm
                }
            }
        }

        stage('Download Dependencies') {
            steps {
                container('golang') {
                    sh 'go mod download'
                }
            }
        }

        stage('Unit Testing') {
            steps {
                container('golang') {
                    sh 'go test -v ./...'
                }
            }
        }
    }
}
