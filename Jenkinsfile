@Library('sharedLibrary') _

def REPO = getRepoName()

podTemplate {
    node(POD_LABEL) {
        stage('setup env') {
            sh 'echo setting up env'
            checkout scm
        }
        stage('cat Jenkinsfile') {
            sh 'cat Jenkinsfile'
            sh "cat repo: $REPO"
        }
    }
}

