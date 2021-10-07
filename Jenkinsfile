podTemplate {
    stage('setup env') {
        node(POD_LABEL) {
            sh 'echo setting up env'
        }
    }
    stage('Run shell') {
        node(POD_LABEL) {
            sh 'echo hello world'
        }
    }
}
