// // Run on an agent where we want to use Go
// node {
//     // Ensure the desired Go version is installed
//     def root = tool type: 'go', name: 'go1.18'
//
//     // Export environment variables pointing to the directory where Go was installed
//     withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
//         sh 'go version'
//     }
// }

pipeline {
    // install golang 1.18 on Jenkins node
    agent any
    tools {
        go 'go1.18'
    }
    environment {
//         HOME = "/tmp"
//         GOROOT = "${root}"
//         GOBIN = "${root}/bin"
//         GOPATH = "${root}/go"
//         PATH = "${root}/bin"
//         GO = "${root}/bin"
        GO114MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage("Build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
//                 sh 'go build'
                sh 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GO_FLAGS) -v -o bin/bark-send-linux'
                archiveArtifacts artifacts: 'bin/*', fingerprint: true
                sh 'cp bin/bark-send-linux /tmp/bark-send-linux'
            }
        }
    }
}