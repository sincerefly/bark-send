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
        go 'go1.18'
    }
    environment {
        GOROOT = "${root}"
        PATH = "${root}/bin"
        GO = "${root}/bin"
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