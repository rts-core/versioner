podTemplate(cloud: 'kubernetes', containers: [
    containerTemplate(name: 'nix', image: 'dwdraju/alpine-curl-jq', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'golang', image: 'golang:1.16', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'docker', image: 'docker:20.10', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'helm', image: 'dtzar/helm-kubectl:3.6.0', ttyEnabled: true, command: 'cat'),
    ],
    volumes: [
        hostPathVolume(hostPath: '/var/run/docker.sock', mountPath: '/var/run/docker.sock')
    ]) {
    node(POD_LABEL) {
        checkout scm
        def applicationName = "rts-core/versioner"
        def major = 1
        def minor = 1
        def patch = 0
        def version = 'Unknown'
        stage('Get Version') {
            container('nix') {
                stage('Get Version for Build') {
                    version = sh(returnStdout: true, script: "curl -X PUT -H 'Content-Type: application/json' -d '{\"major\":${major}, \"minor\":${minor}, \"patch\":${patch}}' versioner.roeden.local/v1/applications/${java.net.URLEncoder.encode(applicationName, "UTF-8")} | jq -r '.version'").trim()
                }
            }
        }

        def versionValues = version.split('\\.')
        echo "Build: ${version}"
 
        stage('Test Stage') {
            container('golang') {
                stage('Run Tests') {
                    sh 'go test $(go list ./... | grep -v generated | grep -v mocks) -coverprofile .testcoverage'
                }
            }
        }
        stage('Deploy Stage') {
            container('docker') {
                stage('Build/Tag/Push Docker image') {                            
                    if (env.BRANCH_NAME == 'main') {
                        sh 'docker build --rm -t docker.roeden.local/versioner .'
                        sh "docker tag docker.roeden.local/versioner docker.roeden.local/versioner:${version}"
                        sh "docker push docker.roeden.local/versioner:${version}"
                        def currentBuiltVersion = ''
                        versionValues.each{ item ->
                            currentBuiltVersion = currentBuiltVersion + item
                            sh "docker tag docker.roeden.local/versioner docker.roeden.local/versioner:${currentBuiltVersion}"
                            sh "docker push docker.roeden.local/versioner:${currentBuiltVersion}"
                            currentBuiltVersion = currentBuiltVersion + '.'
                        }
                        sh 'docker tag docker.roeden.local/versioner docker.roeden.local/versioner:latest'
                        sh 'docker push docker.roeden.local/versioner:latest'
                    } else if(env.BRANCH_NAME != 'main') {
                        sh 'docker build --rm -t docker.roeden.local/versioner .'
                        sh "docker tag docker.roeden.local/versioner docker.roeden.local/versioner:${version}-beta"
                        sh "docker push docker.roeden.local/versioner:${version}-beta"
                    } else {
                        echo 'Skipped Docker Stage'
                    }
                }
            }
        }
        stage('Kube Deploy') {            
            container('helm') {
                stage('Deploy to Kube') {                         
                    if (env.BRANCH_NAME == 'main') {
                        withCredentials([file(credentialsId: 'roedenKubeConfig', variable: 'kubeConfig')]) {
                            sh 'mkdir ~/.kube'
                            sh "cp $kubeConfig ~/.kube/config"
                            sh 'apk update'
                            sh 'apk add gettext'
                            sh "VERSION=${version} envsubst < ./values.tmpl.yaml > ./values.yaml"
                            sh 'helm repo add rts http://helm-repository.roeden.local/'
                            sh 'helm repo update'
                            sh 'kubectl create namespace versioner --dry-run=client -o yaml | kubectl apply -f -'
                            sh 'helm upgrade --install --wait versioner rts/sdhc -f values.yaml -n versioner --version ~1.2'
                            sh 'rm ~/.kube/config' //cleanup
                        }                        
                    } else {
                        echo 'Skipped Deployment Stage'
                    }
                }
            }
        }
    }
}