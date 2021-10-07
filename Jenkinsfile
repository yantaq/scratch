podTemplate {
    stage('setup env') {
        node(POD_LABEL) {
            sh 'echo setting up env'
            checkout scm
        }
    }
    stage('Run shell') {
        node(POD_LABEL) {
            sh 'echo hello world'
        }
    }
}
