pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo env.GIT_BRANCH
            }
        }
        stage('Test') {
            steps {
                echo 'testing...'
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