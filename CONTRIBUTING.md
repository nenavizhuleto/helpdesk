## Branching Strategy

### Develop branch

Used for development process
All new features would be merged into `develop` branch first

### feature/[feature-name] branch

Used for developing specific application feature
After completion should be merged into `develop` branch

### Release branch

Used for testing and QA of development iteration
`release` branch should go to production after all test's are successfully completed

Hotfixes during `monitoring` phase should be commited to `release` branch and then back-merged into `develop` and `main`

### Main branch

Used for storing current stable release after `monitoring` phase
Hotfixes during `production` should be commited to `main` branch and then back-merged into `develop` and/or `release` branch

### hotfix/[issue] branch

Used for hotfixes during `production` phase
Should be merged into `main` branch

### Git workflow

Target: new feature
Steps:
1. Create new `feature` branch based of `develop` branch
2. Implement new feature and commit changes to `feature` branch
3. Create merge request to `develop` branch
4. Review and accept/decline merge request

---

Target: new release
Steps:
1. Create new merge request from `develop` branch to `release` branch
2. Accept merge request and proceed to QA/Testing
3. Make sure all tests completed successfully and code satisfies QA
4. Build application based of `release` branch and push it to production
5. Start `monitoring` phase (1d or more)
    - If application needs any hotfixes during `monitoring` phase, they should be commited to `release` branch directly
6. `release` branch should be back-merged into `develop` and `main` branch with all hotfixes and changes during `monitoring` phase

--- 

Target: hotfix during `production` phase
Steps:
1. Create new `hotfix` branch based of `main` branch
2. Fix the issue
3. Merge `hotfix` into `main` branch and push to `production`
4. `main` branch should be back-merged into `develop` and/or `release` if neccessary 
 
