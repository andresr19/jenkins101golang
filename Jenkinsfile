
pipeline {
    def app
    agent any
    stages {
        stage('Cloning') {
            steps {
                checkout scm
            }
        }
        stage('Docker Building') {
            steps {
                app = docker.build("kkthnxbye/jenkinsgolang101")
            }
        }
        stage('Testing') {
            steps {
                echo 'testing...'
            }
        }
        stage('Docker Pushing') {
            steps {
                docker.withRegistry('', 'personal-docker-hub') {
                    app.push("${env.BUILD_NUMBER}")
                    app.push("latest")
                }
            }
        }
        stage('Stagging') {
           when {
                expression {
                    return env.GIT_BRANCH == "origin/branch1"
                }
           }
           steps {
               echo 'deploying to branch1'
           }
        }
        stage('Production') {
           when {
                expression {
                    return env.GIT_BRANCH == "origin/master"
                }
           }
           steps {
               echo 'deploying to master'
           }
        }
    }
}