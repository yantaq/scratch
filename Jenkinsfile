podTemplate {
    node(POD_LABEL) {
        stage('setup env') {
            sh 'echo setting up env'
            checkout scm
        }
        stage('cat Jenkinsfile') {
            sh 'cat Jenkinsfile'
        }
    }
}
