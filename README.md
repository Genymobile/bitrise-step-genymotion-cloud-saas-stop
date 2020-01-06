# genymotion-cloud-saas-stop

Stop Genymotion Cloud SaaS android devices

## Prerequisite

To use this step, you must have started a Android instance on Genymotion Cloud SaaS with [Start Genymotion Cloud SaaS android devices](https://github.com/genymobile/bitrise-step-genymotion-cloud-saas-start.git) step.

## How to setup Bitrise.yml

This step take one input :
  * `instance_uuid`: instance UUID is the identifier of the Android instance. 

Example: 
`GMCLOUD_SAAS_INSTANCE_UUID` is automatically set as an output of [Start Genymotion Cloud SaaS android devices](https://github.com/genymobile/bitrise-step-genymotion-cloud-saas-start.git) step. 

```
inputs:
  - instance_uuid: $GMCLOUD_SAAS_INSTANCE_UUID
```

## How to contribute to this Step

1. Fork this repository
2. `git clone` it
3. Create a branch you'll work on
4. To use/test the step just follow the **How to use this Step** section
5. Do the changes you want to
6. Run/test the step before sending your contribution
  * You can also test the step in your `bitrise` project, either on your Mac or on [bitrise.io](https://www.bitrise.io)
  * You just have to replace the step ID in your project's `bitrise.yml` with either a relative path, or with a git URL format
  * (relative) path format: instead of `- original-step-id:` use `- path::./relative/path/of/script/on/your/Mac:`
  * direct git URL format: instead of `- original-step-id:` use `- git::https://github.com/user/step.git@branch:`
  * You can find more example of alternative step referencing at: https://github.com/bitrise-io/bitrise/blob/master/_examples/tutorials/steps-and-workflows/bitrise.yml
7. Once you're done just commit your changes & create a Pull Request
