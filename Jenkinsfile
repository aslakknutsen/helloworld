#!/usr/bin/groovy

@Library('github.com/aslakknutsen/fabric8-pipeline-library@master')
def canaryVersion = "1.0.${env.BUILD_NUMBER}"
def utils = new io.fabric8.Utils()
def stashName = "buildpod.${env.JOB_NAME}.${env.BUILD_NUMBER}".replace('-', '_').replace('/', '_')
def envStage = utils.environmentNamespace('stage')
def envProd = utils.environmentNamespace('run')
def setupScript = null

goNode {
  checkout scm
  if (utils.isCI()) {

    goCI{}
    
  } else if (utils.isCD()) {
    
    echo 'NOTE: running pipelines for the first time will take longer as build and base docker images are pulled onto the node'
    container(name: 'go') {
      stage('Build Release') {
      }
    }
  }
}

if (utils.isCD()) {
  node {
    stage('Rollout to Stage') {
      unstash stashName
      setupScript?.setupEnvironmentPre(envStage)
      apply {
        environment = envStage
      }
      setupScript?.setupEnvironmentPost(envStage)
    }

    stage('Approve') {
      approve {
        room = null
        version = canaryVersion
        environment = 'Stage'
      }
    }
    
    stage('Rollout to Run') {
      unstash stashName
      setupScript?.setupEnvironmentPre(envProd)
      apply {
        environment = envProd
      }
      setupScript?.setupEnvironmentPost(envProd)
    }
  }
}

