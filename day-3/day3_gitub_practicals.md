
# Day 3

## 1. Git SHA (Secure Hashing Algorithm)

- **Definition**: Git uses SHA-1 (Secure Hashing Algorithm) to uniquely identify each commit.  
- Each commit generates a **40-character hash** that acts like a fingerprint.  
- SHA is based on:
  - Commit message
  - Author details
  - Timestamp
  - Snapshot of project files

**Importance of Git SHA:**
- Guarantees **uniqueness** of commits.
- Provides **integrity**: if even a small change is made, the SHA will be different.
- Allows you to reference commits directly using their SHA.

**Example:**
```bash
commit e3b0c44298fc1c149afbf4c8996fb92427ae41e4
Author: John Doe <johndoe@example.com>
Date:   Mon Sep 3 10:00:00 2025 +0100

    Added user authentication module
```
Here, `e3b0c44298fc1c149afbf4c8996fb92427ae41e4` is the commit SHA.

You can checkout a commit using SHA:
```bash
git checkout e3b0c44298fc1c149afbf4c8996fb92427ae41e4
```

---

## 2. Git Stash

- Temporarily saves changes in your working directory without committing them.
- Useful when you need to switch branches but don’t want to lose progress.

**Commands:**
```bash
git stash          # Save current work
git stash list     # Show stashed changes
git stash apply    # Reapply the stashed changes
git stash drop     # Delete a stash
```

**Example:**
```bash
git stash
# Switch branch
git checkout main
# Come back and restore work
git checkout feature-branch
git stash apply
```

---

## 3. Git Merge

- Combines changes from one branch into another.
- Keeps track of history and commits.

**Example:**
```bash
git checkout main
git merge feature-branch
```

---

## 4. Git Collaboration

- Clone repository to work on a shared project.
- Create a feature branch and push it to the remote.
- Submit a **Pull Request (PR)** or **Merge Request (MR)** to merge code into main.

**Workflow:**
1. Clone repository  
   ```bash
   git clone https://github.com/user/repo.git
   ```
2. Create feature branch  
   ```bash
   git checkout -b feature-login
   ```
3. Push branch  
   ```bash
   git push origin feature-login
   ```
4. Open PR/MR on GitHub and assign reviewers.

---

## 5. Git Merge Options

When merging PRs/MRs on GitHub, you get three choices:

1. **Create a Merge Commit**
   - Default option, creates a merge commit in the history.
   - Preserves complete commit history.

2. **Squash and Merge**
   - Combines all commits into a single commit.
   - Makes history cleaner.

3. **Rebase and Merge**
   - Places commits from the feature branch on top of the main branch.
   - Creates a linear history.

**Example (Squash Merge):**
```bash
# In GitHub UI, select "Squash and merge"
```

---

## 6. Git Delete (Branch Deletion)

- **Delete a branch locally:**
```bash
git branch -d feature-branch
```

- **Delete a branch remotely:**
```bash
git push origin --delete feature-branch
```

---


## 7. GitHub Actions

**What is GitHub Actions?**  
GitHub Actions is a CI/CD (Continuous Integration / Continuous Deployment) service built into GitHub.  
It allows you to automate tasks in your software development workflow.

**Why do we need GitHub Actions?**
- To **automate unit testing** whenever code is pushed or a pull request is created.
- To **automate pipelines** (build → test → deploy) without manual intervention.
- To ensure **code quality** before merging into main branches.
- To trigger workflows based on **events** such as:
  - `push` (when new commits are pushed to the repository)
  - `pull_request` (when someone creates or updates a pull request)
  - `schedule` (run at specific times, like cron jobs)
  - `workflow_dispatch` (manually triggered workflows)

**Importance:**
- Saves developer time by automating repetitive tasks.  
- Reduces human errors by ensuring consistent checks.  
- Makes collaboration smoother, as everyone’s code is tested and validated automatically.  
- Enables faster delivery with continuous integration and deployment.  

**Example Workflow File (`.github/workflows/ci.yml`):**

```yaml
name: CI Pipeline

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v3
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - name: Install dependencies
        run: go mod tidy
      - name: Run Unit Tests
        run: go test ./...
```
