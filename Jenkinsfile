// // Run on an agent where we want to use Go
// node {
//     // Ensure the desired Go version is installed
//     def root = tool type: 'go', name: 'Go 1.18'
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
        go 'Go 1.18'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh 'make bl_linux'
            }
        }
    }
}