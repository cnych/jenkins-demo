podTemplate(label: 'haimaxy-jnlp', cloud: 'kubernetes') {
    node ('haimaxy-jnlp') {
        stage('pre') {
          echo "0. clone project"
          git url: "https://github.com/cnych/jenkins-demo.git"
          script {
            build_tag = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
          }
        }
        stage('test') {
          echo "1.This is in test stage"
          sleep 10
        }
        stage('build') {
          echo "2.Build docker image stage"
          sh "docker build -t cnych/jenkins-demo:${build_tag} ."
        }
        stage('push') {
          echo "3. Push docker image to docker hub stage"
          withCredentials([usernamePassword(credentialsId: 'dockerHub', passwordVariable: 'dockerHubPassword', usernameVariable: 'dockerHubUser')]) {
            sh "docker login -u ${dockerHubUser} -p ${dockerHubPassword}"
            sh "docker push cnych/jenkins-demo:${build_tag}"
          }
        }
        stage('yaml') {
            echo "4. Update kubernetes yaml file stage"
            input "Do you want to deploy this image to kubernetes env?" 
            sh "sed -i 's/<BUILD_TAG>/${build_tag}/' k8s.yaml"
        }
        stage('deploy') {
          echo "5. Deploy to kubernetes stage"
          sh "kubectl apply -f k8s.yaml"
        }
    }
}
