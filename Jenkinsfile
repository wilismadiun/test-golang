
pipeline {
    agent {
        kubernetes {
            defaultContainer 'golang'

            yaml '''
apiVersion: v1
kind: Pod
spec:
  containers:
    # Container untuk Go Testing
    - name: golang
      image: golang:1.26-alpine
      command:
        - cat
      tty: true

    # Container untuk Build dan Push Image
    - name: kaniko
      image: gcr.io/kaniko-project/executor:v1.24.0-debug
      command:
        - /busybox/cat
      tty: true
      volumeMounts:
        - name: docker-config
          mountPath: /kaniko/.docker/

  volumes:
    - name: docker-config
      secret:
        secretName: dockerhub-credentials
        items:
          - key: .dockerconfigjson
            path: config.json
'''
        }
    }

    options {
        timestamps()
    }

    environment {
        DOCKER_IMAGE = 'USERNAME_DOCKERHUB/test-golang'
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

        stage('Build and Push') {
            steps {
                container('kaniko') {
                    sh '''
                        /kaniko/executor \
                          --context=dir://${WORKSPACE} \
                          --dockerfile=${WORKSPACE}/Dockerfile \
                          --destination=${DOCKER_IMAGE}:${BUILD_NUMBER} \
                          --destination=${DOCKER_IMAGE}:latest
                    '''
                }
            }
        }
    }
}
