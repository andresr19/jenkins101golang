
pipeline {
    def app
    agent any
    stages {
        stage('Clone') {
            checkout scm
        }
        stage('Build') {
            steps {
                app = docker.build("kkthnxbye/jenkinsgolang101")
            }
        }
        stage('Test') {
            steps {
                echo 'testing...'
            }
        }
        stage('Push docker image') {
            steps {
                docker.withRegistry('', 'personal-docker-hub') {
                    app.push("${env.BUILD_NUMBER}")
                    app.push("latest")
                }
            }
        }
        stage('Deploy') {
           when {
                expression {
                    return env.GIT_BRANCH == "origin/master"
                }
           }
           steps {
               echo 'deploying to master'
           }
        }
        stage('Deploy Dev') {
           when {
                expression {
                    return env.GIT_BRANCH == "origin/branch1"
                }
           }
           steps {
               echo 'deploying to branch1'
           }
        }
    }
}